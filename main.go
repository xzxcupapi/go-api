package main

import (
	"database/sql"
	"fmt"
	"log"

	// "net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "Since2024."
	DBName     = "dbenigma"
)

var db *sql.DB

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}

type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Transaction struct {
	ID         int       `json:"id"`
	BillDate   string    `json:"billDate"`
	EntryDate  string    `json:"entryDate"`
	FinishDate string    `json:"finishDate"`
	Employee   Employee  `json:"employee"`
	Customer   Customer  `json:"customer"`
	Details    []Details `json:"billDetails"`
	TotalBill  int       `json:"totalBill"`
}

type Details struct {
	ID           int     `json:"id"`
	BillID       int     `json:"billId"`
	Product      Product `json:"product"`
	ProductPrice int     `json:"productPrice"`
	Qty          int     `json:"qty"`
}

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()

	// Customer API
	router.POST("/customers", createCustomer)
	router.GET("/customers/:id", getCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)

	// Product API
	router.POST("/products", createProduct)
	router.GET("/products", listProducts)
	router.GET("/products/:id", getProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// Transaction API
	router.POST("/transactions", createTransaction)
	router.GET("/transactions/:id", getTransaction)
	router.GET("/transactions", listTransactions)

	log.Fatal(router.Run(":8080"))
}

// Initialize the database connection
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

// CRUD operations for Customer, Product, and Transaction

// Customer API Handlers
func createCustomer(c *gin.Context) {
	// Implement your logic to create a new customer
}

func getCustomer(c *gin.Context) {
	// Implement your logic to get a customer by ID
}

func updateCustomer(c *gin.Context) {
	// Implement your logic to update a customer by ID
}

func deleteCustomer(c *gin.Context) {
	// Implement your logic to delete a customer by ID
}

// Product API Handlers
func createProduct(c *gin.Context) {
	// Implement your logic to create a new product
}

func listProducts(c *gin.Context) {
	// Implement your logic to list all products
}

func getProduct(c *gin.Context) {
	// Implement your logic to get a product by ID
}

func updateProduct(c *gin.Context) {
	// Implement your logic to update a product by ID
}

func deleteProduct(c *gin.Context) {
	// Implement your logic to delete a product by ID
}

// Transaction API Handlers
func createTransaction(c *gin.Context) {
	// Implement your logic to create a new transaction
}

func getTransaction(c *gin.Context) {
	// Implement your logic to get a transaction by ID
}

func listTransactions(c *gin.Context) {
	// Implement your logic to list transactions based on query parameters
}
