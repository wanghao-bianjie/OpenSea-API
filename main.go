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

	r.GET("/api/contract", func(c *gin.Context) {	//test1 test3
		c.JSON(http.StatusOK, gin.H{
			"name":                    "collection-1155",
			"description":             "这里是 ERC-1155 的集合的描述。。。",
			"image":                   "https://img0.baidu.com/it/u=2555070557,2516571774&fm=253&fmt=auto&app=120&f=PNG?w=503&h=485",
			"external_link":           "https://eips.ethereum.org/EIPS/eip-1155",
			"seller_fee_basis_points": 111,                                          // # Indicates a 1.11% seller fee.(x/10000)
			"fee_recipient":           "0x5fbca2ed2040924f65133A25194eeEBd45864f24", // # Where seller fees will be paid to.
		})
	})
	r.GET("/api/contract-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":                    "collection-1155-0909",
			"description":             "这里是 ERC-1155 的集合的描述，我是合约的 contractURI 中配置的接口返回的数据。。。",
			"image":                   "https://thehustle.co/wp-content/uploads/2021/07/News-Brief_2021-07-23T014014.791Z.jpg",
			"external_link":           "https://eips.ethereum.org/EIPS/eip-1155",
			"seller_fee_basis_points": 750,                                          // # Indicates a 1.11% seller fee.(x/10000)
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
		case "6":
			c.JSON(http.StatusOK, gin.H{
				"name": "《清漓渔歌》荣宝斋1:1限量版精品复制",
				"description": `# 传世精彩&bull;忠于原作 #
&emsp;&emsp;《清漓渔歌》由350年历史的“中华老字号”荣宝斋出品、制作。


# 产品信息 #
&emsp;&emsp;作品名称：《清漓渔歌》荣宝斋1:1限量版精品复制

&emsp;&emsp;作    家：李可染

&emsp;&emsp;出品单位：荣宝斋

&emsp;&emsp;作品规格：68.7x45cm (因手工测量方式不同，存在1~2cm的误差值，属于合理范围。)

&emsp;&emsp;产品大小：约6.5平尺

&emsp;&emsp;产品材质：设色纸本

&emsp;&emsp;装裱形式：镜心

&emsp;&emsp;发行总数：限量10000幅


# 作者介绍 #
&emsp;&emsp;李可染(1907-1989)，江苏徐州人。中国近代杰出的画家、诗人，画家齐白石的弟子。李可染自幼即喜绘画，13岁时学画山水。43岁任中央美术学院教授，49岁为变革山水画，行程数万里旅行写生。72岁任中国美术家协会幅主席、中国画研究院院长。晚年用笔趋于老辣。擅长画山水、人物，尤其擅长画牛。代表画作有《漓江胜境图》、《万山红遍》、《井冈山》等。
代表画集有《李可染水墨写生画集》、《李可染中国画集》、《李可染画牛》等。

# 产品信息 #
&emsp;&emsp;《清漓渔歌图》是李可染先生存世不多的作品之一，此幅作品以淡为主，把南方的水乡景色用俯视的手法，以朦胧、清秀手法表现出来，央视《鉴宝》鉴定专家-王洪斌先生曾讲此幅作品的极具收藏价值，未来的升值潜力不可估量。


# 产品构成 #
&emsp;&emsp;&bull;限量复制画一副

&emsp;&emsp;&bull;防伪卡袋一个

&emsp;&emsp;&bull;防伪卡一张

&emsp;&emsp;&bull;外包装盒一个


# 提货凭证 #
&emsp;&emsp;一体化数字艺术品（IDA）电子提货凭证，登录[www.ip.pub](http://www.ip.pub "www.ip.pub")在个人中心查看凭证。


# 作品工艺 #
&emsp;&emsp;作为荣宝斋1:1限量版精品复制作品，由荣宝斋独家工艺制作，荣宝斋的馆藏资料室采用当代最先进的设备对原作进行扫描、输入，确保将原作信息准确、全面的转换成数字信息。遵循原作的材料，选择高档的宣纸、绢布，引入水性颜料复制，增加渗透性及国画感。并经过传统工艺及专家的润色、增色处理，使其克服了一般高仿作品墨纸分离、暗调不够暗、层次不分明的问题，完全可以保证原作的真实性。


# 防伪溯源 #
&emsp;&emsp;每件作品附有唯一编号和标识码的荣宝斋证书，用户可通过下载“荣宝斋”手机APP（目前仅支持苹果手机下载），扫码防伪证书，通过扫描结果了解是否为荣宝斋官方出品。

&bull;APP扫码验证

&emsp;&emsp;1.应用商店下载“荣宝斋”APP

&emsp;&emsp;2.在APP内打开[防伪认证]页面

&emsp;&emsp;3.扫描防伪证书上的二维码

&emsp;&emsp;4.查看验证结果


&bull;字画背面编号

&emsp;&emsp;字画背面印有“荣宝制作”红章及编号。

&bull;识别伪造仿品

&emsp;&emsp;1. 标注非1:1限量复制
	
&emsp;&emsp;2. 无防伪证书
	
&emsp;&emsp;3. 非荣宝斋制作

	
	
# 购买须知 #
&bull;制作时间

&emsp;&emsp;您选择购买的荣宝斋1:1限量版精品复制作品，属于高端定制产品，由荣宝斋独家工艺制作。自您付款之日起，预计制作时间最长为20个工作日。

&bull;专业托管

&emsp;&emsp;您购买后如果选择托管，我们将保管在恒温恒湿的专业仓库，您在提货时若发生任何因托管产生的质量问题，我们将负责免费重新制作。
在托管过程中，您可以选择通过摄像头随时查看所托管的作品。

&bull;自行提货

&emsp;&emsp;您选择提货后，将在3个工作日内进行多重质检，确保无任何质量问题后发出。
提货后建议您尽快选择专业的书画装裱机构，对书画进行装裱，以进行更好的保护。并保持书画存放环境干燥、卫生，避免潮湿、紫外线直射。勿用尖锐物品接触画，平时可用软毛刷轻轻掸扫书画表面。

&bull;关于快递

&emsp;&emsp;合作快递为顺丰，每幅作品物流我们将购买物流保险，若因物流发生损坏，将由物流公司进行赔偿。

&bull;关于签收

&emsp;&emsp;产品均经过严格检验后发货，收到货后请立即拍摄开盒视频，若出现数量不对、破损等物流过程问题，请拒绝签收，并通过物流索赔。`,
				"image":            "http://saastestimages.digit.vc/testsaas/4/775145175260/20210617155030941.jpg",
				"external_url":     "https://www.ip.pub",
				"background_color": "FFFFFF",
				"animation_url":    "",
				"youtube_url":      "",
				"attributes": []gin.H{
					{
						"trait_type": "作家",
						"value":      "李可染",
					},
					{
						"trait_type": "实物规格",
						"value":      "68.7x45cm",
					},
					{
						"trait_type": "实物材质",
						"value":      "设色纸本",
					},
					{
						"trait_type": "装裱",
						"value":      "镜心",
					},
					{
						"display_type": "number",
						"trait_type":   "排名",
						"value":        9999,
						"max_value":    10000,
					},
				},
			})
			break
		default:
			c.JSON(http.StatusOK, gin.H{
				//"description":  "nothing to show ",
				"description": `# 一级标题测试
#### 四级标题测试

*斜体测试*

|  表头   | 表头  |
|  ----  | ----  |
| 单元格  | 单元格 |
| 单元格  | 单元格 |

> 以下是图片测试

![alt text](https://www.markdownguide.org/assets/images/tux.png)`,
				"external_url": "https://www.baidu.com/",
				"image":        "https://img2.baidu.com/it/u=3908224415,1649806984&fm=26&fmt=auto&gp=0.jpg",
				"name":         "nothing",
			})
			break
		}
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
