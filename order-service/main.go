package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	UserID int         `json:"user_id"`
	Items  []OrderItem `json:"items"`
}

type OrderItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func main() {
	router := buildRouter()
	log.Println("order service started on :8003")
	if err := router.Run(":8003"); err != nil {
		log.Fatal(err)
	}
}

func buildRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.POST("/order", func(c *gin.Context) {
		var req OrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "invalid order payload",
			})
			return
		}

		total := 0.0
		itemCount := 0
		for _, item := range req.Items {
			total += item.Price * float64(item.Count)
			itemCount += item.Count
		}

		c.JSON(http.StatusOK, gin.H{
			"code":          200,
			"message":       "order created",
			"order_id":      fmt.Sprintf("ORD%d", time.Now().UnixMilli()),
			"item_count":    itemCount,
			"total_amount":  total,
			"delivery_fee":  3,
			"estimate_time": "25-35分钟",
			"address":       "北京市朝阳区建国路88号",
		})
	})

	return r
}
