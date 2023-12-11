package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "Since2024."
	DBName     = "dbenigma"
)

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"address"`
}

func initDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Define your API routes and handlers here
	router.POST("/customers", createCustomer)

	// Run the server
	port := 8080
	log.Printf("Server is running on :%d...\n", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}

// Add your API handlers here
func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "INSERT INTO customers (id, name, phonenumber, address ) VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	err = stmt.QueryRow(customer.Name, customer.PhoneNumber, customer.Address).Scan(&customer.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert customer into the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully", "data": customer})
}
