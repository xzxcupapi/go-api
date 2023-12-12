package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Transaction Struct
type Transaction struct {
	Id         int       `json:"id"`
	BillDate   time.Time `json:"billdate"`
	EntryDate  time.Time `json:"entrydate"`
	FinishDate time.Time `json:"finishdate"`
	EmployeeId int       `json:"employeeid"`
	CustomerId int       `json:"customerid"`
}

// CreateTransaction
func CreateTransaction(c *gin.Context, db *sql.DB) {
	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Format dates
	formattedBillDate := transaction.BillDate.Format("2006-01-02")
	formattedEntryDate := transaction.EntryDate.Format("2006-01-02")
	formattedFinishDate := transaction.FinishDate.Format("2006-01-02")

	query := "INSERT INTO transaction (id, billdate, entrydate, finishdate, emploeeid, customerid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Query Not Valid"})
	}
	defer stmt.Close()

	//Log
	fmt.Println("SQL Query", query)

	//Execute SQL Query
	err = stmt.QueryRow(&transaction.Id, formattedBillDate, formattedEntryDate, formattedFinishDate, &transaction.EmployeeId, &transaction.CustomerId).Scan(&transaction.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed Insert Transaction"})
		return
	}

	c.JSON(201, gin.H{"data": transaction, "message": "Insert Products Successfully"})
}
