package handler

import (
	"gateway/proto/budget"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)
// @Router           /api/budgets/create [post]
// @Summary          Create a new budget
// @Description      This method creates a new budget
// @Tags             BUDGETS
// @Accept           json
// @Produce          json
// @Param            body    body    budget.BudgetInfo    true  "Budget details"
// @Success          201     {object}  string              "Budget created successfully"
// @Failure          400     {object}  error               "Invalid request body"
// @Failure          500     {object}  error               "Unable to create budget"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) CreateBudget(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createBudget budget.BudgetInfo
	if err := protojson.Unmarshal(bytes, &createBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetID, err := h.Budget.CreateBudget(c.Request.Context(), &createBudget)
	if err != nil {
		log.Println("Error creating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create budget"})
		return
	}

	c.IndentedJSON(201, budgetID)
}

// @Router           /api/budgets [put]
// @Summary          Update budget amount
// @Description      This method updates the amount of an existing budget
// @Tags             BUDGETS
// @Accept           json
// @Produce          json
// @Param            body    body    budget.BudgetUpdate   true  "Updated budget details"
// @Success          200     {object}  budget.BudgetInfo    "Budget updated successfully"
// @Failure          400     {object}  error                "Invalid request body"
// @Failure   500     {object}  error                "Unable to update budget"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) UpdateBudgetAmount(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}	

	var updateBudget budget.BudgetUpdate
	if err := protojson.Unmarshal(bytes, &updateBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.UpdateBudgetAmount(c.Request.Context(), &updateBudget)
	if err != nil {
		log.Println("Error updating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}

// @Router           /api/budgets [get]
// @Summary          Get all budgets
// @Description      This method retrieves all budgets
// @Tags             BUDGETS
// @Security         BearerAuth
// @Produce          json
// @Success          200     {array}  budget.BudgetInfo    "List of budgets"
// @Failure       500     {object}  error                "Unable to retrieve budgets"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) GetBudgets(c *gin.Context) {
	var budgetEmpty budget.Empty
	budgetWithID, err := h.Budget.GetBudgets(c.Request.Context(), &budgetEmpty)
	if err != nil {
		log.Println("Error getting budgets:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to retrieve budgets"})
		return
	}

	c.IndentedJSON(200, budgetWithID)
}

// @Router           /api/budgets/{id} [delete]
// @Summary          Delete a budget by ID
// @Description      This method deletes an existing budget by its ID
// @Tags             BUDGETS
// @Accept           json
// @Produce          json
// @Param            id      path    string               true  "Budget ID"
// @Success          200     {object}  string              "Budget deleted successfully"
// @Failure          400     {object}  error               "Invalid request body"
// @Failure          500     {object}  error               "Unable to delete budget"
// @Failure          403     {object}  error             "Permission Denied"

func (h *Handler) DeleteBudgetByID(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var budgetID budget.BudgetID
	if err := protojson.Unmarshal(bytes, &budgetID); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.DeleteBudgetByID(c.Request.Context(), &budgetID)
	if err != nil {
		log.Println("Error deleting budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}
