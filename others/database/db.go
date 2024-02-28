package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
)

const alphanumericChars = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphanumericChars[rand.Intn(len(alphanumericChars))]
	}
	return string(result)
}

func DbInsert(db *sql.DB, code string) string {
	// Create the database table codes with the following schema (id will be 6 characters long and alphanumeric):
	// code string, id string
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS codes (id VARCHAR(6) PRIMARY KEY, code TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully.")
	// insert the code into the database
	id := generateRandomString(6)
	_, err = db.Exec("INSERT INTO codes (id, code) VALUES ($1, $2)", id, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Code inserted successfully With ID: ", id)
	// close the connection
	defer db.Close()
	// return the id
	return id

}

func DbGet(db *sql.DB, id string) string {
	// get the code from the database
	var code string
	err := db.QueryRow("SELECT code FROM codes WHERE id = $1", id).Scan(&code)
	if err != nil {
		return "Code Ain't Found"
	}
	fmt.Println("Code retrieved successfully")
	// close the connection
	defer db.Close()
	// return the code
	return code
}
