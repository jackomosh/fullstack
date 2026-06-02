.PHONY: all dev build migrate seed setup \
        db-up db-down db-shell db-reset db-logs db-check db-wait \
        test test-api test-db ping ship clean

# ── Config ────────────────────────────────────────────────────
DB_CONTAINER  := bursaryhub_db
DB_USER       := bursary_user
DB_NAME       := bursaryhub
DB_PASSWORD   := bursary_pass
DB_PORT       := 5432
SERVER_PORT   := 8080

# Load .env if it exists
-include .env
export

# ── Dev ───────────────────────────────────────────────────────
dev:
	@echo "🚀 Starting BursaryHub backend on port $(SERVER_PORT)..."
	go run ./backend/main.go

build:
	@echo "🔨 Building binary..."
	@mkdir -p bin
	go build -o bin/bursaryhub ./backend/main.go
	@echo "✅ Binary at bin/bursaryhub"

# ── Database (Docker, no sudo) ────────────────────────────────
db-up:
	@echo "🐘 Starting PostgreSQL container..."
	@docker run -d \
		--name $(DB_CONTAINER) \
		--restart unless-stopped \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		-v bursaryhub_pgdata:/var/lib/postgresql/data \
		postgres:15-alpine 2>/dev/null \
	|| echo "ℹ️  Container already exists — starting if stopped..." \
	&& docker start $(DB_CONTAINER) 2>/dev/null || true
	@$(MAKE) db-wait

db-wait:
	@echo "⏳ Waiting for PostgreSQL to be ready..."
	@until docker exec $(DB_CONTAINER) pg_isready -U $(DB_USER) -d $(DB_NAME) > /dev/null 2>&1; do \
		printf '.'; sleep 1; \
	done
	@echo ""
	@echo "✅ PostgreSQL is ready"

db-down:
	@echo "🛑 Stopping database container..."
	docker stop $(DB_CONTAINER) && docker rm $(DB_CONTAINER)
	@echo "✅ Container stopped and removed"

db-shell:
	@echo "🐚 Opening psql shell (type \q to exit)..."
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

db-logs:
	docker logs -f $(DB_CONTAINER)

db-check:
	@echo "📋 Tables in database:"
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -c "\dt"

db-reset:
	@echo "⚠️  Resetting database (drop + recreate schema)..."
	@docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) \
		-c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" 2>/dev/null || true
	@$(MAKE) migrate
	@$(MAKE) seed
	@echo "✅ Database reset complete"

# ── Migrations (via docker exec — no psql on host needed) ─────
migrate:
	@echo "📐 Running schema migration..."
	@docker exec -i $(DB_CONTAINER) psql \
		-U $(DB_USER) -d $(DB_NAME) \
		< db/schema.sql
	@echo "✅ Schema applied"

seed:
	@echo "🌱 Running seeds..."
	@docker exec -i $(DB_CONTAINER) psql \
		-U $(DB_USER) -d $(DB_NAME) \
		< db/seeds.sql
	@echo "✅ Seeds loaded"

# ── Connection Tests ──────────────────────────────────────────
ping:
	@echo "🔌 Testing database connection..."
	go run cmd/dbtest/main.go

# ── Full Setup (one command from zero to running) ─────────────
setup:
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "  BursaryHub — Full Setup"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@[ -f .env ] || cp .env.example .env
	@$(MAKE) db-up
	@$(MAKE) migrate
	@$(MAKE) seed
	@$(MAKE) ping
	@echo ""
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "  ✅ Setup complete! Run: make dev"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# ── Tests ─────────────────────────────────────────────────────
test-db:
	@echo "🧪 Testing database layer..."
	go test ./backend/repository/... -v -timeout 30s

test-api:
	@echo "🧪 Testing API endpoints..."
	go test ./backend/handlers/... -v -timeout 60s

test:
	@echo "🧪 Running all Go tests..."
	go test ./... -v -timeout 120s
	@echo "✅ All tests passed"

test-contracts:
	@echo "🧪 Running smart contract tests..."
	cd contracts && npx hardhat test

# ── Ship ──────────────────────────────────────────────────────
ship: setup test build
	@echo ""
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "  🚀 BursaryHub is ready to ship!"
	@echo "  Binary: bin/bursaryhub"
	@echo "  Start:  make dev"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# ── Cleanup ───────────────────────────────────────────────────
clean:
	rm -rf bin/
	docker stop $(DB_CONTAINER) 2>/dev/null || true
	docker rm $(DB_CONTAINER) 2>/dev/null || true
	docker volume rm bursaryhub_pgdata 2>/dev/null || true
	@echo "✅ Cleaned up"