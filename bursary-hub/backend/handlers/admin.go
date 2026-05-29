package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/altradits/bursaryhub/backend/repository"
)

// WhitelistSchool handles POST /admin/schools/{id}/whitelist
func WhitelistSchool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schoolID := parseUint(vars["id"])

	schoolRepo := repository.NewSchoolRepo(repository.GetDB())
	if err := schoolRepo.WhitelistSchool(schoolID, true); err != nil {
		http.Error(w, "Failed to whitelist school", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "whitelisted"})
}

// ApproveDonorKYC handles POST /admin/donors/{id}/kyc
func ApproveDonorKYC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	donorID := parseUint(vars["id"])

	donorRepo := repository.NewDonorRepo(repository.GetDB())
	if err := donorRepo.ApproveDonorKYC(donorID, "approved"); err != nil {
		http.Error(w, "Failed to approve KYC", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "kyc_approved"})
}

// GetMismatches handles GET /admin/mismatches
func GetMismatches(w http.ResponseWriter, r *http.Request) {
	adminRepo := repository.NewAdminRepo(repository.GetDB())
	mismatches, err := adminRepo.GetMismatches()
	if err != nil {
		http.Error(w, "Failed to fetch mismatches", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mismatches": mismatches,
		"count":      len(mismatches),
	})
}

// ResolveMismatch handles POST /admin/mismatches/{id}/resolve
func ResolveMismatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mismatchID := parseUint(vars["id"])

	var req struct {
		Resolution string `json:"resolution"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adminRepo := repository.NewAdminRepo(repository.GetDB())
	if err := adminRepo.ResolveMismatch(mismatchID, req.Resolution); err != nil {
		http.Error(w, "Failed to resolve mismatch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "resolved"})
}