-- Seed 5 mock Kenyan universities (whitelisted)
INSERT INTO users (email, phone, national_id, full_name, role, password_hash, is_whitelisted) VALUES
('admin@uonbi.ac.ke', '+254700000001', 'UNIV001', 'University of Nairobi Admin', 'school_admin', '$2a$10$hash1', TRUE),
('admin@jkuit.ac.ke', '+254700000002', 'UNIV002', 'JKUAT Admin', 'school_admin', '$2a$10$hash2', TRUE),
('admin@maseno.ac.ke', '+254700000003', 'UNIV003', 'Maseno University Admin', 'school_admin', '$2a$10$hash3', TRUE),
('admin@egerton.ac.ke', '+254700000004', 'UNIV004', 'Egerton University Admin', 'school_admin', '$2a$10$hash4', TRUE),
('admin@uonbi2.ac.ke', '+254700000005', 'UNIV005', 'Kenyatta University Admin', 'school_admin', '$2a$10$hash5', TRUE);

INSERT INTO schools (user_id, registration_number, school_name, ministry_verified, bank_account_number, wallet_address, is_whitelisted) VALUES
(1, 'EDU-KE-2024-001', 'University of Nairobi', TRUE, '0123456789', '0x742d35Cc6634C0532925a3b8D7389D75cF8b8F9C', TRUE),
(2, 'EDU-KE-2024-002', 'JKUAT', TRUE, '0234567891', '0x853d35Cc6634C0532925a3b8D7389D75cF8b8F9D', TRUE),
(3, 'EDU-KE-2024-003', 'Maseno University', TRUE, '0345678912', '0x964d35Cc6634C0532925a3b8D7389D75cF8b8F9E', TRUE),
(4, 'EDU-KE-2024-004', 'Egerton University', TRUE, '0456789123', '0xa75d35Cc6634C0532925a3b8D7389D75cF8b8F9F', TRUE),
(5, 'EDU-KE-2024-005', 'Kenyatta University', TRUE, '0567891234', '0xb86d35Cc6634C0532925a3b8D7389D75cF8b8F90', TRUE);

-- Seed 3 mock donors
INSERT INTO users (email, phone, national_id, full_name, role, password_hash, is_whitelisted) VALUES
('donor1@foundation.or.ke', '+254711000001', 'DONOR001', 'Safaricom Foundation', 'donor', '$2a$10$hashd1', TRUE),
('donor2@trust.or.ke', '+254711000002', 'DONOR002', 'Equity Bank Trust', 'donor', '$2a$10$hashd2', TRUE),
('donor3@ngo.or.ke', '+254711000003', 'DONOR003', 'Tech Innovators NGO', 'donor', '$2a$10$hashd3', TRUE);

INSERT INTO donors (user_id, organization_name, tax_id, kyc_status, total_donated_usd) VALUES
(6, 'Safaricom Foundation', 'TAX-001', 'approved', 247500),
(7, 'Equity Bank Trust', 'TAX-002', 'approved', 185000),
(8, 'Tech Innovators NGO', 'TAX-003', 'approved', 95000);

-- Seed 20 mock students across different courses and years
INSERT INTO users (email, phone, national_id, full_name, role, password_hash, is_whitelisted) VALUES
('student1@uonbi.ac.ke', '+254720000001', 'STUD001', 'Brian Otieno', 'student', '$2a$10$hashs1', TRUE),
('student2@uonbi.ac.ke', '+254720000002', 'STUD002', 'Mercy Akinyi', 'student', '$2a$10$hashs2', TRUE),
('student3@jkuit.ac.ke', '+254720000003', 'STUD003', 'John Kimani', 'student', '$2a$10$hashs3', TRUE),
('student4@jkuit.ac.ke', '+254720000004', 'STUD004', 'Sarah Wambui', 'student', '$2a$10$hashs4', TRUE),
('student5@maseno.ac.ke', '+254720000005', 'STUD005', 'Peter Mwangi', 'student', '$2a$10$hashs5', TRUE),
('student6@maseno.ac.ke', '+254720000006', 'STUD006', 'Ann Njeri', 'student', '$2a$10$hashs6', TRUE),
('student7@egerton.ac.ke', '+254720000007', 'STUD007', 'David Ochieng', 'student', '$2a$10$hashs7', TRUE),
('student8@egerton.ac.ke', '+254720000008', 'STUD008', 'Grace Wanjiru', 'student', '$2a$10$hashs8', TRUE),
('student9@uonbi.ac.ke', '+254720000009', 'STUD009', 'James Mutua', 'student', '$2a$10$hashs9', TRUE),
('student10@uonbi.ac.ke', '+254720000010', 'STUD010', 'Mary Auma', 'student', '$2a$10$hashs10', TRUE),
('student11@jkuit.ac.ke', '+254720000011', 'STUD011', 'Samuel Kiprop', 'student', '$2a$10$hashs11', TRUE),
('student12@jkuit.ac.ke', '+254720000012', 'STUD012', 'Lucy Chebet', 'student', '$2a$10$hashs12', TRUE),
('student13@maseno.ac.ke', '+254720000013', 'STUD013', 'Michael Omondi', 'student', '$2a$10$hashs13', TRUE),
('student14@maseno.ac.ke', '+254720000014', 'STUD014', 'Elizabeth Adhiambo', 'student', '$2a$10$hashs14', TRUE),
('student15@egerton.ac.ke', '+254720000015', 'STUD015', 'Robert Maina', 'student', '$2a$10$hashs15', TRUE),
('student16@egerton.ac.ke', '+254720000016', 'STUD016', 'Cynthia Achieng', 'student', '$2a$10$hashs16', TRUE),
('student17@uonbi.ac.ke', '+254720000017', 'STUD017', 'Joseph Kamau', 'student', '$2a$10$hashs17', TRUE),
('student18@uonbi.ac.ke', '+254720000018', 'STUD018', 'Jane Wambui', 'student', '$2a$10$hashs18', TRUE),
('student19@jkuit.ac.ke', '+254720000019', 'STUD019', 'Daniel Kiprono', 'student', '$2a$10$hashs19', TRUE),
('student20@jkuit.ac.ke', '+254720000020', 'STUD020', 'Faith Chelangat', 'student', '$2a$10$hashs20', TRUE);

INSERT INTO students (user_id, school_id, student_reg_number, course, year_of_study, county, gpa) VALUES
(9, 1, 'CS001', 'Computer Science', 1, 'Nairobi', 3.8),
(10, 1, 'IT002', 'Information Technology', 2, 'Kisumu', 3.6),
(11, 2, 'CS003', 'Computer Science', 1, 'Nakuru', 3.7),
(12, 2, 'EN004', 'Engineering', 3, 'Eldoret', 3.5),
(13, 3, 'BS005', 'Business', 2, 'Kakamega', 3.4),
(14, 3, 'NU006', 'Nursing', 1, 'Kitale', 3.9),
(15, 4, 'ED007', 'Education', 3, 'Nanyuki', 3.6),
(16, 4, 'LA008', 'Law', 2, 'Mombasa', 3.7),
(17, 1, 'CS009', 'Computer Science', 2, 'Nairobi', 3.5),
(18, 1, 'IT010', 'Information Technology', 1, 'Nakuru', 3.8),
(19, 2, 'CS011', 'Computer Science', 3, 'Eldoret', 3.6),
(20, 2, 'EN012', 'Engineering', 1, 'Kisumu', 3.4),
(21, 3, 'BS013', 'Business', 1, 'Kakamega', 3.7),
(22, 3, 'NU014', 'Nursing', 2, 'Kitale', 3.8),
(23, 4, 'ED015', 'Education', 1, 'Nairobi', 3.5),
(24, 4, 'LA016', 'Law', 4, 'Mombasa', 3.9),
(25, 5, 'CS017', 'Computer Science', 3, 'Nairobi', 3.6),
(26, 5, 'IT018', 'Information Technology', 2, 'Nakuru', 3.7),
(27, 5, 'EN019', 'Engineering', 4, 'Eldoret', 3.5),
(28, 5, 'BS020', 'Business', 3, 'Kakamega', 3.8);

-- Sample fee master entries
INSERT INTO fee_master (school_id, academic_year, course, year_of_study, tuition_amount, accommodation_amount, food_amount, transport_amount) VALUES
(1, '2024-2025', 'Computer Science', 1, 55000, 25000, 15000, 5000),
(1, '2024-2025', 'Information Technology', 1, 50000, 20000, 15000, 5000),
(2, '2024-2025', 'Computer Science', 1, 45000, 20000, 12000, 4000),
(2, '2024-2025', 'Engineering', 3, 60000, 15000, 12000, 4000),
(3, '2024-2025', 'Business', 2, 40000, 18000, 10000, 3000),
(3, '2024-2025', 'Nursing', 1, 35000, 15000, 8000, 3000),
(4, '2024-2025', 'Education', 3, 30000, 12000, 8000, 2500),
(4, '2024-2025', 'Law', 2, 65000, 18000, 10000, 3500),
(5, '2024-2025', 'Computer Science', 3, 52000, 22000, 14000, 4500),
(5, '2024-2025', 'Engineering', 4, 58000, 16000, 12000, 4000);

-- Sample student balances
INSERT INTO student_balances (student_id, academic_year, coverage_type, original_fee, amount_paid, balance_remaining) VALUES
(1, '2024-2025', 'tuition', 55000, 0, 55000),
(2, '2024-2025', 'tuition', 50000, 0, 50000),
(3, '2024-2025', 'tuition', 45000, 0, 45000),
(4, '2024-2025', 'tuition', 60000, 0, 60000),
(5, '2024-2025', 'tuition', 40000, 0, 40000);

-- Sample scholarships
INSERT INTO scholarships (donor_id, title, coverage_type, max_amount_per_student, number_of_slots, eligible_courses, eligible_years, min_gpa, is_active, application_start_date, application_end_date) VALUES
(1, 'Tech Leaders Fund', 'tuition', 120000, 50, ARRAY['Computer Science', 'Information Technology'], ARRAY[1, 2, 3], 3.5, TRUE, '2024-01-01', '2024-12-31'),
(2, 'Women in STEM', 'all', 200000, 30, ARRAY['Computer Science', 'Engineering', 'Nursing'], ARRAY[1, 2], 3.0, TRUE, '2024-02-01', '2024-11-30'),
(3, 'Business Excellence', 'tuition', 80000, 25, ARRAY['Business', 'Education'], ARRAY[2, 3, 4], 3.2, TRUE, '2024-03-01', '2024-10-31');