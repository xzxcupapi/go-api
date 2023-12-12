package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Customers struct to represent a customer
type Customers struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"address"`
}

// CreateCustomer handles the creation of a new customer
func CreateCustomer(c *gin.Context, db *sql.DB) {
	var customer Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "INSERT INTO customers (name, phonenumber, address) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Tambahkan log untuk mencetak query SQL yang akan dieksekusi
	fmt.Println("SQL Query:", query)

	// Execute the SQL query
	err = stmt.QueryRow(&customer.Name, &customer.PhoneNumber, &customer.Address).Scan(&customer.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert customer into the database"})
		return
	}

	c.JSON(201, gin.H{"message": "Customer created successfully", "data": customer})
}

// GetCustomer retrieves a customer by ID
func GetCustomer(c *gin.Context, db *sql.DB) {
	customerID := c.Param("id")

	// Prepare the SQL query
	query := "SELECT id, name, phonenumber, address FROM customers WHERE id = $1"
	row := db.QueryRow(query, customerID)

	// Create a customer variable to store the result
	var customer Customers

	// Scan the row data into the customer variable
	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve customer from the database"})
		return
	}

	c.JSON(200, gin.H{"data": customer})
}

// UpdateCustomer updates an existing customer by ID
func UpdateCustomer(c *gin.Context, db *sql.DB) {
	var customer Customers
	customerID := c.Param("id")

	// Bind JSON input to customer struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "UPDATE customers SET name=$1, phonenumber=$2, address=$3 WHERE id=$4"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(customer.Name, customer.PhoneNumber, customer.Address, customerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update customer in the database"})
		return
	}

	c.JSON(200, gin.H{"message": "Customer updated successfully"})
}

// DeleteCustomer deletes a customer by ID
func DeleteCustomer(c *gin.Context, db *sql.DB) {
	customerID := c.Param("id")

	// Prepare the SQL query
	query := "DELETE FROM customers WHERE id=$1"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(customerID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete customer from the database"})
		return
	}

	c.JSON(200, gin.H{"message": "Customer deleted successfully"})
}
