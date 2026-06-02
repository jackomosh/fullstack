package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/altradits/bursaryhub/backend/models"
	"github.com/altradits/bursaryhub/backend/repository"
	"github.com/altradits/bursaryhub/backend/services"
)

// GetAvailableScholarships handles GET /student/scholarships
func GetAvailableScholarships(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	studentRepo := repository.NewStudentRepo(repository.GetDB())
	student, err := studentRepo.GetStudentByID(userID)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	scholarships, err := studentRepo.GetAvailableScholarships(student.Course, student.YearOfStudy, student.GPA)
	if err != nil {
		http.Error(w, "Failed to fetch scholarships", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"scholarships": scholarships,
		"count":        len(scholarships),
	})
}

// ApplyForScholarship handles POST /student/scholarships/{id}/apply
func ApplyForScholarship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scholarshipID := parseUint(vars["id"])

	userID := r.Context().Value("user_id").(uint)

	studentRepo := repository.NewStudentRepo(repository.GetDB())
	student, _ := studentRepo.GetStudentByID(userID)

	application := &models.Application{
		ScholarshipID:     scholarshipID,
		StudentID:         student.ID,
		Status:            "pending",
		AmountRequested:   0,
	}

	if err := studentRepo.CreateApplication(application); err != nil {
		http.Error(w, "Failed to apply for scholarship", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"application_id": application.ID,
		"status":         "applied",
	})
}

// GetStudentBalance handles GET /student/balance
func GetStudentBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	studentRepo := repository.NewStudentRepo(repository.GetDB())
	student, _ := studentRepo.GetStudentByID(userID)

	balances, err := studentRepo.GetStudentBalance(student.ID)
	if err != nil {
		http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balances)
}

// StudentThreeWayVerify handles POST /student/three-way-verify
func StudentThreeWayVerify(w http.ResponseWriter, r *http.Request) {
	var req models.ThreeWayVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// In production, save to database
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"match":             false,
		"student_amount":    req.EnteredAmount,
		"school_amount":     0,
		"fee_master_amount": 0,
		"status":            "pending",
		"message":           "Waiting for school verification",
	})
}

// RequestOTP handles POST /student/claims/{claimId}/request-otp
func RequestOTP(w http.ResponseWriter, r *http.Request) {
	_ = parseUint(mux.Vars(r)["id"]) // claimID used for future validation

	otp := services.GenerateOTP()
	services.SendOTP("", otp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "sent",
		"message":    "OTP sent to your phone",
		"expires_in": 300,
	})
}

// ApproveClaimWithOTP handles POST /student/claims/{claimId}/approve
func ApproveClaimWithOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OTP string `json:"otp"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "approved",
		"message": "Payment approved successfully",
	})
}

// GetStudentDisbursements handles GET /student/disbursements
func GetStudentDisbursements(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	studentRepo := repository.NewStudentRepo(repository.GetDB())
	student, _ := studentRepo.GetStudentByID(userID)

	disbursements, err := studentRepo.GetStudentDisbursements(student.ID)
	if err != nil {
		http.Error(w, "Failed to fetch disbursements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(disbursements)
}