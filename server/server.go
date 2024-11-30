package main

import (
	"database/sql"
	"fmt"
	"go-react-blog/db"
	"go-react-blog/web"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := dataSource()
	log.Printf("Connecting to database with: %s", connStr)

	d, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer d.Close()

	// Add a retry mechanism for database connection
	for i := 0; i < 5; i++ {
		err = d.Ping()
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/5): %v", i+1, err)
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		log.Fatal("Failed to connect to database after 5 attempts:", err)
	}

	// Check current database
	var dbname string
	err = d.QueryRow("SELECT current_database()").Scan(&dbname)
	if err != nil {
		log.Fatal("Error checking current database:", err)
	}
	log.Printf("Connected to database: %s", dbname)

	// Call verifyDatabaseSetup to ensure tables and data exist
	if err := verifyDatabaseSetup(d); err != nil {
		log.Fatal("Database setup verification failed:", err)
	}

	app := web.NewApp(db.NewDB(d), true)
	err = app.Serve()
	log.Fatal("Server error:", err)
}

func dataSource() string {
	host := "localhost"
	pass := "mysecretpassword"
	dbname := "postgres"
	user := "postgres"
	if os.Getenv("profile") == "prod" {
		host = "db"
		pass = os.Getenv("db_pass")
		dbname = "goxygen"
	}
	return fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s search_path=public sslmode=disable",
		host, user, pass, dbname,
	)
}

func verifyDatabaseSetup(db *sql.DB) error {
	// Check if tables exist and are accessible
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name IN ('blogs', 'technologies')
	`).Scan(&count)

	if err != nil {
		return fmt.Errorf("error checking tables: %v", err)
	}

	if count != 2 {
		// Try to create tables if they don't exist
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS public.technologies (
				name    VARCHAR(255) PRIMARY KEY,
				details VARCHAR(255)
			);

			CREATE TABLE IF NOT EXISTS public.blogs (
				id SERIAL PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				coverURL VARCHAR(255),
				body TEXT NOT NULL
			);
		`)
		if err != nil {
			return fmt.Errorf("error creating tables: %v", err)
		}
	}

	// Check if data exists
	var blogCount, techCount int
	err = db.QueryRow("SELECT COUNT(*) FROM public.blogs").Scan(&blogCount)
	if err != nil {
		return fmt.Errorf("error checking blogs count: %v", err)
	}

	err = db.QueryRow("SELECT COUNT(*) FROM public.technologies").Scan(&techCount)
	if err != nil {
		return fmt.Errorf("error checking technologies count: %v", err)
	}

	// Insert sample data if tables are empty
	if blogCount == 0 {
		_, err = db.Exec(`
			INSERT INTO public.blogs (title, coverURL, body) VALUES 
			('First Blog', 'https://img.freepik.com/free-photo/networking-media-sharing-icons-graphic-concept_53876-120836.jpg?t=st=1733001062~exp=1733004662~hmac=9f1a3e07b05f357ad985863ba753c78071db2591c42aada40a82c1eaba545883&w=740', 'This is the first blog post'),
			('Second Blog', 'https://img.freepik.com/free-vector/blogging-concept-with-man_23-2148653963.jpg?t=st=1733001110~exp=1733004710~hmac=625f20845479d3b915f5aabba67e35dcc804c4b8ef399cbdc12fe7034e6721c9&w=740', 'This is the second blog post')
		`)
		if err != nil {
			return fmt.Errorf("error inserting sample blogs: %v", err)
		}
		log.Println("Inserted sample blogs")
	} else {
		log.Println("Blogs already exist, skipping insertion")
	}

	if techCount == 0 {
		_, err = db.Exec(`
			INSERT INTO public.technologies (name, details) VALUES 
			('Go', 'An open source programming language that makes it easy to build simple and efficient software.'),
			('JavaScript', 'A lightweight, interpreted, or just-in-time compiled programming language with first-class functions.'),
			('PostgreSQL', 'A powerful, open source object-relational database system')
		`)
		if err != nil {
			return fmt.Errorf("error inserting sample technologies: %v", err)
		}
		log.Println("Inserted sample technologies")
	} else {
		log.Println("Technologies already exist, skipping insertion")
	}

	return nil
}
