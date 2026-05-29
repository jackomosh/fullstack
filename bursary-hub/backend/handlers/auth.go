package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/altradits/bursaryhub/backend/models"
	"github.com/altradits/bursaryhub/backend/repository"
	"github.com/altradits/bursaryhub/backend/services"
)

// Login handles POST /auth/login - sends OTP to user
func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adminRepo := repository.NewAdminRepo(repository.GetDB())
	user, err := adminRepo.GetUserByPhone(req.Phone)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Generate OTP
	otp := services.GenerateOTP()
	services.SendOTP(req.Phone, otp)
	services.StoreOTP(req.Phone, otp, user.ID, user.Role)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP sent successfully",
		"status":  "pending",
	})
}

// VerifyOTP handles POST /auth/verify-otp - verifies OTP and returns JWT
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req models.VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, valid := services.VerifyOTP(req.Phone, req.OTP)
	if !valid {
		http.Error(w, "Invalid or expired OTP", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		secret = []byte("bursaryhub-dev-secret-change-in-production")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": record.UserID,
		"role":    record.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user": map[string]interface{}{
			"id":  record.UserID,
			"role": record.Role,
		},
	})
}