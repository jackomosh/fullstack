package repository

import (
	"database/sql"
	"time"

	"github.com/altradits/bursaryhub/backend/models"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

// GetStudentByID retrieves a student by their user ID
func (r *StudentRepo) GetStudentByID(userID uint) (*models.Student, error) {
	var student models.Student
	var schoolID sql.NullInt32

	err := r.db.QueryRow(`
		SELECT id, user_id, school_id, student_reg_number, course, year_of_study, county, gpa 
		FROM students WHERE user_id = $1
	`, userID).Scan(
		&student.ID, &student.UserID, &schoolID, &student.RegNumber,
		&student.Course, &student.YearOfStudy, &student.County, &student.GPA,
	)
	if err != nil {
		return nil, err
	}
	if schoolID.Valid {
		student.SchoolID = uint(schoolID.Int32)
	}
	return &student, nil
}

// GetStudentBalance retrieves the current balance for a student
func (r *StudentRepo) GetStudentBalance(studentID uint) ([]*models.StudentBalance, error) {
	rows, err := r.db.Query(`
		SELECT id, student_id, academic_year, coverage_type, original_fee, amount_paid, balance_remaining
		FROM student_balances WHERE student_id = $1
	`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var balances []*models.StudentBalance
	for rows.Next() {
		var b models.StudentBalance
		err := rows.Scan(&b.ID, &b.StudentID, &b.AcademicYear, &b.CoverageType, &b.OriginalFee, &b.AmountPaid, &b.BalanceRemaining)
		if err != nil {
			return nil, err
		}
		balances = append(balances, &b)
	}
	return balances, nil
}

// CreateStudentBalance creates a new student balance record
func (r *StudentRepo) CreateStudentBalance(b *models.StudentBalance) error {
	_, err := r.db.Exec(`
		INSERT INTO student_balances (student_id, academic_year, coverage_type, original_fee, amount_paid, balance_remaining)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, b.StudentID, b.AcademicYear, b.CoverageType, b.OriginalFee, b.AmountPaid, b.BalanceRemaining)
	return err
}

// UpdateStudentBalance updates a student's balance
func (r *StudentRepo) UpdateStudentBalance(id uint, balanceRemaining float64) error {
	_, err := r.db.Exec(`
		UPDATE student_balances SET balance_remaining = $1 WHERE id = $2
	`, balanceRemaining, id)
	return err
}

// GetAvailableScholarships retrieves scholarships available to a student
func (r *StudentRepo) GetAvailableScholarships(course string, year int, gpa float64) ([]*models.Scholarship, error) {
	query := `
		SELECT id, donor_id, title, coverage_type, max_amount_per_student, number_of_slots, min_gpa, is_active, application_start_date, application_end_date
		FROM scholarships 
		WHERE is_active = TRUE 
		AND application_end_date >= $1
	`
	
	args := []interface{}{time.Now()}
	
	if course != "" {
		query += " AND eligible_courses @> ARRAY[$2]::text[]"
		args = append(args, course)
	}
	
	if year > 0 {
		query += " AND eligible_years @> ARRAY[$3]::integer[]"
		args = append(args, year)
	}
	
	if gpa > 0 {
		query += " AND min_gpa <= $4"
		args = append(args, gpa)
	}

	rows, err := r.db.Query(query, args...)
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
			s.ApplicationStartDate = sql.NullTime{Time: startDate.Time, Valid: true}
		}
		if endDate.Valid {
			s.ApplicationEndDate = sql.NullTime{Time: endDate.Time, Valid: true}
		}
		scholarships = append(scholarships, &s)
	}
	return scholarships, nil
}

// GetStudentApplications retrieves all applications for a student
func (r *StudentRepo) GetStudentApplications(studentID uint) ([]*models.Application, error) {
	rows, err := r.db.Query(`
		SELECT a.id, a.scholarship_id, a.status, a.amount_requested, a.applied_at, s.title
		FROM applications a
		JOIN scholarships s ON a.scholarship_id = s.id
		WHERE a.student_id = $1
		ORDER BY a.applied_at DESC
	`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []*models.Application
	for rows.Next() {
		var a models.Application
		var appliedAt sql.NullTime
		err := rows.Scan(&a.ID, &a.ScholarshipID, &a.Status, &a.AmountRequested, &appliedAt, &a.ScholarshipTitle)
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

// CreateApplication creates a new scholarship application
func (r *StudentRepo) CreateApplication(app *models.Application) error {
	_, err := r.db.Exec(`
		INSERT INTO applications (scholarship_id, student_id, status, amount_requested)
		VALUES ($1, $2, 'pending', $3)
	`, app.ScholarshipID, app.StudentID, app.AmountRequested)
	return err
}

// GetStudentDisbursements retrieves disbursements for a student
func (r *StudentRepo) GetStudentDisbursements(studentID uint) ([]*models.Disbursement, error) {
	rows, err := r.db.Query(`
		SELECT id, scholarship_id, school_id, amount_usdt, amount_ksh, status, transaction_hash, three_way_match_status, created_at
		FROM disbursements WHERE student_id = $1 ORDER BY created_at DESC
	`, studentID)
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