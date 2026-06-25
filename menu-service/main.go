package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Food struct {
	ID       int     `json:"id"`
	ShopName string  `json:"shop_name"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Sales    int     `json:"sales"`
	Rating   float64 `json:"rating"`
	Desc     string  `json:"desc"`
	Badge    string  `json:"badge,omitempty"`
	Image    string  `json:"image"`
}

var menus = map[int][]Food{
	1: {
		{ID: 1, ShopName: "麦当劳", Name: "巨无霸", Price: 32, Sales: 2600, Rating: 4.9, Desc: "经典牛肉汉堡，层次丰富", Badge: "招牌", Image: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=300&h=200&fit=crop"},
		{ID: 2, ShopName: "麦当劳", Name: "薯条", Price: 12, Sales: 3200, Rating: 4.8, Desc: "外脆里软，现炸更香", Image: "https://images.unsplash.com/photo-1573080496219-bb080dd4f877?w=300&h=200&fit=crop"},
		{ID: 3, ShopName: "麦当劳", Name: "麦辣鸡腿堡", Price: 29, Sales: 2100, Rating: 4.7, Desc: "香辣过瘾，口感十足", Badge: "热销", Image: "https://images.unsplash.com/photo-1550317138-10000687a72b?w=300&h=200&fit=crop"},
		{ID: 4, ShopName: "麦当劳", Name: "可乐", Price: 8, Sales: 2800, Rating: 4.8, Desc: "冰爽畅快，搭配首选", Image: "https://images.unsplash.com/photo-1622483767028-3f66f32aef97?w=300&h=200&fit=crop"},
	},
	2: {
		{ID: 5, ShopName: "肯德基", Name: "香辣鸡腿堡", Price: 28, Sales: 2400, Rating: 4.8, Desc: "香辣多汁，经典人气款", Badge: "招牌", Image: "https://images.unsplash.com/photo-1606755962773-d324e0a13086?w=300&h=200&fit=crop"},
		{ID: 6, ShopName: "肯德基", Name: "新奥尔良烤翅", Price: 22, Sales: 1900, Rating: 4.7, Desc: "鲜嫩入味，甜辣适中", Image: "https://images.unsplash.com/photo-1562967914-608f82629710?w=300&h=200&fit=crop"},
		{ID: 7, ShopName: "肯德基", Name: "吮指原味鸡", Price: 19, Sales: 3100, Rating: 4.9, Desc: "外酥里嫩，经典口味", Badge: "必点", Image: "https://images.unsplash.com/photo-1513639776629-7b61b0ac49cb?w=300&h=200&fit=crop"},
		{ID: 8, ShopName: "肯德基", Name: "土豆泥", Price: 10, Sales: 1500, Rating: 4.6, Desc: "细腻绵软，口感顺滑", Image: "https://images.unsplash.com/photo-1608500218808-3442ec7a520f?w=300&h=200&fit=crop"},
	},
	3: {
		{ID: 9, ShopName: "沙县小吃", Name: "拌面", Price: 14, Sales: 2800, Rating: 4.7, Desc: "酱香浓郁，简单耐吃", Badge: "招牌", Image: "https://images.unsplash.com/photo-1612929633738-8fe44f7ec841?w=300&h=200&fit=crop"},
		{ID: 10, ShopName: "沙县小吃", Name: "蒸饺", Price: 16, Sales: 2300, Rating: 4.8, Desc: "皮薄馅足，现包现蒸", Image: "https://images.unsplash.com/photo-1544025162-d76694265947?w=300&h=200&fit=crop"},
		{ID: 11, ShopName: "沙县小吃", Name: "扁肉汤", Price: 12, Sales: 1700, Rating: 4.6, Desc: "汤鲜肉嫩，暖胃舒心", Image: "https://images.unsplash.com/photo-1547592180-85f173990554?w=300&h=200&fit=crop"},
		{ID: 12, ShopName: "沙县小吃", Name: "鸡腿饭", Price: 22, Sales: 2100, Rating: 4.7, Desc: "分量扎实，性价比高", Badge: "热销", Image: "https://images.unsplash.com/photo-1512058564366-18510be2db19?w=300&h=200&fit=crop"},
	},
	4: {
		{ID: 13, ShopName: "兰州拉面", Name: "牛肉拉面", Price: 18, Sales: 3400, Rating: 4.8, Desc: "汤清面劲道，牛肉鲜香", Badge: "招牌", Image: "https://images.unsplash.com/photo-1617093727343-374698b1b08d?w=300&h=200&fit=crop"},
		{ID: 14, ShopName: "兰州拉面", Name: "加肉拉面", Price: 24, Sales: 2200, Rating: 4.9, Desc: "分量更足，肉香浓郁", Image: "https://images.unsplash.com/photo-1569718212165-3a8278d5f624?w=300&h=200&fit=crop"},
		{ID: 15, ShopName: "兰州拉面", Name: "凉菜拼盘", Price: 13, Sales: 1600, Rating: 4.6, Desc: "清爽开胃，搭配首选", Image: "https://images.unsplash.com/photo-1544025162-d76694265947?w=300&h=200&fit=crop"},
		{ID: 16, ShopName: "兰州拉面", Name: "红烧牛肉面", Price: 22, Sales: 2500, Rating: 4.7, Desc: "浓郁汤底，香气十足", Badge: "必点", Image: "https://images.unsplash.com/photo-1574471440-2b6d80cba3e4?w=300&h=200&fit=crop"},
	},
}

func main() {
	router := buildRouter()
	log.Println("menu service started on :8002")
	if err := router.Run(":8002"); err != nil {
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

	r.GET("/menu/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid shop id"})
			return
		}

		menu, ok := menus[id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"message": "shop menu not found"})
			return
		}

		c.JSON(http.StatusOK, menu)
	})

	return r
}
