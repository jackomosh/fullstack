#!/bin/bash
set -e

BASE="http://localhost:8080/api"
PASS=0
FAIL=0

check() {
  local label=$1
  local expected=$2
  local actual=$3
  if echo "$actual" | grep -q "$expected"; then
    echo "  ✅ $label"
    PASS=$((PASS+1))
  else
    echo "  ❌ $label"
    echo "     Expected: $expected"
    echo "     Got:      $actual"
    FAIL=$((FAIL+1))
  fi
}

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  BursaryHub API Test Suite"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# ── Health ────────────────────────────────────────────────────
echo ""
echo "[ Health Check ]"
R=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health 2>/dev/null || echo "000")
check "GET /health returns 200" "200" "$R"

# ── Auth ──────────────────────────────────────────────────────
echo ""
echo "[ Auth Flow ]"

R=$(curl -s -X POST $BASE/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone":"+254700000001"}' 2>/dev/null || echo "{}")
check "POST /auth/login returns OTP confirmation" "otp\|message" "$R"

# Use dev stub OTP
R=$(curl -s -X POST $BASE/auth/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"phone":"+254700000001","otp":"123456"}' 2>/dev/null || echo "{}")
check "POST /auth/verify-otp returns JWT token" "token" "$R"

export SCHOOL_TOKEN=$(echo $R | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
echo "  🔑 School JWT: ${SCHOOL_TOKEN:0:40}..."

# Login as donor for donor endpoints
R=$(curl -s -X POST $BASE/auth/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"phone":"+254711000001","otp":"123456"}' 2>/dev/null || echo "{}")
export DONOR_TOKEN=$(echo $R | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

# ── Donor ─────────────────────────────────────────────────────
echo ""
echo "[ Donor Endpoints ]"

R=$(curl -s -X POST $BASE/donor/scholarships \
  -H "Authorization: Bearer $DONOR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title":"API Test Scholarship",
    "coverage_type":"tuition",
    "max_amount_per_student":55000,
    "number_of_slots":5,
    "eligible_courses":["Computer Science"],
    "eligible_years":[1],
    "min_gpa":3.0
  }' 2>/dev/null || echo "{}")
check "POST /donor/scholarships creates scholarship" "scholarship_id\|status" "$R"

R=$(curl -s $BASE/donor/disbursements \
  -H "Authorization: Bearer $DONOR_TOKEN" 2>/dev/null || echo "{}")
check "GET /donor/disbursements returns list" "disbursements\|count" "$R"

R=$(curl -s "$BASE/donor/cost-breakdown?amount=10000&currency=USD" \
  -H "Authorization: Bearer $DONOR_TOKEN" 2>/dev/null || echo "{}")
check "GET /donor/cost-breakdown returns fee breakdown" "platform_fee\|usdt_locked\|donor_deposit" "$R"

# ── School ────────────────────────────────────────────────────
echo ""
echo "[ School Endpoints ]"

R=$(curl -s -X POST $BASE/school/fee-master \
  -H "Authorization: Bearer $SCHOOL_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "academic_year":"2025-2026",
    "course":"Computer Science",
    "year_of_study":1,
    "tuition_amount":55000,
    "accommodation_amount":25000,
    "food_amount":15000,
    "transport_amount":5000
  }' 2>/dev/null || echo "{}")
check "POST /school/fee-master publishes fee structure" "success\|status" "$R"

R=$(curl -s -X POST $BASE/school/fee-master/bulk-update \
  -H "Authorization: Bearer $SCHOOL_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "academic_year":"2025-2026",
    "course":"Computer Science",
    "year_of_study":1,
    "new_tuition":60000
  }' 2>/dev/null || echo "{}")
check "POST /school/fee-master/bulk-update applies auto-balance" "students_updated\|formula\|success" "$R"

# ── Student ───────────────────────────────────────────────────
echo ""
echo "[ Student Endpoints ]"

# Login as student for student endpoints
R=$(curl -s -X POST $BASE/auth/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"phone":"+254720000001","otp":"123456"}' 2>/dev/null || echo "{}")
export STUDENT_TOKEN=$(echo $R | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

R=$(curl -s $BASE/student/scholarships \
  -H "Authorization: Bearer $STUDENT_TOKEN" 2>/dev/null || echo "{}")
check "GET /student/scholarships returns available scholarships" "scholarships\|count" "$R"

R=$(curl -s $BASE/student/balance \
  -H "Authorization: Bearer $STUDENT_TOKEN" 2>/dev/null || echo "{}")
check "GET /student/balance returns balance" "balance\|amount\|coverage" "$R"

R=$(curl -s -X POST $BASE/student/three-way-verify \
  -H "Authorization: Bearer $STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"disbursement_id":1,"entered_amount":55000}' 2>/dev/null || echo "{}")
check "POST /student/three-way-verify returns match result" "match\|student_amount" "$R"

# ── Three-Way Mismatch Test ───────────────────────────────────
echo ""
echo "[ Three-Way Verification — Mismatch Block ]"

R=$(curl -s -X POST $BASE/student/three-way-verify \
  -H "Authorization: Bearer $STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"disbursement_id":1,"entered_amount":99999}' 2>/dev/null || echo "{}")
check "Mismatch amount is correctly blocked" "mismatch\|false\|blocked\|pending" "$R"

# ── OTP Flow ──────────────────────────────────────────────────
echo ""
echo "[ OTP Flow ]"

R=$(curl -s -X POST $BASE/student/claims/1/request-otp \
  -H "Authorization: Bearer $STUDENT_TOKEN" 2>/dev/null || echo "{}")
check "POST /student/claims/1/request-otp sends OTP" "otp\|sent\|success\|status" "$R"

# ── Admin ─────────────────────────────────────────────────────
echo ""
echo "[ Admin Endpoints ]"

R=$(curl -s $BASE/admin/mismatches \
  -H "Authorization: Bearer $SCHOOL_TOKEN" 2>/dev/null || echo "{}")
check "GET /admin/mismatches returns mismatch list" "mismatches" "$R"

# ── Auth Guard ────────────────────────────────────────────────
echo ""
echo "[ JWT Protection ]"

R=$(curl -s -o /dev/null -w "%{http_code}" $BASE/donor/scholarships 2>/dev/null || echo "000")
check "Unauthenticated request returns 401" "401" "$R"

# ── Summary ───────────────────────────────────────────────────
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Results: ✅ $PASS passed  ❌ $FAIL failed"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ $FAIL -gt 0 ]; then exit 1; fi