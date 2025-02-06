package handlers

import (
	"finance-operations-service/internal/finance"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, service finance.Service) {
	h := NewHandler(service)

	financeEndpoints := router.Group("/finance")
	{
		financeEndpoints.POST("/replenish", h.Replenish)
		financeEndpoints.POST("/transfer", h.Transfer)

		operationsEndpoints := financeEndpoints.Group("/operations")
		{
			operationsEndpoints.GET("/last", h.LastOperations)
		}
	}
}
