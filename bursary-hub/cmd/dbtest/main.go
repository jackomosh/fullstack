package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using env vars")
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

	fmt.Printf("🔌 Testing connection to:\n   %s\n\n", url)

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("❌ Open failed:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Ping failed — is the Docker container running?\n", err)
	}

	fmt.Println("✅ Database connection OK")

	var version string
	db.QueryRow("SELECT version()").Scan(&version)
	fmt.Println("📦 PostgreSQL:", version)

	rows, _ := db.Query(`
		SELECT table_name, 
		       (SELECT COUNT(*) FROM information_schema.columns 
		        WHERE table_name = t.table_name AND table_schema = 'public') as col_count
		FROM information_schema.tables t
		WHERE table_schema = 'public'
		ORDER BY table_name
	`)
	defer rows.Close()

	fmt.Println("\n📋 Schema verification:")
	tableCount := 0
	for rows.Next() {
		var name string
		var cols int
		rows.Scan(&name, &cols)
		fmt.Printf("  ✅ %-35s (%d columns)\n", name, cols)
		tableCount++
	}

	if tableCount < 10 {
		fmt.Printf("\n⚠️  Only %d tables found. Run: make migrate\n", tableCount)
		os.Exit(1)
	}

	fmt.Printf("\n✅ All %d tables verified. Database is ready.\n", tableCount)
}