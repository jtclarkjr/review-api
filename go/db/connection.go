package db

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func Connect() {
	// Get the database connection string from the environment variable
	// Locally use export DATABASE_URL=dbUrlHere
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Environment variable DATABASE_URL is not set")
	}

	var err error
	Conn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = Conn.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	log.Println("Database connection established")
}

func SeedDatabase() {
	// Hash passwords for admin and employees
	adminPassword, err := hashPassword("admin")
	if err != nil {
		log.Fatalf("Error hashing admin password: %v", err)
	}
	employeePassword, err := hashPassword("employee")
	if err != nil {
		log.Fatalf("Error hashing employee password: %v", err)
	}

	// Seed the admin user
	_, err = Conn.Exec(`
        INSERT INTO users (email, password, role)
        VALUES ($1, $2, 'admin')
        ON CONFLICT (email) DO NOTHING;
    `, "admin@example.com", adminPassword)
	if err != nil {
		log.Printf("Error seeding admin user: %v", err)
	}

	// Seed employees
	_, err = Conn.Exec(`
        INSERT INTO users (email, password, role)
        VALUES 
        ($1, $2, 'employee'),
        ($3, $2, 'employee')
        ON CONFLICT (email) DO NOTHING;
    `, "employee1@example.com", employeePassword, "employee2@example.com")
	if err != nil {
		log.Printf("Error seeding employee users: %v", err)
	}

	// Seed corresponding employees in the employees table
	_, err = Conn.Exec(`
        INSERT INTO employees (email, position)
        VALUES 
        ($1, 'Developer'),
        ($2, 'Designer')
        ON CONFLICT (email) DO NOTHING;
    `, "employee1@example.com", "employee2@example.com")
	if err != nil {
		log.Printf("Error seeding employees: %v", err)
	}
}

// hashPassword generates a bcrypt hash for the given password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
