package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Shop struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Rating       float64 `json:"rating"`
	Sales        int     `json:"sales"`
	DeliveryTime string  `json:"deliveryTime"`
	Tag1         string  `json:"tag1"`
	Tag2         string  `json:"tag2"`
	Image        string  `json:"image"`
}

var shops = []Shop{
	{
		ID:           1,
		Name:         "麦当劳",
		Rating:       4.9,
		Sales:        8562,
		DeliveryTime: "20-30",
		Tag1:         "准时必达",
		Tag2:         "品质优选",
		Image:        "https://images.unsplash.com/photo-1565299624946-b28f40a0ae38?w=400&h=300&fit=crop",
	},
	{
		ID:           2,
		Name:         "肯德基",
		Rating:       4.8,
		Sales:        10234,
		DeliveryTime: "25-35",
		Tag1:         "炸鸡专家",
		Tag2:         "新品上架",
		Image:        "https://images.unsplash.com/photo-1626645738196-c2a7c87a8f58?w=400&h=300&fit=crop",
	},
	{
		ID:           3,
		Name:         "沙县小吃",
		Rating:       4.7,
		Sales:        5630,
		DeliveryTime: "15-25",
		Tag1:         "地道美食",
		Tag2:         "实惠超值",
		Image:        "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=400&h=300&fit=crop",
	},
	{
		ID:           4,
		Name:         "兰州拉面",
		Rating:       4.6,
		Sales:        4230,
		DeliveryTime: "20-30",
		Tag1:         "手工拉面",
		Tag2:         "汤鲜肉嫩",
		Image:        "https://images.unsplash.com/photo-1574471440-2b6d80cba3e4?w=400&h=300&fit=crop",
	},
}

func main() {
	router := buildRouter()
	log.Println("shop service started on :8080")
	if err := router.Run(":8080"); err != nil {
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

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, shops)
	})

	return r
}
