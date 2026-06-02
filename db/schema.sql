-- ENUM TYPES
CREATE TYPE user_role AS ENUM ('donor', 'school_admin', 'student', 'admin');
CREATE TYPE coverage_type AS ENUM ('tuition', 'accommodation', 'food', 'transport', 'all', 'unrestricted');
CREATE TYPE disbursement_status AS ENUM ('pending', 'verified', 'completed', 'failed', 'blocked');
CREATE TYPE scholarship_status AS ENUM ('active', 'closed', 'expired');
CREATE TYPE application_status AS ENUM ('pending', 'approved', 'rejected');

-- USERS TABLE
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    national_id VARCHAR(20) UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_whitelisted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- DONORS TABLE
CREATE TABLE donors (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    organization_name VARCHAR(255),
    tax_id VARCHAR(100),
    kyc_status VARCHAR(50) DEFAULT 'pending',
    total_donated_usd DECIMAL(20,2) DEFAULT 0
);

-- SCHOOLS TABLE
CREATE TABLE schools (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    registration_number VARCHAR(100) UNIQUE NOT NULL,
    school_name VARCHAR(255) NOT NULL,
    ministry_verified BOOLEAN DEFAULT FALSE,
    bank_account_number VARCHAR(100),
    wallet_address VARCHAR(255) UNIQUE,
    is_whitelisted BOOLEAN DEFAULT FALSE
);

-- STUDENTS TABLE
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    school_id INTEGER REFERENCES schools(id),
    student_reg_number VARCHAR(100) UNIQUE NOT NULL,
    course VARCHAR(255),
    year_of_study INTEGER,
    county VARCHAR(100),
    gpa DECIMAL(3,2)
);

-- FEE MASTER TABLE
CREATE TABLE fee_master (
    id SERIAL PRIMARY KEY,
    school_id INTEGER REFERENCES schools(id) ON DELETE CASCADE,
    academic_year VARCHAR(20) NOT NULL,
    course VARCHAR(255),
    year_of_study INTEGER,
    tuition_amount DECIMAL(15,2) DEFAULT 0,
    accommodation_amount DECIMAL(15,2) DEFAULT 0,
    food_amount DECIMAL(15,2) DEFAULT 0,
    transport_amount DECIMAL(15,2) DEFAULT 0,
    UNIQUE(school_id, academic_year, course, year_of_study)
);

-- STUDENT BALANCES TABLE
CREATE TABLE student_balances (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES students(id) ON DELETE CASCADE,
    academic_year VARCHAR(20),
    coverage_type coverage_type NOT NULL,
    original_fee DECIMAL(15,2) NOT NULL,
    amount_paid DECIMAL(15,2) DEFAULT 0,
    balance_remaining DECIMAL(15,2) NOT NULL
);

-- SCHOLARSHIPS TABLE
CREATE TABLE scholarships (
    id SERIAL PRIMARY KEY,
    donor_id INTEGER REFERENCES donors(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    coverage_type coverage_type NOT NULL,
    max_amount_per_student DECIMAL(15,2),
    number_of_slots INTEGER,
    eligible_courses TEXT[],
    eligible_years INTEGER[],
    min_gpa DECIMAL(3,2),
    is_active BOOLEAN DEFAULT TRUE,
    application_start_date DATE,
    application_end_date DATE
);

-- APPLICATIONS TABLE
CREATE TABLE applications (
    id SERIAL PRIMARY KEY,
    scholarship_id INTEGER REFERENCES scholarships(id),
    student_id INTEGER REFERENCES students(id),
    status application_status DEFAULT 'pending',
    amount_requested DECIMAL(15,2),
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- DISBURSEMENTS TABLE
CREATE TABLE disbursements (
    id SERIAL PRIMARY KEY,
    scholarship_id INTEGER REFERENCES scholarships(id),
    student_id INTEGER REFERENCES students(id),
    school_id INTEGER REFERENCES schools(id),
    amount_usdt DECIMAL(20,8) NOT NULL,
    amount_ksh DECIMAL(15,2) NOT NULL,
    status disbursement_status DEFAULT 'pending',
    transaction_hash VARCHAR(255),
    three_way_match_status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- THREE-WAY VERIFICATION LOGS
CREATE TABLE three_way_verification (
    id SERIAL PRIMARY KEY,
    disbursement_id INTEGER REFERENCES disbursements(id),
    student_entered_amount DECIMAL(15,2),
    school_keyed_amount DECIMAL(15,2),
    fee_master_amount DECIMAL(15,2),
    match_result BOOLEAN,
    mismatch_reason TEXT
);

-- TRANSACTION FEES LOG
CREATE TABLE transaction_fees (
    id SERIAL PRIMARY KEY,
    disbursement_id INTEGER REFERENCES disbursements(id),
    conversion_in_fee_usd DECIMAL(10,4),
    network_gas_fee_usd DECIMAL(10,4),
    conversion_out_fee_ksh DECIMAL(10,2),
    withdrawal_fee_ksh DECIMAL(10,2),
    platform_fee_ksh DECIMAL(10,2)
);

-- AUDIT LOGS
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    action VARCHAR(255) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INDEXES
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_students_school_id ON students(school_id);
CREATE INDEX idx_fee_master_school_year ON fee_master(school_id, academic_year);
CREATE INDEX idx_disbursements_status ON disbursements(status);
CREATE INDEX idx_scholarships_donor_id ON scholarships(donor_id);
CREATE INDEX idx_applications_student_id ON applications(student_id);
CREATE INDEX idx_applications_scholarship_id ON applications(scholarship_id);