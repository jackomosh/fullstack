# BursaryHub

**Version:** 1.0.0 | **Last Updated:** 2026-05-29

Free, fraud-proof scholarship disbursement platform for Kenya that connects donors, schools, and students on a single transparent system. Money moves only after all parties confirm the same amount through three-way verification.

## What It Does

- **Zero subscription fees** for schools and students
- **1% success fee** taken only when money is disbursed
- **Three-way verification** ensures student, school, and fee master amounts match
- **Auto-balance calculation** eliminates manual fee updates
- **SMS OTP approval** works on any phone, no smartphone required

## Tech Stack

| Layer | Technology | Version |
|-------|------------|---------|
| Language | Go | 1.21 |
| HTTP Router | Gorilla Mux | v1.8.1 |
| Database Driver | lib/pq | v1.10.9 |
| Auth | JWT (golang-jwt) | v5.0.0 |
| Database | PostgreSQL | 15+ (Alpine) |
| Blockchain | go-ethereum | v1.13.0 |
| Smart Contracts | Solidity | ^0.8.19 |
| Frontend | Tailwind CSS + Vanilla JS | CDN |

## Project Structure

```
bursaryhub/
├── backend/           # Go HTTP API server
│   ├── main.go        # Entry point (port 8080)
│   ├── handlers/      # HTTP handlers for all endpoints
│   ├── middleware/    # JWT auth middleware
│   ├── repository/    # Database queries
│   └── services/      # Business logic (OTP, fees, blockchain)
├── db/                # Database migrations and seeds
│   ├── schema.sql     # 12 tables with ENUMs and constraints
│   └── seeds.sql      # 28 users, 5 schools, 20 students
├── frontend/          # Three dashboard clients
│   ├── index.html     # Landing page
│   ├── grantor/       # Donor dashboard
│   ├── institution/   # School dashboard
│   └── beneficiary/   # Student portal
├── contracts/         # Solidity smart contracts
│   ├── BursaryEscrow.sol
│   └── VendorRegistry.sol
└── scripts/           # Test scripts
    └── test-api.sh
```

## Prerequisites

- Go 1.21+
- Docker 27.3+ (rootless mode supported)
- Node.js 18+ (for smart contract compilation)
- Python 3+ (for frontend server)

## Quick Start

### 1. Clone & Enter

```bash
git clone https://github.com/altradits/bursaryhub.git
cd bursaryhub
```

### 2. Copy Env

```bash
cp .env.example .env
```

### 3. Start Database (Docker, no sudo)

```bash
make db-up
```

### 4. Run Migrations

```bash
make migrate
```

### 5. Seed Data

```bash
make seed
```

### 6. Verify Database

```bash
make db-check
```

### 7. Start Backend

```bash
setsid env GOPATH=~/go go run ./backend/main.go
```

### 8. Serve Frontend

```bash
cd frontend && python3 -m http.server 3000
```

## Environment Variables

| Variable | Description | Example | Required |
|----------|-------------|---------|----------|
| DATABASE_URL | PostgreSQL connection string | postgres://user:pass@localhost:5432/bursaryhub | Yes |
| JWT_SECRET | JWT signing key | bursaryhub-dev-secret-change-in-production | Yes |
| MPESA_CONSUMER_KEY | M-Pesa API consumer key | sandbox | No |
| MPESA_CONSUMER_SECRET | M-Pesa API consumer secret | sandbox | No |
| MPESA_SHORTCODE | M-Pesa shortcode | 174379 | No |
| MPESA_PASSKEY | M-Pesa passkey | dev-stub | No |
| MPESA_CALLBACK_URL | M-Pesa callback URL | http://localhost:8080/mpesa/callback | No |
| AFRICAS_TALKING_API_KEY | SMS provider API key | dev-stub-key | No |
| AFRICAS_TALKING_USERNAME | SMS provider username | sandbox | No |
| INFURA_PROJECT_ID | Blockchain RPC endpoint | dev-stub | No |
| ESCROW_CONTRACT_ADDRESS | Deployed contract address | 0x... | No |
| USDT_CONTRACT_ADDRESS | USDT token address | 0x... | No |
| PLATFORM_WALLET | Platform fee wallet | 0x... | No |
| PORT | Server port | 8080 | No |

## API Overview

### Auth
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/auth/login | Request OTP via SMS |
| POST | /api/auth/verify-otp | Verify OTP, receive JWT |

### Donor
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/donor/scholarships | Create new scholarship |
| GET | /api/donor/scholarships | List donor's scholarships |
| GET | /api/donor/scholarships/{id}/applications | View applications |
| POST | /api/donor/scholarships/{id}/applications/{appId}/approve | Approve application |
| GET | /api/donor/disbursements | View disbursements |
| GET | /api/donor/cost-breakdown | Preview all fees |

### School
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/school/fee-master | Publish fee structure |
| POST | /api/school/fee-master/bulk-update | Bulk update fees |
| PUT | /api/school/students/{id}/balance | Update student balance |
| POST | /api/school/claims | Create payment claim |
| GET | /api/school/claims | List payment claims |
| POST | /api/school/three-way-verify | Submit school verification |
| GET | /api/school/disbursements | View school disbursements |

### Student
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/student/scholarships | Get available scholarships |
| POST | /api/student/scholarships/{id}/apply | Apply for scholarship |
| GET | /api/student/balance | Get student balance |
| POST | /api/student/three-way-verify | Submit balance verification |
| POST | /api/student/claims/{id}/request-otp | Request SMS OTP |
| POST | /api/student/claims/{id}/approve | Approve claim with OTP |
| GET | /api/student/disbursements | View student disbursements |

### Admin
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/admin/schools/{id}/whitelist | Whitelist school |
| POST | /api/admin/donors/{id}/kyc | Approve donor KYC |
| GET | /api/admin/mismatches | Get mismatched verifications |
| POST | /api/admin/mismatches/{id}/resolve | Resolve mismatch |

## Running Tests

```bash
# Database connection test
make ping

# Full API test suite (requires server running)
bash scripts/test-api.sh

# Go unit tests
make test

# Smart contract tests
make test-contracts
```

## Smart Contracts

Deployed to local Hardhat network for testing.

```bash
# Install dependencies
cd contracts && npm install

# Run tests
npx hardhat test

# Deploy (local)
npx hardhat run scripts/deploy.js
```

**Contracts:**
- `BursaryEscrow.sol` - Core escrow with 1% platform fee
- `VendorRegistry.sol` - School/donor whitelist registry

## Team Setup

For machines without admin access:

1. Install rootless Docker: `dockerd-rootless-setuptool.sh install`
2. No sudo required - all database operations via `docker exec`
3. Use `make db-up` to start PostgreSQL container
4. Use `make migrate` and `make seed` to initialize database
5. Frontend runs on any port with `python3 -m http.server`

## License

MIT License