package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, using environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	fmt.Println("🔌 Connecting to database...")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("❌ Failed to open DB:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Connection failed:", err)
	}
	fmt.Println("✅ Connected successfully")

	files := []string{
		"db/schema.sql",
		"db/seeds.sql",
	}

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			log.Fatalf("❌ Cannot read %s: %v", f, err)
		}

		// Split on semicolons and run statement by statement
		stmts := strings.Split(string(content), ";")
		for _, stmt := range stmts {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			if _, err := db.Exec(stmt); err != nil {
				// Log but don't fatal on "already exists" errors
				if strings.Contains(err.Error(), "already exists") {
					continue
				}
				log.Printf("⚠️  Warning in %s: %v", f, err)
			}
		}
		fmt.Printf("✅ Applied: %s\n", f)
	}

	// Verify tables
	rows, err := db.Query(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public'
		ORDER BY table_name
	`)
	if err != nil {
		log.Fatal("❌ Table check failed:", err)
	}

	fmt.Println("\n📋 Tables in database:")
	count := 0
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf("  ✅ %s\n", name)
		count++
	}
	fmt.Printf("\n✅ Total: %d tables\n", count)
}