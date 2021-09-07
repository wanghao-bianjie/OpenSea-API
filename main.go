package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type",
			"Authorization", "X-Forwarded-For", "X-Real-Ip",
			"X-Appengine-Remote-Addr", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"x-pagination", "Content-Disposition"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//type metaData struct {
	//	image            string
	//	image_data       string
	//	external_url     string
	//	description      string
	//	name             string
	//	attributes       string
	//	background_color string
	//	animation_url    string
	//	youtube_url      string
	//}
	//m := make(map[string]metaData)
	r.GET("/api/:token_id", func(c *gin.Context) {
		tokenId := c.Param("token_id")
		switch tokenId {
		case "1":
			c.JSON(http.StatusOK, gin.H{
				"description":  "Holiday soon...",
				"external_url": "https://www.baidu.com/",
				"image":        "https://img1.baidu.com/it/u=1228507199,1542939359&fm=26&fmt=auto&gp=0.jpg",
				"name":         "Mid-Autumn Festival",
				"attributes": []gin.H{
					{
						"trait_type": "历法",
						"value":      "阴历",
					},
					{
						"trait_type": "月",
						"value":      "八",
					},
					{
						"trait_type": "日",
						"value":      "十五",
					},
					{
						"trait_type": "放假天数",
						"value":      1,
					},
				},
			})
			break
		case "2":
			c.JSON(http.StatusOK, gin.H{
				"description":  "The man is walking",
				"external_url": "https://www.taobao.com/",
				"image":        "https://img2.baidu.com/it/u=363177954,2709021080&fm=26&fmt=auto&gp=0.jpg",
				"name":         "Walking",
			})
			break
		default:
			c.JSON(http.StatusOK, gin.H{
				"description":  "nothing to show ",
				"external_url": "https://www.baidu.com/",
				"image":        "https://img2.baidu.com/it/u=3908224415,1649806984&fm=26&fmt=auto&gp=0.jpg",
				"name":         "nothing",
			})
			break
		}
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
