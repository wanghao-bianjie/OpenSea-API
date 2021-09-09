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

	r.GET("/api/contract", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":                    "collection-1155",
			"description":             "这里是 ERC-1155 的集合的描述。。。",
			"image":                   "https://img0.baidu.com/it/u=2555070557,2516571774&fm=253&fmt=auto&app=120&f=PNG?w=503&h=485",
			"external_link":           "https://eips.ethereum.org/EIPS/eip-1155",
			"seller_fee_basis_points": 111,                                          // # Indicates a 1.11% seller fee.(x/10000)
			"fee_recipient":           "0x5fbca2ed2040924f65133A25194eeEBd45864f24", // # Where seller fees will be paid to.
		})
	})

	r.GET("/api/:token_id", func(c *gin.Context) {
		tokenId := c.Param("token_id")
		switch tokenId {
		case "1":
			c.JSON(http.StatusOK, gin.H{
				"name":          "Mid-Autumn Festival",
				"description":   "Holiday soon...",
				"external_url":  "https://www.baidu.com/",
				"image":         "https://img1.baidu.com/it/u=1228507199,1542939359&fm=26&fmt=auto&gp=0.jpg",
				"animation_url": "http://vjs.zencdn.net/v/oceans.mp4",
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
				"name":             "Walking",
				"description":      "The man is walking",
				"external_url":     "https://www.taobao.com/",
				"image":            "https://img2.baidu.com/it/u=363177954,2709021080&fm=26&fmt=auto&gp=0.jpg",
				"background_color": "FFFFFF",
				"animation_url":    "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
				"youtube_url":      "https://www.youtube.com/watch?v=UCxA8sDltMY",
				"attributes": []gin.H{
					{
						"trait_type": "天气",
						"value":      "晴",
					},
					{
						"trait_type": "颜色",
						"value":      "白色",
					},
					{
						"trait_type": "季节",
						"value":      "秋天",
					},
					{
						"trait_type": "等级",
						"value":      1,
						"max_value":  30,
					},
					{
						"trait_type": "攻击力",
						"value":      100.5,
					},
					{
						"display_type": "number",
						"trait_type":   "排名",
						"value":        3,
					},
					{
						"display_type": "boost_percentage",
						"trait_type":   "占比",
						"value":        77.7,
					},
					{
						"display_type": "boost_number",
						"trait_type":   "速度 km/h",
						"value":        60,
					},
					{
						"display_type": "date",
						"trait_type":   "生日",
						"value":        time.Now().Unix(),
					},
					{
						"value": "没有设置trait_type",
					},
				},
			})
			break
		case "3":
			c.JSON(http.StatusOK, gin.H{
				"name":             "孙悟空",
				"description":      "龙珠孙悟空",
				"external_url":     "https://www.taobao.com/",
				"image":            "https://img.cppng.com/download/2020-06/20414-7-goku-transparent-picture.png",
				"background_color": "FFFFFF",
				//"animation_url":    "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
				"youtube_url": "https://www.youtube.com/watch?v=UCxA8sDltMY",
				"attributes": []gin.H{
					{
						"trait_type": "国家",
						"value":      "日本",
					},
					{
						"trait_type": "动画",
						"value":      "龙珠",
					},
					{
						"trait_type": "性别",
						"value":      "男",
					},
					{
						"trait_type": "战斗力",
						"value":      9999,
						"max_value":  10000,
					},
					{
						"trait_type": "护甲",
						"value":      66,
					},
					{
						"display_type": "number",
						"trait_type":   "排名",
						"value":        1,
						"max_value":    100,
					},
					{
						"display_type": "boost_percentage",
						"trait_type":   "受欢迎占比",
						"value":        88.8,
					},
					{
						"display_type": "boost_number",
						"trait_type":   "移动速度 km/h",
						"value":        120,
					},
					{
						"display_type": "date",
						"trait_type":   "生日",
						"value":        time.Now().Unix(),
					},
					{
						"value": 777,
					},
				},
			})
			break
		case "4":
			c.JSON(http.StatusOK, gin.H{
				"name":             "布欧",
				"description":      "龙珠布欧",
				"external_url":     "https://www.taobao.com/",
				"image":            "https://img2.baidu.com/it/u=2960282707,3861898616&fm=26&fmt=auto&gp=0.jpg",
				"background_color": "FFFFFF",
				//"animation_url":    "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
				//"youtube_url": "https://www.youtube.com/watch?v=UCxA8sDltMY",
				"attributes": []gin.H{
					{
						"trait_type": "国家",
						"value":      "日本",
					},
					{
						"trait_type": "动画",
						"value":      "龙珠",
					},
					{
						"trait_type": "性别",
						"value":      "男",
					},
					{
						"trait_type": "战斗力",
						"value":      8888,
						"max_value":  10000,
					},
					{
						"trait_type": "护甲",
						"value":      77,
					},
					{
						"display_type": "number",
						"trait_type":   "排名",
						"value":        2,
						"max_value":    100,
					},
					{
						"display_type": "boost_percentage",
						"trait_type":   "受欢迎占比",
						"value":        66.6,
					},
					{
						"display_type": "boost_number",
						"trait_type":   "移动速度 km/h",
						"value":        115,
					},
					{
						"display_type": "date",
						"trait_type":   "生日",
						"value":        time.Now().Unix(),
					},
				},
			})
			break
		case "5":
			c.JSON(http.StatusOK, gin.H{
				"name":             "人造人18号",
				"description":      "龙珠人造人18号",
				"external_url":     "https://www.taobao.com/",
				"image":            "https://img2.baidu.com/it/u=4063123354,1403985731&fm=26&fmt=auto&gp=0.jpg",
				"background_color": "FFFFFF",
				//"animation_url":    "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
				//"youtube_url": "https://www.youtube.com/watch?v=UCxA8sDltMY",
				"attributes": []gin.H{
					{
						"trait_type": "国家",
						"value":      "日本",
					},
					{
						"trait_type": "动画",
						"value":      "龙珠",
					},
					{
						"trait_type": "性别",
						"value":      "女",
					},
					{
						"trait_type": "战斗力",
						"value":      5678,
						"max_value":  10000,
					},
					{
						"trait_type": "护甲",
						"value":      20,
					},
					{
						"display_type": "number",
						"trait_type":   "排名",
						"value":        4,
						"max_value":    100,
					},
					{
						"display_type": "boost_percentage",
						"trait_type":   "受欢迎占比",
						"value":        98.7,
					},
					{
						"display_type": "boost_number",
						"trait_type":   "移动速度 km/h",
						"value":        99,
					},
					{
						"display_type": "date",
						"trait_type":   "生日",
						"value":        time.Now().Unix(),
					},
				},
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
