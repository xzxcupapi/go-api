package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"net/http"

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
	// router.DELETE("/customers/:id", deleteCustomer)

	// Product API
	// router.POST("/products", createProduct)
	// router.GET("/products", listProducts)
	// router.GET("/products/:id", getProduct)
	// router.PUT("/products/:id", updateProduct)
	// router.DELETE("/products/:id", deleteProduct)

	// Transaction API
	// router.POST("/transactions", createTransaction)
	// router.GET("/transactions/:id", getTransaction)
	// router.GET("/transactions", listTransactions)

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

// createCustomer creates a new customer.
func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "INSERT INTO customers (name, phone_number, address) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	err = stmt.QueryRow(customer.Name, customer.PhoneNumber, customer.Address).Scan(&customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert customer into the database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully", "data": customer})
}

// getCustomer retrieves a customer by ID.
func getCustomer(c *gin.Context) {
	id := c.Param("id")

	// Convert the ID to an integer
	customerID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Query to fetch a customer by ID
	row := db.QueryRow("SELECT id, name, phone_number, address FROM customers WHERE id = $1", customerID)

	// Create a Customer object to hold the result
	var customer Customer

	// Scan the row data into the customer object
	err = row.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer retrieved successfully", "data": customer})
}

// updateCustomer updates a customer by ID.
func updateCustomer(c *gin.Context) {
	id := c.Param("id")

	// Convert the ID to an integer
	customerID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Create a Customer object to hold the request payload
	var updatedCustomer Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the request payload
	// Add your custom validation logic here

	// Update query
	result, err := db.Exec("UPDATE customers SET name=$1, phone_number=$2, address=$3 WHERE id=$4",
		updatedCustomer.Name, updatedCustomer.PhoneNumber, updatedCustomer.Address, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	// Check the number of rows affected by the update
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

// // deleteCustomer deletes a customer by ID.
// func deleteCustomer(c *gin.Context) {
// 	id := c.Param("id")

// 	// Implement your logic to delete the customer from the database by ID
// 	// ...

// 	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully", "data": "OK"})
// }

// // Product API Handlers

// // createProduct creates a new product.
// func createProduct(c *gin.Context) {
// 	var product Product
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Implement your logic to save the product to the database
// 	// ...

// 	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "data": product})
// }

// // listProducts lists all products.
// func listProducts(c *gin.Context) {
// 	// Implement your logic to fetch all products from the database
// 	// ...

// 	// Example response (replace with actual fetched data)
// 	products := []Product{
// 		{ID: 1, Name: "Product A", Price: 50, Unit: "Unit"},
// 		{ID: 2, Name: "Product B", Price: 30, Unit: "Kg"},
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Products retrieved successfully", "data": products})
// }

// // getProduct retrieves a product by ID.
// func getProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	// Implement your logic to fetch the product from the database by ID
// 	// ...

// 	// Example response (replace with actual fetched data)
// 	product := Product{ID: 1, Name: "Product A", Price: 50, Unit: "Unit"}
// 	c.JSON(http.StatusOK, gin.H{"message": "Product retrieved successfully", "data": product})
// }

// // updateProduct updates a product by ID.
// func updateProduct(c *gin.Context) {
// 	id := c.Param("id")
// 	var updatedProduct Product
// 	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Implement your logic to update the product in the database by ID
// 	// ...

// 	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "data": updatedProduct})
// }

// // deleteProduct deletes a product by ID.
// func deleteProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	// Implement your logic to delete the product from the database by ID
// 	// ...

// 	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "data": "OK"})
// }

// // Transaction API Handlers

// // createTransaction creates a new transaction.
// func createTransaction(c *gin.Context) {
// 	var transaction Transaction
// 	if err := c.ShouldBindJSON(&transaction); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Implement your logic to save the transaction to the database
// 	// ...

// 	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "data": transaction})
// }

// // getTransaction retrieves a transaction by ID.
// func getTransaction(c *gin.Context) {
// 	id := c.Param("id")

// 	// Implement your logic to fetch the transaction from the database by ID
// 	// ...

// 	// Example response (replace with actual fetched data)
// 	transaction := Transaction{ID: 1, BillDate: "2023-01-01", EntryDate: "2023-01-01", FinishDate: "2023-01-02", EmployeeID: "emp123", CustomerID: "cust456"}
// 	c.JSON(http.StatusOK, gin.H{"message": "Transaction retrieved successfully", "data": transaction})
// }

// // listTransactions lists transactions based on query parameters.
// func listTransactions(c *gin.Context) {
// 	// Implement your logic to fetch transactions from the database based on query parameters
// 	// ...

// 	// Example response (replace with actual fetched data)
// 	transactions := []Transaction{
// 		{ID: 1, BillDate: "2023-01-01", EntryDate: "2023-01-01", FinishDate: "2023-01-02", EmployeeID: "emp123", CustomerID: "cust456"},
// 		{ID: 2, BillDate: "2023-01-02", EntryDate: "2023-01-02", FinishDate: "2023-01-03", EmployeeID: "emp456", CustomerID: "cust789"},
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Transactions retrieved successfully", "data": transactions})
// }
