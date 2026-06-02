package repository

import (
	"database/sql"

	"github.com/altradits/bursaryhub/backend/models"
)

type DonorRepo struct {
	db *sql.DB
}

func NewDonorRepo(db *sql.DB) *DonorRepo {
	return &DonorRepo{db: db}
}

// GetDonorByID retrieves a donor by their user ID
func (r *DonorRepo) GetDonorByID(userID uint) (*models.Donor, error) {
	var donor models.Donor
	err := r.db.QueryRow(`
		SELECT id, user_id, organization_name, tax_id, kyc_status, total_donated_usd
		FROM donors WHERE user_id = $1
	`, userID).Scan(
		&donor.ID, &donor.UserID, &donor.OrganizationName, &donor.TaxID, &donor.KYCStatus, &donor.TotalDonatedUSD,
	)
	if err != nil {
		return nil, err
	}
	return &donor, nil
}

// CreateScholarship creates a new scholarship
func (r *DonorRepo) CreateScholarship(s *models.Scholarship) error {
	var startDate, endDate interface{}
	if s.ApplicationStartDate.Valid {
		startDate = s.ApplicationStartDate.Time
	} else {
		startDate = nil
	}
	if s.ApplicationEndDate.Valid {
		endDate = s.ApplicationEndDate.Time
	} else {
		endDate = nil
	}

	err := r.db.QueryRow(`
		INSERT INTO scholarships (donor_id, title, coverage_type, max_amount_per_student, number_of_slots, eligible_courses, eligible_years, min_gpa, is_active, application_start_date, application_end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`, s.DonorID, s.Title, s.CoverageType, s.MaxAmountPerStudent, s.NumberOfSlots, s.EligibleCourses, s.EligibleYears, s.MinGPA, s.IsActive, startDate, endDate).Scan(&s.ID)
	return err
}

// GetDonorScholarships retrieves all scholarships for a donor
func (r *DonorRepo) GetDonorScholarships(donorID uint) ([]*models.Scholarship, error) {
	rows, err := r.db.Query(`
		SELECT id, donor_id, title, coverage_type, max_amount_per_student, number_of_slots, min_gpa, is_active, application_start_date, application_end_date
		FROM scholarships WHERE donor_id = $1
		ORDER BY id DESC
	`, donorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scholarships []*models.Scholarship
	for rows.Next() {
		var s models.Scholarship
		var startDate, endDate sql.NullTime
		err := rows.Scan(&s.ID, &s.DonorID, &s.Title, &s.CoverageType, &s.MaxAmountPerStudent, &s.NumberOfSlots, &s.MinGPA, &s.IsActive, &startDate, &endDate)
		if err != nil {
			return nil, err
		}
		if startDate.Valid {
			s.ApplicationStartDate = startDate
		}
		if endDate.Valid {
			s.ApplicationEndDate = endDate
		}
		scholarships = append(scholarships, &s)
	}
	return scholarships, nil
}

// GetScholarshipApplications retrieves applications for a scholarship
func (r *DonorRepo) GetScholarshipApplications(scholarshipID uint) ([]*models.Application, error) {
	rows, err := r.db.Query(`
		SELECT a.id, a.student_id, a.status, a.amount_requested, a.applied_at, u.full_name as student_name, st.course, st.year_of_study
		FROM applications a
		JOIN students st ON a.student_id = st.id
		JOIN users u ON st.user_id = u.id
		WHERE a.scholarship_id = $1
		ORDER BY a.applied_at DESC
	`, scholarshipID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []*models.Application
	for rows.Next() {
		var a models.Application
		var appliedAt sql.NullTime
		err := rows.Scan(&a.ID, &a.StudentID, &a.Status, &a.AmountRequested, &appliedAt, &a.StudentName, &a.StudentCourse, &a.StudentYear)
		if err != nil {
			return nil, err
		}
		if appliedAt.Valid {
			a.AppliedAt = appliedAt.Time
		}
		apps = append(apps, &a)
	}
	return apps, nil
}

// ApproveApplication updates application status
func (r *DonorRepo) ApproveApplication(appID uint, status string) error {
	_, err := r.db.Exec(`
		UPDATE applications SET status = $1 WHERE id = $2
	`, status, appID)
	return err
}

// GetDonorDisbursements retrieves all disbursements for a donor
func (r *DonorRepo) GetDonorDisbursements(donorID uint) ([]*models.Disbursement, error) {
	rows, err := r.db.Query(`
		SELECT d.id, d.scholarship_id, d.school_id, d.amount_usdt, d.amount_ksh, d.status, d.transaction_hash, d.three_way_match_status, d.created_at, s.title, st.school_name
		FROM disbursements d
		JOIN scholarships s ON d.scholarship_id = s.id
		JOIN schools st ON d.school_id = st.id
		WHERE s.donor_id = $1
		ORDER BY d.created_at DESC
	`, donorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disbursements []*models.Disbursement
	for rows.Next() {
		var d models.Disbursement
		var createdAt sql.NullTime
		err := rows.Scan(&d.ID, &d.ScholarshipID, &d.SchoolID, &d.AmountUSDT, &d.AmountKSH, &d.Status, &d.TransactionHash, &d.ThreeWayMatchStatus, &createdAt, &d.ScholarshipTitle, &d.SchoolName)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			d.CreatedAt = createdAt.Time
		}
		disbursements = append(disbursements, &d)
	}
	return disbursements, nil
}

// ApproveDonorKYC updates donor KYC status
func (r *DonorRepo) ApproveDonorKYC(donorID uint, status string) error {
	_, err := r.db.Exec(`
		UPDATE donors SET kyc_status = $1 WHERE id = $2
	`, status, donorID)
	return err
}