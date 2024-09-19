package handler

import (
	"gateway/proto/income"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)
// @Router           /api/income [post]
// @Summary          Create a new income transaction
// @Description      This method creates a new income transaction
// @Tags             INCOME
// @Accept           json
// @Produce          json
// @Param            body    body    income.TransactionInfo    true  "Transaction details"
// @Success          201     {object}  string                  "Transaction created successfully"
// @Failure          400     {object}  error                   "Invalid request body"
// @Failure          500     {object}  error                   "Unable to create income"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) CreateTransaction(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createIncome income.TransactionInfo
	if err := protojson.Unmarshal(bytes, &createIncome); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	trID, err := h.Income.CreateTransaction(c.Request.Context(), &createIncome)
	if err != nil {
		log.Println("Error creating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create income"})
		return
	}

	c.IndentedJSON(201, trID)
}

// @Router           /api/income/{id} [put]
// @Summary          Update an existing transaction by ID
// @Description      This method updates an existing income transaction
// @Tags             INCOME
// @Accept           json
// @Produce          json
// @Param            id      path    string                   true  "Transaction ID"
// @Param            body    body    income.TransactionWithID   true  "Transaction ID and details"
// @Success          200     {object}  income.TransactionInfo    "Transaction updated successfully"
// @Failure          400     {object}  error                     "Invalid request body"
// @Failure          500     {object}  error                     "Unable to update income"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) UpdateTransactionByID(c *gin.Context) {
	id :=c.Param("id")
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var TransactionWithID income.TransactionWithID
	TransactionWithID.Id = id
	if err := protojson.Unmarshal(bytes, &TransactionWithID); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}

	tres, err := h.Income.UpdateTransactionByID(c.Request.Context(), &TransactionWithID)
	if err != nil {
		log.Println("Error updating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update income"})
		return
	}

	c.IndentedJSON(200, tres)
}
// @Router           /api/income/{id} [delete]
// @Summary          Delete a transaction by ID
// @Description      This method deletes an existing income transaction by its ID
// @Tags             INCOME
// @Accept           json
// @Produce          json
// @Param            id      path    string                   true  "Transaction ID"
// @Success          200     {object}  string                  "Transaction deleted successfully"
// @Failure          500     {object}  error                   "Unable to delete income"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	tres, err := h.Income.DeleteTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error deleted income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, tres)
}

// @Router           /api/income/{id} [get]
// @Summary          Get a transaction by ID
// @Description      This method retrieves an income transaction by its ID
// @Tags             INCOME
// @Accept           json
// @Produce          json
// @Param            id      path    string                   true  "Transaction ID"
// @Success          200     {object}  income.TransactionInfo   "Transaction retrieved successfully"
// @Failure          500     {object}  error                   "Unable to retrieve income"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	transactionWithID, err := h.Income.GetTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, transactionWithID)
}

// @Router           /api/income/category/{category} [get]
// @Summary          Get transactions by category
// @Description      This method retrieves income transactions for a specific category
// @Tags             INCOME
// @Security         BearerAuth
// @Produce          json
// @Param            category  path    string                   true  "Transaction category"
// @Success          200       {array}  income.TransactionInfo   "List of transactions for the category"
// @Failure          500       {object}  error                   "Unable to retrieve income transactions"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) GetTransactionsByCategory(c *gin.Context) {
	category := c.Param("categoryy")
	var trCategory income.TransactionCategory
	trCategory.Category = category
	list, err := h.Income.GetTransactionsByCategory(c.Request.Context(), &trCategory)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get income"})
		return
	}

	c.IndentedJSON(200, list)
}
// @Router           /api/income/date [get]
// @Summary          Get transactions by date
// @Description      This method retrieves income transactions for a specific date
// @Tags             INCOME
// @Accept           json
// @Produce          json
// @Param            body    body    income.TransactionDate    true  "Date for transactions"
// @Success          200     {array}  income.TransactionInfo    "List of transactions for the date"
// @Failure          400     {object}  error                   "Invalid request body"
// @Failure          500     {object}  error                   "Unable to retrieve transactions"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) GetTransactionByDate(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var trDate income.TransactionDate
	if err := protojson.Unmarshal(bytes, &trDate); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	lst, err := h.Income.GetTransactionByDate(c.Request.Context(), &trDate)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, lst)
}
