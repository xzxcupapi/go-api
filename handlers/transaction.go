package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Transaction Struct
type Transaction struct {
	Id         int       `json:"id"`
	BillDate   time.Time `json:"billdate" format:"2006-01-02"`
	EntryDate  time.Time `json:"entrydate" format:"2006-01-02"`
	FinishDate time.Time `json:"finishdate" format:"2006-01-02"`
	EmployeeId int       `json:"employeeid"`
	CustomerId int       `json:"customerid"`
}

// CreateTransaction
// CreateTransaction
func CreateTransaction(c *gin.Context, db *sql.DB) {
	var transaction Transaction
	if err := c.ShouldBindWith(&transaction, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO transactions (id, billdate, entrydate, finishdate, employeeid, customerid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to prepare the database query"})
		return
	}
	defer stmt.Close()

	// Log
	fmt.Println("SQL Query", query)

	// Execute SQL Query
	err = stmt.QueryRow(&transaction.Id, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.EmployeeId, &transaction.CustomerId).Scan(&transaction.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert transaction into the database"})
		return
	}

	c.JSON(201, gin.H{"data": transaction, "message": "Insert Transaction Successfully"})
}

// GetTransaction by ID
func GetTransaction(c *gin.Context, db *sql.DB) {
	transactionID := c.Param("id")

	//Prepare Query SQL
	query := "SELECT id, billdate, entrydate, finishdate, employeeid, customerid FROM transactions WHERE id = $1"
	row := db.QueryRow(query, transactionID)

	var transaction Transaction

	//Scan the row data
	err := row.Scan(&transaction.Id, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.EmployeeId, &transaction.CustomerId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Transaction not Exist"})
		return
	}

	c.JSON(200, gin.H{"data": transaction})
}
