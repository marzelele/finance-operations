package handlers

import (
	"finance-operations-service/internal/finance/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ErrResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func (h *Handler) Replenish(c *gin.Context) {
	var req models.Funds

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrResponse(err))
		return
	}

	operationID, err := h.financeService.Replenish(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"operation_id": operationID})
}

func (h *Handler) LastOperations(c *gin.Context) {
	var (
		id  uuid.UUID
		err error
	)

	if id, err = uuid.Parse(c.Query("user_id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid user id: %s", id)})
		return
	}

	ops, err := h.financeService.LastOperations(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"operations": ops})
}

func (h *Handler) Transfer(c *gin.Context) {
	var req models.DetailsOperation

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrResponse(err))
		return
	}

	operationID, err := h.financeService.Transfer(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"operation_id": operationID})
}
