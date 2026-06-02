package repository

import (
	"database/sql"

	"github.com/altradits/bursaryhub/backend/models"
)

type AdminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

// GetMismatches retrieves all mismatched three-way verifications
func (r *AdminRepo) GetMismatches() ([]*models.ThreeWayVerification, error) {
	rows, err := r.db.Query(`
		SELECT id, disbursement_id, student_entered_amount, school_keyed_amount, fee_master_amount, match_result, mismatch_reason
		FROM three_way_verification WHERE match_result = FALSE
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mismatches []*models.ThreeWayVerification
	for rows.Next() {
		var m models.ThreeWayVerification
		err := rows.Scan(&m.ID, &m.DisbursementID, &m.StudentEnteredAmount, &m.SchoolKeyedAmount, &m.FeeMasterAmount, &m.MatchResult, &m.MismatchReason)
		if err != nil {
			return nil, err
		}
		mismatches = append(mismatches, &m)
	}
	return mismatches, nil
}

// ResolveMismatch marks a mismatch as resolved
func (r *AdminRepo) ResolveMismatch(id uint, resolution string) error {
	_, err := r.db.Exec(`
		UPDATE three_way_verification SET mismatch_reason = $1, match_result = TRUE WHERE id = $2
	`, resolution, id)
	return err
}

// GetUserByPhone retrieves a user by phone number
func (r *AdminRepo) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, email, phone, national_id, full_name, role, password_hash, is_whitelisted, created_at
		FROM users WHERE phone = $1
	`, phone).Scan(
		&user.ID, &user.Email, &user.Phone, &user.NationalID, &user.FullName, &user.Role, &user.PasswordHash, &user.IsWhitelisted, &user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}