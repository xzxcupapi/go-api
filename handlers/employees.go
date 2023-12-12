package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Employees Struct
type Employees struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Address     string `json:"address"`
}

// CreateEmployees handles the creation of a new customer
func CreateEmployees(c *gin.Context, db *sql.DB) {
	var employees Employees
	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL query
	query := "INSERT INTO employees (id, name, phonenumber, address) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
		return
	}
	defer stmt.Close()

	// Log
	fmt.Println("SQL Query:", query)

	// Execute the SQL query
	err = stmt.QueryRow(&employees.Id, &employees.Name, &employees.PhoneNumber, &employees.Address).Scan(&employees.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed Insert Employees"})
		return
	}

	c.JSON(201, gin.H{"data": employees, "message": "Insert Employees Successfully"})
}

// GetEmployees by ID
func GetEmployees(c *gin.Context, db *sql.DB) {
	employeesID := c.Param("id")

	// Prepare the SQL query
	query := "SELECT id, name, phonenumber, address FROM employees WHERE id = $1"
	row := db.QueryRow(query, employeesID)

	var employees Employees

	// Scan the row data
	err := row.Scan(&employees.Id, &employees.Name, &employees.PhoneNumber, &employees.Address)
	if err != nil {
		c.JSON(500, gin.H{"error": "Employees not Exist"})
		return
	}

	c.JSON(200, gin.H{"data": employees})
}

// UpdateEmployees by ID
func UpdateEmployees(c *gin.Context, db *sql.DB) {
	var employees Employees
	employeesID := c.Param("id")

	// Bind JSON input to customer struct
	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate id
	if _, err := strconv.Atoi(employeesID); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Employees ID"})
		return
	}

	// Validate id no exist
	if !customerExists(db, employeesID) {
		c.JSON(404, gin.H{"error": "Employees Not Exist"})
		return
	}

	// Prepare the SQL query
	query := "UPDATE employees SET name=$1, phonenumber=$2, address=$3 WHERE id=$4"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query not valid"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(employees.Name, employees.PhoneNumber, employees.Address, employeesID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed Update Employees"})
		return
	}

	c.JSON(200, gin.H{"message": "Employees updated successfully"})
}

// DeleteEmployees by ID
func DeleteEmployees(c *gin.Context, db *sql.DB) {
	employeesID := c.Param("id")

	// Validate Id
	checkQuery := "SELECT COUNT(*) FROM employees WHERE id = $1"
	var count int
	err := db.QueryRow(checkQuery, employeesID).Scan(&count)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
		return
	}

	if count == 0 {
		c.JSON(404, gin.H{"error": "Employees not Found"})
		return
	}

	// Prepare the SQL query
	query := "DELETE FROM employees WHERE id=$1"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Employees Not Exist"})
		return
	}
	defer stmt.Close()

	// Execute the SQL query
	_, err = stmt.Exec(employeesID)
	if err != nil {
		fmt.Println("Error executing delete query:", err)
		c.JSON(500, gin.H{"error": "Failed to Delete Employees"})
		return
	}

	c.JSON(200, gin.H{"message": "Employees Deleted Successfully"})
}
