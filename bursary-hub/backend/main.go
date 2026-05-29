package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/altradits/bursaryhub/backend/handlers"
	"github.com/altradits/bursaryhub/backend/middleware"
	"github.com/altradits/bursaryhub/backend/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db, err := repository.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// CORS middleware
	r.Use(middleware.CORS)

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Auth routes (no protection)
	r.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")
	r.HandleFunc("/api/auth/verify-otp", handlers.VerifyOTP).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Donor routes
	api.HandleFunc("/donor/scholarships", handlers.CreateScholarship).Methods("POST")
	api.HandleFunc("/donor/scholarships", handlers.GetDonorScholarships).Methods("GET")
	api.HandleFunc("/donor/scholarships/{id}/applications", handlers.GetScholarshipApplications).Methods("GET")
	api.HandleFunc("/donor/scholarships/{id}/applications/{appId}/approve", handlers.ApproveApplication).Methods("POST")
	api.HandleFunc("/donor/disbursements", handlers.GetDonorDisbursements).Methods("GET")
	api.HandleFunc("/donor/impact-report", handlers.GetImpactReport).Methods("GET")
	api.HandleFunc("/donor/cost-breakdown", handlers.GetCostBreakdown).Methods("GET")

	// School routes
	api.HandleFunc("/school/fee-master", handlers.CreateFeeMaster).Methods("POST")
	api.HandleFunc("/school/fee-master/bulk-update", handlers.BulkUpdateFeeMaster).Methods("POST")
	api.HandleFunc("/school/students/{id}/balance", handlers.UpdateStudentBalance).Methods("PUT")
	api.HandleFunc("/school/claims", handlers.CreatePaymentClaim).Methods("POST")
	api.HandleFunc("/school/claims", handlers.GetPaymentClaims).Methods("GET")
	api.HandleFunc("/school/roster/upload", handlers.UploadRoster).Methods("POST")
	api.HandleFunc("/school/three-way-verify", handlers.SchoolThreeWayVerify).Methods("POST")
	api.HandleFunc("/school/disbursements", handlers.GetSchoolDisbursements).Methods("GET")

	// Student routes
	api.HandleFunc("/student/scholarships", handlers.GetAvailableScholarships).Methods("GET")
	api.HandleFunc("/student/scholarships/{id}/apply", handlers.ApplyForScholarship).Methods("POST")
	api.HandleFunc("/student/balance", handlers.GetStudentBalance).Methods("GET")
	api.HandleFunc("/student/three-way-verify", handlers.StudentThreeWayVerify).Methods("POST")
	api.HandleFunc("/student/claims/{id}/approve", handlers.ApproveClaimWithOTP).Methods("POST")
	api.HandleFunc("/student/claims/{id}/request-otp", handlers.RequestOTP).Methods("POST")
	api.HandleFunc("/student/disbursements", handlers.GetStudentDisbursements).Methods("GET")

	// Admin routes
	api.HandleFunc("/admin/schools/{id}/whitelist", handlers.WhitelistSchool).Methods("POST")
	api.HandleFunc("/admin/donors/{id}/kyc", handlers.ApproveDonorKYC).Methods("POST")
	api.HandleFunc("/admin/mismatches", handlers.GetMismatches).Methods("GET")
	api.HandleFunc("/admin/mismatches/{id}/resolve", handlers.ResolveMismatch).Methods("POST")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}