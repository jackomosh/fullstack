# BursaryHub — Technical Documentation

## Table of Contents

1. [System Overview](#1-system-overview)
2. [Architecture](#2-architecture)
3. [Database](#3-database)
4. [Backend](#4-backend)
5. [Smart Contracts](#5-smart-contracts)
6. [Frontend](#6-frontend)
7. [Development Setup](#7-development-setup)
8. [Testing](#8-testing)
9. [Deployment](#9-deployment)
10. [API Quick Reference Card](#10-api-quick-reference-card)

---

## 1. System Overview

### 1.1 Purpose

BursaryHub eliminates scholarship fraud and payment leakage in Kenya's education funding system. It serves three user types:

- **Donors** - Foundations, NGOs, and individuals creating scholarships
- **Schools** - Universities and colleges publishing fee structures and receiving payments
- **Students** - Beneficiaries applying for scholarships and approving disbursements

### 1.2 Core Principles

| Principle | Implementation |
|-----------|----------------|
| Zero subscription | Schools and students pay no platform fees |
| 1% success fee | Fee collected only on successful disbursement (platformFeeBps = 100) |
| Three-way verification | Student + School + Fee Master must match before payment |
| Auto-balance formula | Automated calculation when fee master changes |
| Soulbound USDT escrow | Tokens locked in smart contract until verification passes |

### 1.3 User Roles

| Role | Description | Permissions |
|------|-------------|-------------|
| donor | Creates scholarships, deposits funds | Create scholarships, view disbursements, cost breakdown |
| school_admin | Publishes fee masters, creates claims | Fee master CRUD, claims, three-way verification, disbursements |
| student | Applies for scholarships, approves payments | View scholarships, submit balance, request/approve OTP |
| admin | Resolves mismatches | Whitelist schools, approve donor KYC, resolve mismatches |

### 1.4 Data Flow

1. **Login** - User sends phone to `/api/auth/login`, receives OTP, verifies with `/api/auth/verify-otp`
2. **Scholarship Creation** - Donor creates scholarship with eligibility criteria
3. **Fee Master Publication** - School publishes fee structure for courses/years
4. **Student Application** - Student applies for matching scholarship
5. **Payment Claim** - School creates claim with coverage type and amount
6. **Three-Way Verification** - Student submits amount, school confirms, system compares to fee master
7. **OTP Approval** - Student requests and enters OTP to approve disbursement
8. **Disbursement** - USDT transferred from escrow to school (minus 1% platform fee)

---

## 2. Architecture

### 2.1 Component Overview

```
┌─────────────────────┐
│ Frontend (3 clients)│
│ index.html          │
│ grantor/dashboard.html│
│ institution/dashboard.html│
│ beneficiary/portal.html│
└──────────┬──────────┘
           │ HTTP/JSON (fetch API)
           ▼
┌─────────────────────┐
│ Go Backend (port 8080)│
│ Gorilla Mux router  │
│ JWT middleware      │
└──────────┬──────────┘
           │ SQL
           ▼
┌─────────────────────┐
│ PostgreSQL 15       │
│ (Docker container)  │
└─────────────────────┘

┌─────────────────────┐
│ Smart Contracts     │
│ Polygon/Celo testnet│
└─────────────────────┘

┌─────────────────────┐
│ External Services   │
│ M-Pesa (STK push)   │
│ Africa's Talking    │
└─────────────────────┘
```

### 2.2 Request Lifecycle

1. Browser loads HTML from `python3 -m http.server`
2. JavaScript calls `fetch()` to `/api/...` endpoints
3. Request hits Gorilla Mux router in `backend/main.go`
4. JWT middleware in `backend/middleware/auth.go` validates token
5. Handler in `backend/handlers/` processes request
6. Repository in `backend/repository/` executes SQL
7. PostgreSQL returns data
8. Handler returns JSON response

### 2.3 Directory Structure

| File | Responsibility |
|------|----------------|
| backend/main.go | Server entry point, routes, CORS middleware |
| backend/middleware/auth.go | JWT verification, user context extraction |
| backend/handlers/auth.go | OTP login, verification handlers |
| backend/handlers/donor.go | Scholarship CRUD, disbursement endpoints |
| backend/handlers/school.go | Fee master, claims, verification handlers |
| backend/handlers/student.go | Scholarship list, balance, OTP approval |
| backend/handlers/admin.go | Mismatches, whitelist, KYC handlers |
| backend/handlers/claims.go | Payment claim creation logic |
| backend/handlers/feemaster.go | Fee master CRUD with auto-balance |
| backend/repository/db.go | Database connection pool, GetDB() singleton |
| backend/repository/admin_repo.go | User lookups, mismatch queries |
| backend/repository/donor_repo.go | Scholarship and disbursement queries |
| backend/repository/school_repo.go | Fee master, claims, balance queries |
| backend/repository/student_repo.go | Student profile, applications, disbursements |
| backend/services/otp.go | 6-digit OTP generation, 5-min expiry |
| backend/services/fees.go | Fee calculation (15+2 USD fees, 1% platform) |
| backend/services/mpesa.go | M-Pesa integration stubs |
| backend/services/blockchain.go | Smart contract call stubs |
| db/schema.sql | 12 tables, 5 ENUMs, 6 indexes |
| db/seeds.sql | 28 users, 5 schools, 20 students, 3 scholarships |

---

## 3. Database

### 3.1 Schema Overview

| Table | Purpose | Seed Count |
|-------|---------|------------|
| users | Authentication, roles | 28 |
| donors | Donor profiles | 3 |
| schools | School profiles | 5 |
| students | Student profiles | 20 |
| fee_master | Fee structures per course/year | 9 |
| student_balances | Student fee balances | 5 |
| scholarships | Scholarship definitions | 3 |
| applications | Student applications | - |
| disbursements | Payment records | - |
| three_way_verification | Verification logs | - |
| transaction_fees | Fee breakdowns | - |
| audit_logs | Action logs | - |

### 3.2 Full Schema

**ENUM Types:**
- `user_role`: 'donor', 'school_admin', 'student', 'admin'
- `coverage_type`: 'tuition', 'accommodation', 'food', 'transport', 'all', 'unrestricted'
- `disbursement_status`: 'pending', 'verified', 'completed', 'failed', 'blocked'
- `scholarship_status`: 'active', 'closed', 'expired' (referenced but not in schema.sql)
- `application_status`: 'pending', 'approved', 'rejected'

**Foreign Key Relationships:**
- `donors.user_id` → `users.id`
- `schools.user_id` → `users.id`
- `students.user_id` → `users.id`
- `students.school_id` → `schools.id`
- `fee_master.school_id` → `schools.id`
- `scholarships.donor_id` → `donors.id`
- `applications.scholarship_id` → `scholarships.id`
- `applications.student_id` → `students.id`
- `disbursements.scholarship_id` → `scholarships.id`
- `disbursements.student_id` → `students.id`
- `disbursements.school_id` → `schools.id`
- `three_way_verification.disbursement_id` → `disbursements.id`
- `transaction_fees.disbursement_id` → `disbursements.id`
- `audit_logs.user_id` → `users.id`

### 3.3 Auto-Balance Formula

Implemented in `backend/handlers/feemaster.go`:

```
New Balance = Previous Unpaid + (New Fee - Paid to Date)
```

When school runs bulk update with new tuition, all students in that course/year get recalculated.

### 3.4 Seed Data

**Schools:**
- University of Nairobi (EDU-KE-2024-001)
- JKUAT (EDU-KE-2024-002)
- Maseno University (EDU-KE-2024-003)
- Egerton University (EDU-KE-2024-004)
- Kenyatta University (EDU-KE-2024-005)

**Seed Users:** 5 school admins, 3 donors, 20 students across Computer Science, IT, Engineering, Business, Nursing, Education, and Law programs.

---

## 4. Backend

### 4.1 Server Setup

- **Port:** 8080 (from `backend/main.go`)
- **Router:** Gorilla Mux v1.8.1
- **CORS:** Allows all origins, headers include Authorization

### 4.2 Authentication

**OTP Flow (auth.go):**
1. POST `/api/auth/login` with phone - generates 6-digit random OTP
2. POST `/api/auth/verify-otp` with phone + OTP - returns JWT token
3. Dev stub: OTP "123456" accepted for seeded phone numbers

**JWT Structure:**
- Signing method: HS256
- Claims: user_id (uint), role (string), exp (Unix timestamp)
- Expiry: 24 hours
- Secret: From `JWT_SECRET` env var, defaults to "bursaryhub-dev-secret-change-in-production"

### 4.3 API Endpoints - Full Reference

#### POST /api/auth/login
- Auth: No
- Body: `{"phone": "+254700000001"}`
- Response: `{"message": "OTP sent successfully", "status": "pending"}`

#### POST /api/auth/verify-otp
- Auth: No
- Body: `{"phone": "+254700000001", "otp": "123456"}`
- Response: `{"token": "jwt...", "user": {"id": 1, "role": "school_admin"}}`

#### POST /api/donor/scholarships
- Auth: Yes (donor role)
- Body: `{"title", "coverage_type", "max_amount_per_student", "number_of_slots", "eligible_courses", "eligible_years", "min_gpa"}`
- Response: `{"scholarship_id": X, "status": "created", "matching_students": 0}`

#### GET /api/donor/scholarships
- Auth: Yes (donor role)
- Response: `{"scholarships": [...], "count": N}`

#### GET /api/donor/disbursements
- Auth: Yes (donor role)
- Response: `{"disbursements": [...], "count": N}`

#### GET /api/donor/cost-breakdown
- Auth: Yes (donor role)
- Query: `amount=10000&currency=USD`
- Response: From `services/fees.go` - `{"donor_deposit", "platform_fee", "usdt_locked", ...}`

#### POST /api/school/fee-master
- Auth: Yes (school_admin role)
- Body: `{"academic_year", "course", "year_of_study", "tuition_amount", "accommodation_amount", "food_amount", "transport_amount"}`
- Response: `{"success": true, "status": "published"}`

#### POST /api/school/fee-master/bulk-update
- Auth: Yes (school_admin role)
- Body: `{"academic_year", "course", "year_of_study", "new_tuition"}`
- Response: `{"success": true, "students_updated": N, "formula": "New Balance = Previous Unpaid + (New Fee - Paid to Date)"}`

#### GET /api/student/scholarships
- Auth: Yes (student role)
- Response: `{"scholarships": [...], "count": N}`

#### POST /api/student/three-way-verify
- Auth: Yes (student role)
- Body: `{"disbursement_id", "entered_amount"}`
- Response: `{"match": bool, "student_amount", "status": "pending"}`

#### POST /api/admin/mismatches
- Auth: Yes (school_admin/admin role)
- Response: `{"mismatches": [...], "count": N}`

### 4.4 Fee Calculation Engine

From `backend/services/fees.go`:

| Fee | Formula | When Applied |
|-----|---------|--------------|
| Conversion In Fee | amount × 0.0015 | On donor deposit to USDT |
| Network Gas Fee | Fixed $2 | On deposit |
| USDT Locked | amount - conversionFee - gasFee | After deposit |
| Platform Fee | USDT amount × 0.01 | On disbursement |
| Conversion Out | USDT × 0.01 | On disbursement |
| Withdrawal Fee | Fixed 100 KSH | On disbursement |

**Worked Example ($10,000 deposit):**
- Donor Deposit: $10,000
- Conversion Fee (0.15%): $15
- Gas Fee: $2
- USDT Locked: $9,983
- When disbursed to school:
  - Platform Fee (1%): KSH 7,000
  - Conversion to KSH (1%): KSH 500
  - Withdrawal Fee: KSH 100
  - School receives: KSH 9283.17

---

## 5. Smart Contracts

### 5.1 BursaryEscrow.sol

**Solidity Version:** ^0.8.19

**Inherited Contracts:** Ownable, ReentrancyGuard

**State Variables:**
- `usdt`: IERC20 token address
- `platformFeeBps`: 100 (1%)
- `feeCollector`: Address for fee collection
- `nextDisbursementId`: Auto-increment counter

**Key Functions:**
- `deposit(scholarshipId, amount)` - Donor deposits USDT (nonReentrant)
- `executeDisbursement(disbursementId, school, amountUSDT, matchHash)` - Owner only, sends to school minus 1% fee
- `setWhitelistedSchool(school, status)` - Owner only
- `setWhitelistedDonor(donor, status)` - Owner only

### 5.2 VendorRegistry.sol

Whitelisting contract with:
- `addSchool(school)` / `removeSchool(school)` - Owner only
- `addDonor(donor)` / `removeDonor(donor)` - Owner only
- `isSchoolWhitelisted(school)` / `isDonorWhitelisted(donor)` - View functions

### 5.3 Deployment

Network: Hardhat (localhost:8545)

```bash
npx hardhat run scripts/deploy.js
```

### 5.4 Tests

```bash
npx hardhat test
```

Tests cover: deposits, disbursements, whitelist management, soulbound transfer prevention.

---

## 6. Frontend

### 6.1 Landing Page (index.html)

12 sections: Navigation, Hero, Trust Badges, Problem, How It Works, Who It's For, Key Features, Cost Transparency Calculator, Testimonials, FAQ, CTA, Footer.

### 6.2 Donor Dashboard (grantor/)

- Summary statistics cards (total donated, active scholarships, students funded)
- Create scholarship form (title, coverage type, amount, slots, GPA)
- Cost preview box
- Active scholarships table
- Pending applications queue
- Recent disbursements
- Impact report panel

API calls: `POST /api/donor/scholarships`, `GET /api/donor/scholarships`, `POST /api/scholarships/{id}/applications/{appId}/approve`

### 6.3 School Dashboard (institution/)

- Header with school name, registration, verification status
- Summary statistics (students, pending claims, completed, active scholarships)
- Fee master management panel with bulk update
- Auto-balance preview table
- Student balances table with export
- Three-way verification panel
- Digital payment claims table
- Recent disbursements
- Bank account settings

API calls: `POST /api/school/fee-master`, `POST /api/school/fee-master/bulk-update`

### 6.4 Student Portal (beneficiary/)

Mobile-first design with:
- Available scholarship notification
- My applications list
- STEP 1: Confirm balance (amount input, upload, verification status)
- STEP 2: Approve payment (OTP request, 5-min countdown, approve button)
- Payment history
- Help & support footer

### 6.5 Shared Frontend Utilities

**api.js:** `apiRequest()`, `login()`, `verifyOTP()`, `get()`, `post()`, `put()`, `del()` - wraps fetch with JWT injection

**auth.js:** `saveAuth()`, `getToken()`, `getUser()`, `logout()`, `loginAndRedirect()`

---

## 7. Development Setup

### 7.1 Prerequisites

- Go 1.21+
- Docker 27.3+ (rootless compatible)
- Node.js 18+ with npm
- Python 3+

### 7.2 Environment Variables

See README.md for full table.

### 7.3 Database Setup (Rootless Docker)

```bash
# Start PostgreSQL container
make db-up

# Wait for ready
# Database accessible at: postgres://bursary_user:bursary_pass@localhost:5432/bursaryhub
```

### 7.4 Running Migrations

```bash
make migrate   # Uses docker exec, no psql needed
make seed
make db-check  # Lists all tables
```

### 7.5 Starting the Server

```bash
setsid env GOPATH=~/go go run ./backend/main.go
# Or for background:
nohup env GOPATH=~/go go run ./backend/main.go > server.log 2>&1 &
```

### 7.6 Serving the Frontend

```bash
cd frontend && python3 -m http.server 3000
```

### 7.7 Makefile Reference

| Target | Description |
|--------|-------------|
| dev | Start Go backend |
| build | Build binary to bin/bursaryhub |
| db-up | Start PostgreSQL container |
| db-down | Stop and remove container |
| db-shell | Open psql shell in container |
| migrate | Run schema.sql via docker exec |
| seed | Run seeds.sql via docker exec |
| db-check | List all tables |
| db-reset | Drop schema, rerun migrate + seed |
| ping | Test database connection |
| test | Run Go tests |
| test-api | Run handler tests |
| test-contracts | Run Hardhat tests |
| ship | Full setup + test + build |

---

## 8. Testing

### 8.1 Database Connection Test

```bash
make ping
# Output:
# 🔌 Testing database connection...
# ✅ Database connected successfully
```

### 8.2 API Test Suite

```bash
bash scripts/test-api.sh
```

Tests covered:
- Health check (200 OK)
- Auth login flow (OTP request)
- Auth OTP verification (JWT token)
- Donor endpoints (create scholarship, disbursements, cost breakdown)
- School endpoints (fee master, bulk update)
- Student endpoints (scholarships, balance, three-way verify)
- Three-way mismatch blocking
- OTP flow
- Admin endpoints (mismatches)
- JWT protection (401 on unauthenticated)

### 8.3 Go Tests

```bash
make test
# Tests packages: ./...
```

### 8.4 Contract Tests

```bash
make test-contracts
# npx hardhat test
```

---

## 9. Deployment

### 9.1 Current Status

- M-Pesa: Stub mode (requires real API keys)
- Africa's Talking: Stub mode
- Blockchain: Stub mode (requires RPC endpoint and contract addresses)

### 9.2 Environment Checklist

Before production:
- [ ] Change `JWT_SECRET` from dev default
- [ ] Add real M-Pesa credentials
- [ ] Add real Africa's Talking credentials
- [ ] Add real Infura/RPC endpoint
- [ ] Deploy contracts to mainnet, update addresses
- [ ] Set `sslmode=require` in DATABASE_URL

---

## 10. API Quick Reference Card

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /api/auth/login | No | Request OTP |
| POST | /api/auth/verify-otp | No | Verify OTP, get JWT |
| POST | /api/donor/scholarships | Donor | Create scholarship |
| GET | /api/donor/scholarships | Donor | List scholarships |
| GET | /api/donor/scholarships/{id}/applications | Donor | View applications |
| POST | /api/donor/scholarships/{id}/applications/{appId}/approve | Donor | Approve student |
| GET | /api/donor/disbursements | Donor | View disbursements |
| GET | /api/donor/cost-breakdown | Donor | Fee preview |
| POST | /api/school/fee-master | School | Publish fee structure |
| POST | /api/school/fee-master/bulk-update | School | Update all balances |
| PUT | /api/school/students/{id}/balance | School | Update student balance |
| POST | /api/school/claims | School | Create payment claim |
| GET | /api/school/claims | School | List claims |
| POST | /api/school/three-way-verify | School | Submit verification |
| GET | /api/school/disbursements | School | View disbursements |
| GET | /api/student/scholarships | Student | Available scholarships |
| POST | /api/student/scholarships/{id}/apply | Student | Apply for scholarship |
| GET | /api/student/balance | Student | Get balance |
| POST | /api/student/three-way-verify | Student | Submit balance |
| POST | /api/student/claims/{id}/request-otp | Student | Request OTP |
| POST | /api/student/claims/{id}/approve | Student | Approve with OTP |
| GET | /api/student/disbursements | Student | View disbursements |
| POST | /api/admin/schools/{id}/whitelist | Admin | Whitelist school |
| POST | /api/admin/donors/{id}/kyc | Admin | Approve donor KYC |
| GET | /api/admin/mismatches | Admin | Get mismatches |
| POST | /api/admin/mismatches/{id}/resolve | Admin | Resolve mismatch |

---

**Documentation maintained from code analysis as of 2026-05-29**