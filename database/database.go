package database

import (
	"database/sql"
	"fmt"
	"log"

	"project2/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}
	log.Println("Successfully connected to PostgreSQL database!")

	createTable()
}

func CloseDB() {
	DB.Close()
}

func createTable() {
	// People tablosunu oluşturmak için SQL komutu
	createPeopleTableSQL := `CREATE TABLE IF NOT EXISTS people (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`

	// SQL komutunu yürütün ve hataları kontrol edin
	if _, err := DB.Exec(createPeopleTableSQL); err != nil {
		log.Fatalf("Unable to create people table: %v\n", err)
	}
	log.Println("People table created successfully!")

	// Fruits tablosunu oluşturmak için SQL komutu
	createFruitTableSQL := `CREATE TABLE IF NOT EXISTS fruits (
        id SERIAL PRIMARY KEY,
        name TEXT,
        price DECIMAL(10, 2),
        amount INT
    );`

	// SQL komutunu yürütün ve hataları kontrol edin
	if _, err := DB.Exec(createFruitTableSQL); err != nil {
		log.Fatalf("Unable to create fruits table: %v\n", err)
	}
	log.Println("Fruits table created successfully!")
}
