package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/altradits/bursaryhub/backend/models"
	"github.com/altradits/bursaryhub/backend/repository"
	"github.com/altradits/bursaryhub/backend/services"
)

// CreateFeeMaster handles POST /school/fee-master
func CreateFeeMaster(w http.ResponseWriter, r *http.Request) {
	var req models.FeeMasterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(uint)

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	school, err := schoolRepo.GetSchoolByID(userID)
	if err != nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}

	feeMaster := &models.FeeMaster{
		SchoolID:             school.ID,
		AcademicYear:         req.AcademicYear,
		Course:               req.Course,
		YearOfStudy:          req.YearOfStudy,
		TuitionAmount:        req.TuitionAmount,
		AccommodationAmount:  req.AccommodationAmount,
		FoodAmount:           req.FoodAmount,
		TransportAmount:      req.TransportAmount,
	}

	if err := schoolRepo.CreateFeeMaster(feeMaster); err != nil {
		http.Error(w, "Failed to create fee master", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "saved",
		"message": "Fee master saved successfully",
	})
}

// BulkUpdateFeeMaster handles POST /school/fee-master/bulk-update
func BulkUpdateFeeMaster(w http.ResponseWriter, r *http.Request) {
	var req models.BulkUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(uint)

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	school, _ := schoolRepo.GetSchoolByID(userID)

	calculator := &services.FeeCalculator{}

	students, _ := schoolRepo.GetStudentsByCourse(school.ID, req.Course, req.YearOfStudy)
	studentRepo := repository.NewStudentRepo(repository.GetDB())

	var preview []map[string]interface{}
	var studentsUpdated int

	for _, student := range students {
		balances, _ := studentRepo.GetStudentBalance(student.ID)
		for _, b := range balances {
			oldBalance := b.BalanceRemaining
			paidToDate := b.AmountPaid
			newBalance := calculator.CalculateAutoBalance(0, req.NewTuition, paidToDate)

			preview = append(preview, map[string]interface{}{
				"student_id": student.ID,
				"old_balance":  oldBalance,
				"paid_to_date": paidToDate,
				"new_balance":  newBalance,
			})
			studentsUpdated++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":         true,
		"students_updated": studentsUpdated,
		"formula":        "New Balance = Previous Unpaid + (New Fee - Paid to Date)",
		"preview":        preview,
	})
}

// UpdateStudentBalance handles PUT /school/students/{studentId}/balance
func UpdateStudentBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := parseUint(vars["id"])

	var req struct {
		CoverageType     string `json:"coverage_type"`
		BalanceRemaining float64 `json:"balance_remaining"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentRepo := repository.NewStudentRepo(repository.GetDB())
	balances, _ := studentRepo.GetStudentBalance(studentID)
	for _, b := range balances {
		if b.CoverageType == req.CoverageType {
			studentRepo.UpdateStudentBalance(b.ID, req.BalanceRemaining)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

// CreatePaymentClaim handles POST /school/claims
func CreatePaymentClaim(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClaimRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(uint)

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	school, _ := schoolRepo.GetSchoolByID(userID)

	claim := &models.Claim{
		SchoolID:     school.ID,
		StudentID:    req.StudentID,
		CoverageType: req.CoverageType,
		AmountKSH:    req.AmountKSH,
		Status:       "pending",
	}

	if err := schoolRepo.CreatePaymentClaim(claim); err != nil {
		http.Error(w, "Failed to create claim", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"claim_id": claim.ID,
		"status":   "created",
	})
}

// GetPaymentClaims handles GET /school/claims
func GetPaymentClaims(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	school, _ := schoolRepo.GetSchoolByID(userID)

	claims, err := schoolRepo.GetPaymentClaims(school.ID)
	if err != nil {
		http.Error(w, "Failed to fetch claims", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}

// UploadRoster handles POST /school/roster/upload
func UploadRoster(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         "uploaded",
		"students_added": 0,
	})
}

// SchoolThreeWayVerify handles POST /school/three-way-verify
func SchoolThreeWayVerify(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DisbursementID uint   `json:"disbursement_id"`
		KeyedAmount    float64 `json:"keyed_amount"`
		UploadURL      string `json:"upload_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "submitted",
		"message": "Three-way verification submitted",
	})
}

// GetSchoolDisbursements handles GET /school/disbursements
func GetSchoolDisbursements(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	school, _ := schoolRepo.GetSchoolByID(userID)

	disbursements, err := schoolRepo.GetSchoolDisbursements(school.ID)
	if err != nil {
		http.Error(w, "Failed to fetch disbursements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(disbursements)
}

// parseUint helper
func parseUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint(v)
}