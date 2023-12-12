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

	c.JSON(201, gin.H{"data": products, "message": "Insert Products Successfully"})
}

// GetProducts retrieves a customer by ID
func GetProducts(c *gin.Context, db *sql.DB) {
	productsID := c.Param("id")

	// Prepare the SQL query
	query := "SELECT id, name, quantity, unit, price FROM products WHERE id = $1"
	row := db.QueryRow(query, productsID)

	// Create a customer variable to store the result
	var products Products

	// Scan the row data into the customer variable
	err := row.Scan(&products.Id, &products.Name, &products.Quantity, &products.Unit, &products.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": "Products not Exist"})
		return
	}

	c.JSON(200, gin.H{"data": products})
}

// UpdateCustomer updates an existing customer by ID
func UpdateProducts(c *gin.Context, db *sql.DB) {
	var products Products
	productsID := c.Param("id")

	// Bind JSON input to customer struct
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the productsID is a valid integer
	if _, err := strconv.Atoi(productsID); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Product ID"})
		return
	}

	// Prepare the SQL query
	query := "UPDATE products SET name=$1, quantity=$2, unit=$3, price=$4 WHERE id=$5"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query not valid"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(products.Name, products.Quantity, products.Unit, products.Price, productsID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to Update Products"})
		return
	}

	c.JSON(200, gin.H{"message": "Product Updated Successfully"})
}

// Function to check if a products with the given ID exists
func ProductsExists(db *sql.DB, productsID string) bool {
	var count int
	query := "SELECT COUNT(*) FROM products WHERE id = $1"
	err := db.QueryRow(query, productsID).Scan(&count)
	return err == nil && count > 0
}

// DeleteCustomer deletes a customer by ID
func DeleteProducts(c *gin.Context, db *sql.DB) {
	productsID := c.Param("id")

	// Check if customer with the given ID exists
	checkQuery := "SELECT COUNT(*) FROM customers WHERE id = $1"
	var count int
	err := db.QueryRow(checkQuery, productsID).Scan(&count)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
		return
	}

	// Prepare the SQL query
	query := "DELETE FROM products WHERE id=$1"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Products Not Exist"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(productsID)
	if err != nil {
		fmt.Println("Error executing delete query:", err)
		c.JSON(500, gin.H{"error": "Failed to Delete Products"})
		return
	}

	c.JSON(200, gin.H{"message": "Products Deleted Successfully"})
}
