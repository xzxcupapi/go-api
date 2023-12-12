package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-api/handlers"

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

	//Customers Router
	router.POST("/customers", func(c *gin.Context) {
		handlers.CreateCustomer(c, db)
	})
	router.GET("/customers/:id", func(c *gin.Context) {
		handlers.GetCustomer(c, db)
	})
	router.PUT("/customers/:id", func(c *gin.Context) {
		handlers.UpdateCustomer(c, db)
	})
	router.DELETE("/customers/:id", func(c *gin.Context) {
		handlers.DeleteCustomer(c, db)
	})

	//Products Router
	router.POST("/products", func(c *gin.Context) {
		handlers.CreateProducts(c, db)
	})
	router.GET("/products/:id", func(c *gin.Context) {
		handlers.GetProducts(c, db)
	})
	router.PUT("/products/:id", func(c *gin.Context) {
		handlers.UpdateProducts(c, db)
	})
	router.DELETE("/products/:id", func(c *gin.Context) {
		handlers.DeleteProducts(c, db)
	})

	// Run the server
	port := 8080
	log.Printf("Server is running on :%d...\n", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
