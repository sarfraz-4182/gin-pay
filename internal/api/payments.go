package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarfraz-4182/gin-pay/internal/db"
	"github.com/sarfraz-4182/gin-pay/internal/models"
)

type CreatePaymentReq struct {
	Amount     int64  `json:"amount" binding:"required,gt=0"`
	Currency   string `json:"currency" binding:"required"`
	CustomerID string `json:"customer_id" binding:"required"`
}

func RegisterPaymentRoutes(rg *gin.RouterGroup) {
	p := rg.Group("/payments")
	p.POST("/", CreatePayment)
	p.GET("/:id", GetPayment)
	p.GET("/", ListPayments)
}

func CreatePayment(c *gin.Context) {
	var body CreatePaymentReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Payment{
		Amount:     body.Amount,
		Currency:   body.Currency,
		CustomerID: body.CustomerID,
		Status:     "pending",
	}

	if err := db.DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment"})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func GetPayment(c *gin.Context) {
	id := c.Param("id")
	var p models.Payment
	if err := db.DB.First(&p, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func ListPayments(c *gin.Context) {
	var payments []models.Payment
	db.DB.Order("id desc").Limit(100).Find(&payments)
	c.JSON(http.StatusOK, payments)
}
