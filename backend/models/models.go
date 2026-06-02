package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID            uint       `json:"id"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	NationalID    string     `json:"national_id"`
	FullName      string     `json:"full_name"`
	Role          string     `json:"role"`
	PasswordHash  string     `json:"password_hash"`
	IsWhitelisted bool       `json:"is_whitelisted"`
	CreatedAt     time.Time  `json:"created_at"`
}

type Donor struct {
	ID               uint    `json:"id"`
	UserID           uint    `json:"user_id"`
	OrganizationName string  `json:"organization_name"`
	TaxID            string  `json:"tax_id"`
	KYCStatus        string  `json:"kyc_status"`
	TotalDonatedUSD   float64 `json:"total_donated_usd"`
}

type School struct {
	ID                uint   `json:"id"`
	UserID            uint   `json:"user_id"`
	RegistrationNumber string `json:"registration_number"`
	SchoolName        string `json:"school_name"`
	MinistryVerified  bool   `json:"ministry_verified"`
	BankAccountNumber string `json:"bank_account_number"`
	WalletAddress     string `json:"wallet_address"`
	IsWhitelisted     bool   `json:"is_whitelisted"`
}

type Student struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	SchoolID    uint    `json:"school_id"`
	RegNumber   string  `json:"student_reg_number"`
	Course      string  `json:"course"`
	YearOfStudy int     `json:"year_of_study"`
	County      string  `json:"county"`
	GPA         float64 `json:"gpa"`
}

type FeeMaster struct {
	ID                  uint    `json:"id"`
	SchoolID            uint    `json:"school_id"`
	AcademicYear        string  `json:"academic_year"`
	Course              string  `json:"course"`
	YearOfStudy         int     `json:"year_of_study"`
	TuitionAmount       float64 `json:"tuition_amount"`
	AccommodationAmount float64 `json:"accommodation_amount"`
	FoodAmount          float64 `json:"food_amount"`
	TransportAmount     float64 `json:"transport_amount"`
}

type StudentBalance struct {
	ID               uint    `json:"id"`
	StudentID        uint    `json:"student_id"`
	AcademicYear     string  `json:"academic_year"`
	CoverageType     string  `json:"coverage_type"`
	OriginalFee      float64 `json:"original_fee"`
	AmountPaid       float64 `json:"amount_paid"`
	BalanceRemaining float64 `json:"balance_remaining"`
}

type Scholarship struct {
	ID                   uint         `json:"id"`
	DonorID              uint         `json:"donor_id"`
	Title                string       `json:"title"`
	CoverageType         string       `json:"coverage_type"`
	MaxAmountPerStudent  float64      `json:"max_amount_per_student"`
	NumberOfSlots        int          `json:"number_of_slots"`
	EligibleCourses      []string     `json:"eligible_courses"`
	EligibleYears        []int        `json:"eligible_years"`
	MinGPA               float64      `json:"min_gpa"`
	IsActive             bool         `json:"is_active"`
	ApplicationStartDate sql.NullTime `json:"application_start_date"`
	ApplicationEndDate   sql.NullTime `json:"application_end_date"`
}

type Application struct {
	ID              uint         `json:"id"`
	ScholarshipID   uint         `json:"scholarship_id"`
	StudentID       uint         `json:"student_id"`
	Status          string       `json:"status"`
	AmountRequested float64      `json:"amount_requested"`
	AppliedAt       time.Time    `json:"applied_at"`
	ScholarshipTitle string      `json:"scholarship_title"`
	StudentName     string       `json:"student_name"`
	StudentCourse   string       `json:"student_course"`
	StudentYear     int          `json:"student_year"`
}

type Claim struct {
	ID          uint      `json:"id"`
	SchoolID    uint      `json:"school_id"`
	StudentID   uint      `json:"student_id"`
	CoverageType string   `json:"coverage_type"`
	AmountKSH   float64  `json:"amount_ksh"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	StudentName string    `json:"student_name"`
	OTPSent     bool      `json:"otp_sent"`
}

type Disbursement struct {
	ID                 uint         `json:"id"`
	ScholarshipID      uint         `json:"scholarship_id"`
	StudentID          uint         `json:"student_id"`
	SchoolID           uint         `json:"school_id"`
	AmountUSDT         float64      `json:"amount_usdt"`
	AmountKSH          float64      `json:"amount_ksh"`
	Status             string       `json:"status"`
	TransactionHash    string       `json:"transaction_hash"`
	ThreeWayMatchStatus string       `json:"three_way_match_status"`
	CreatedAt          time.Time    `json:"created_at"`
	ScholarshipTitle  string       `json:"scholarship_title,omitempty"`
	SchoolName         string       `json:"school_name,omitempty"`
}

type ThreeWayVerification struct {
	ID                   uint    `json:"id"`
	DisbursementID       uint    `json:"disbursement_id"`
	StudentEnteredAmount float64 `json:"student_entered_amount"`
	SchoolKeyedAmount    float64 `json:"school_keyed_amount"`
	FeeMasterAmount      float64 `json:"fee_master_amount"`
	MatchResult          bool    `json:"match_result"`
	MismatchReason       string  `json:"mismatch_reason"`
}

type TransactionFees struct {
	ID                uint    `json:"id"`
	DisbursementID    uint    `json:"disbursement_id"`
	ConversionInFeeUSD float64 `json:"conversion_in_fee_usd"`
	NetworkGasFeeUSD   float64 `json:"network_gas_fee_usd"`
	ConversionOutFeeKSH float64 `json:"conversion_out_fee_ksh"`
	WithdrawalFeeKSH   float64 `json:"withdrawal_fee_ksh"`
	PlatformFeeKSH     float64 `json:"platform_fee_ksh"`
}

type AuditLog struct {
	ID         uint        `json:"id"`
	UserID     uint        `json:"user_id"`
	Action     string      `json:"action"`
	OldData    interface{} `json:"old_data"`
	NewData    interface{} `json:"new_data"`
	CreatedAt  time.Time   `json:"created_at"`
}

// Request/Response models
type LoginRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

type CreateScholarshipRequest struct {
	Title               string `json:"title"`
	CoverageType        string `json:"coverage_type"`
	MaxAmountPerStudent float64 `json:"max_amount_per_student"`
	NumberOfSlots       int    `json:"number_of_slots"`
	EligibleCourses     []string `json:"eligible_courses"`
	EligibleYears       []int   `json:"eligible_years"`
	MinGPA              float64 `json:"min_gpa"`
	ApplicationStartDate string `json:"application_start_date"`
	ApplicationEndDate   string `json:"application_end_date"`
}

type FeeMasterRequest struct {
	SchoolID            uint    `json:"school_id"`
	AcademicYear        string  `json:"academic_year"`
	Course              string  `json:"course"`
	YearOfStudy         int     `json:"year_of_study"`
	TuitionAmount       float64 `json:"tuition_amount"`
	AccommodationAmount float64 `json:"accommodation_amount"`
	FoodAmount          float64 `json:"food_amount"`
	TransportAmount     float64 `json:"transport_amount"`
}

type BulkUpdateRequest struct {
	SchoolID   uint   `json:"school_id"`
	AcademicYear string `json:"academic_year"`
	Course     string `json:"course"`
	YearOfStudy int    `json:"year_of_study"`
	NewTuition  float64 `json:"new_tuition"`
}

type CreateClaimRequest struct {
	StudentID   uint   `json:"student_id"`
	CoverageType string `json:"coverage_type"`
	AmountKSH   float64 `json:"amount_ksh"`
}

type ThreeWayVerifyRequest struct {
	DisbursementID uint   `json:"disbursement_id"`
	EnteredAmount  float64 `json:"entered_amount"`
	UploadURL      string `json:"upload_url"`
}

type CostBreakdownResponse struct {
	DonorDeposit   float64             `json:"donor_deposit"`
	ConversionToUSDT float64           `json:"conversion_to_usdt"`
	NetworkGasFee  float64             `json:"network_gas_fee"`
	USDTLocked     float64             `json:"usdt_locked"`
	WhenDisbursed DisbursementBreakdown `json:"when_disbursed"`
	TotalFeesUSD   float64             `json:"total_fees_usd"`
	TotalFeesKSH   float64             `json:"total_fees_ksh"`
}

type DisbursementBreakdown struct {
	USDTAmount     float64 `json:"usdt_amount"`
	ConversionToKSH float64 `json:"conversion_to_ksh"`
	WithdrawalFee  float64 `json:"withdrawal_fee_ksh"`
	PlatformFee    float64 `json:"platform_fee_1_percent"`
	SchoolReceives float64 `json:"school_receives_ksh"`
}

type OTPRequest struct {
	ClaimID uint `json:"claim_id"`
}