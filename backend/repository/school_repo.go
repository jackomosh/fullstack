package repository

import (
	"database/sql"

	"github.com/altradits/bursaryhub/backend/models"
)

type SchoolRepo struct {
	db *sql.DB
}

func NewSchoolRepo(db *sql.DB) *SchoolRepo {
	return &SchoolRepo{db: db}
}

// GetSchoolByID retrieves a school by its user ID
func (r *SchoolRepo) GetSchoolByID(userID uint) (*models.School, error) {
	var school models.School
	err := r.db.QueryRow(`
		SELECT id, user_id, registration_number, school_name, ministry_verified, bank_account_number, wallet_address, is_whitelisted
		FROM schools WHERE user_id = $1
	`, userID).Scan(
		&school.ID, &school.UserID, &school.RegistrationNumber, &school.SchoolName,
		&school.MinistryVerified, &school.BankAccountNumber, &school.WalletAddress, &school.IsWhitelisted,
	)
	if err != nil {
		return nil, err
	}
	return &school, nil
}

// CreateFeeMaster creates or updates fee master entry
func (r *SchoolRepo) CreateFeeMaster(f *models.FeeMaster) error {
	_, err := r.db.Exec(`
		INSERT INTO fee_master (school_id, academic_year, course, year_of_study, tuition_amount, accommodation_amount, food_amount, transport_amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (school_id, academic_year, course, year_of_study) 
		DO UPDATE SET 
			tuition_amount = EXCLUDED.tuition_amount,
			accommodation_amount = EXCLUDED.accommodation_amount,
			food_amount = EXCLUDED.food_amount,
			transport_amount = EXCLUDED.transport_amount
	`, f.SchoolID, f.AcademicYear, f.Course, f.YearOfStudy, f.TuitionAmount, f.AccommodationAmount, f.FoodAmount, f.TransportAmount)
	return err
}

// GetFeeMaster retrieves fee master for a school
func (r *SchoolRepo) GetFeeMaster(schoolID uint, academicYear string) ([]*models.FeeMaster, error) {
	rows, err := r.db.Query(`
		SELECT id, school_id, academic_year, course, year_of_study, tuition_amount, accommodation_amount, food_amount, transport_amount
		FROM fee_master WHERE school_id = $1 AND academic_year = $2
	`, schoolID, academicYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fees []*models.FeeMaster
	for rows.Next() {
		var f models.FeeMaster
		err := rows.Scan(&f.ID, &f.SchoolID, &f.AcademicYear, &f.Course, &f.YearOfStudy, &f.TuitionAmount, &f.AccommodationAmount, &f.FoodAmount, &f.TransportAmount)
		if err != nil {
			return nil, err
		}
		fees = append(fees, &f)
	}
	return fees, nil
}

// GetStudentsByCourse retrieves students for a school
func (r *SchoolRepo) GetStudentsByCourse(schoolID uint, course string, year int) ([]*models.Student, error) {
	query := `SELECT id, user_id, school_id, student_reg_number, course, year_of_study, county, gpa FROM students WHERE school_id = $1`
	args := []interface{}{schoolID}

	if course != "" {
		query += " AND course = $2"
		args = append(args, course)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.Student
	for rows.Next() {
		var s models.Student
		var userID sql.NullInt32
		err := rows.Scan(&s.ID, &userID, &s.SchoolID, &s.RegNumber, &s.Course, &s.YearOfStudy, &s.County, &s.GPA)
		if err != nil {
			return nil, err
		}
		if userID.Valid {
			s.UserID = uint(userID.Int32)
		}
		students = append(students, &s)
	}
	return students, nil
}

// CreatePaymentClaim creates a new payment claim
func (r *SchoolRepo) CreatePaymentClaim(c *models.Claim) error {
	_, err := r.db.Exec(`
		INSERT INTO claims (school_id, student_id, coverage_type, amount_ksh, status)
		VALUES ($1, $2, $3, $4, 'pending')
	`, c.SchoolID, c.StudentID, c.CoverageType, c.AmountKSH)
	return err
}

// GetPaymentClaims retrieves all claims for a school
func (r *SchoolRepo) GetPaymentClaims(schoolID uint) ([]*models.Claim, error) {
	rows, err := r.db.Query(`
		SELECT c.id, c.school_id, c.student_id, c.coverage_type, c.amount_ksh, c.status, c.created_at, u.full_name as student_name
		FROM claims c
		JOIN students st ON c.student_id = st.id
		JOIN users u ON st.user_id = u.id
		WHERE c.school_id = $1
		ORDER BY c.created_at DESC
	`, schoolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var claims []*models.Claim
	for rows.Next() {
		var c models.Claim
		var createdAt sql.NullTime
		err := rows.Scan(&c.ID, &c.SchoolID, &c.StudentID, &c.CoverageType, &c.AmountKSH, &c.Status, &createdAt, &c.StudentName)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			c.CreatedAt = createdAt.Time
		}
		claims = append(claims, &c)
	}
	return claims, nil
}

// GetSchoolDisbursements retrieves disbursements for a school
func (r *SchoolRepo) GetSchoolDisbursements(schoolID uint) ([]*models.Disbursement, error) {
	rows, err := r.db.Query(`
		SELECT id, scholarship_id, school_id, amount_usdt, amount_ksh, status, transaction_hash, three_way_match_status, created_at
		FROM disbursements WHERE school_id = $1 ORDER BY created_at DESC
	`, schoolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var disbursements []*models.Disbursement
	for rows.Next() {
		var d models.Disbursement
		var createdAt sql.NullTime
		err := rows.Scan(&d.ID, &d.ScholarshipID, &d.SchoolID, &d.AmountUSDT, &d.AmountKSH, &d.Status, &d.TransactionHash, &d.ThreeWayMatchStatus, &createdAt)
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

// WhitelistSchool updates whitelist status
func (r *SchoolRepo) WhitelistSchool(schoolID uint, status bool) error {
	_, err := r.db.Exec(`
		UPDATE schools SET is_whitelisted = $1 WHERE id = $2
	`, status, schoolID)
	return err
}