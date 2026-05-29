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

// CreateScholarship handles POST /donor/scholarships
func CreateScholarship(w http.ResponseWriter, r *http.Request) {
	var req models.CreateScholarshipRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(uint)

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	donor, err := donorRepo.GetDonorByID(userID)
	if err != nil {
		http.Error(w, "Donor not found", http.StatusNotFound)
		return
	}

	scholarship := &models.Scholarship{
		DonorID:            donor.ID,
		Title:              req.Title,
		CoverageType:       req.CoverageType,
		MaxAmountPerStudent: req.MaxAmountPerStudent,
		NumberOfSlots:      req.NumberOfSlots,
		EligibleCourses:    req.EligibleCourses,
		EligibleYears:      req.EligibleYears,
		MinGPA:             req.MinGPA,
		IsActive:           true,
	}

	if err := donorRepo.CreateScholarship(scholarship); err != nil {
		http.Error(w, "Failed to create scholarship", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"scholarship_id":    scholarship.ID,
		"status":            "created",
		"matching_students": 0,
	})
}

// GetDonorScholarships handles GET /donor/scholarships
func GetDonorScholarships(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	donor, _ := donorRepo.GetDonorByID(userID)
	scholarships, err := donorRepo.GetDonorScholarships(donor.ID)
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

// GetScholarshipApplications handles GET /donor/scholarships/{id}/applications
func GetScholarshipApplications(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scholarshipID := parseUint(vars["id"])

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	apps, err := donorRepo.GetScholarshipApplications(scholarshipID)
	if err != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

// ApproveApplication handles POST /donor/scholarships/{id}/applications/{appId}/approve
func ApproveApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := parseUint(vars["appId"])

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	if err := donorRepo.ApproveApplication(appID, "approved"); err != nil {
		http.Error(w, "Failed to approve application", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "approved"})
}

// GetDonorDisbursements handles GET /donor/disbursements
func GetDonorDisbursements(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	donor, _ := donorRepo.GetDonorByID(userID)
	disbursements, err := donorRepo.GetDonorDisbursements(donor.ID)
	if err != nil {
		http.Error(w, "Failed to fetch disbursements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"disbursements": disbursements,
		"count":         len(disbursements),
	})
}

// GetImpactReport handles GET /donor/impact-report
func GetImpactReport(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	donor, _ := donorRepo.GetDonorByID(userID)
	disbursements, _ := donorRepo.GetDonorDisbursements(donor.ID)
	scholarships, _ := donorRepo.GetDonorScholarships(donor.ID)

	var totalFunded float64
	schoolsMap := make(map[uint]bool)

	for _, d := range disbursements {
		totalFunded += d.AmountKSH
		if _, exists := schoolsMap[d.SchoolID]; !exists {
			schoolsMap[d.SchoolID] = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"total_students_funded": len(disbursements),
		"total_schools_reached": len(schoolsMap),
		"total_ksh_disbursed":   totalFunded,
		"active_scholarships":     len(scholarships),
	})
}

// GetCostBreakdown handles GET /donor/cost-breakdown
func GetCostBreakdown(w http.ResponseWriter, r *http.Request) {
	amount := parseFloat(r.URL.Query().Get("amount"))
	if amount <= 0 {
		amount = 10000
	}

	calculator := &services.FeeCalculator{}
	breakdown := calculator.CalculateCostBreakdown(amount)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breakdown)
}


func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}