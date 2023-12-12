package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Customers struct to represent a customer
type Products struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
	Price    int    `json:"price"`
}

// CreateCustomer handles the creation of a new customer
func CreateProducts(c *gin.Context, db *sql.DB) {
	var products Products
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "INSERT INTO products (name, quantity, unit, price) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
		return
	}
	defer stmt.Close()

	// Tambahkan log untuk mencetak query SQL yang akan dieksekusi
	fmt.Println("SQL Query:", query)

	// Execute the SQL query
	err = stmt.QueryRow(&products.Name, &products.Quantity, &products.Unit, &products.Price).Scan(&products.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to Insert Customer"})
		return
	}

	c.JSON(201, gin.H{"data": products, "message": "Insert Customer Successfully"})
}

// GetCustomer retrieves a customer by ID
func GetProducts(c *gin.Context, db *sql.DB) {
	customerID := c.Param("id")

	// Prepare the SQL query
	query := "SELECT id, name, phonenumber, address FROM customers WHERE id = $1"
	row := db.QueryRow(query, customerID)

	// Create a customer variable to store the result
	var customer Customers

	// Scan the row data into the customer variable
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		c.JSON(500, gin.H{"error": "Customer not Exist"})
		return
	}

	c.JSON(200, gin.H{"data": customer})
}

// UpdateCustomer updates an existing customer by ID
func UpdateProducts(c *gin.Context, db *sql.DB) {
	var customer Customers
	customerID := c.Param("id")

	// Bind JSON input to customer struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the customerID is a valid integer
	if _, err := strconv.Atoi(customerID); err != nil {
		c.JSON(400, gin.H{"error": "Invalid customer ID"})
		return
	}

	// Check if the customer with the given ID exists
	if !customerExists(db, customerID) {
		c.JSON(404, gin.H{"error": "Customer Not Exist"})
		return
	}

	// Prepare the SQL query
	query := "UPDATE customers SET name=$1, phonenumber=$2, address=$3 WHERE id=$4"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query not valid"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(customer.Name, customer.PhoneNumber, customer.Address, customerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(200, gin.H{"message": "Customer updated successfully"})
}

// Function to check if a customer with the given ID exists
func ProductsExists(db *sql.DB, customerID string) bool {
	var count int
	query := "SELECT COUNT(*) FROM customers WHERE id = $1"
	err := db.QueryRow(query, customerID).Scan(&count)
	return err == nil && count > 0
}

// DeleteCustomer deletes a customer by ID
func DeleteProducts(c *gin.Context, db *sql.DB) {
	customerID := c.Param("id")

	// Check if customer with the given ID exists
	checkQuery := "SELECT COUNT(*) FROM customers WHERE id = $1"
	var count int
	err := db.QueryRow(checkQuery, customerID).Scan(&count)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
		return
	}

	if count == 0 {
		c.JSON(404, gin.H{"error": "Customer not Found"})
		return
	}

	// Prepare the SQL query
	query := "DELETE FROM customers WHERE id=$1"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Customer Not Exist"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(customerID)
	if err != nil {
		fmt.Println("Error executing delete query:", err)
		c.JSON(500, gin.H{"error": "Failed to Delete Customer"})
		return
	}

	c.JSON(200, gin.H{"message": "Customer Deleted Successfully"})
}
