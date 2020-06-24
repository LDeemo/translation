package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"translation_web/entry"
	"translation_web/util"
)

var appInfo entry.AppConfig
var transParams entry.TranslationParams
var res entry.Result

func init() {
	appInfo.AppKey = "27b2f18edab025ae"
	appInfo.AppPwd = "uu2VMZ7FfJnnjSmCrZlqiFmsmFitfRyN"
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.POST("/query", query)
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"可以了",
		})
	})
	r.Run(":9090")
}

func query(c *gin.Context) {
	transParams.AppConfig = appInfo
	err := c.ShouldBind(&transParams)
	if err != nil {
		log.Fatalf("banding params failed,err: %v\n", err)
	}
	//绑定参数
	util.BandingParams(&transParams)
	//发请求访问
	jsonByte := util.DoPost(transParams)
	//绑定返回的json
	err = json.Unmarshal(jsonByte, &res)
	if err != nil {
		log.Fatalf("json parse failed,err: %v\n", err)
	}
	c.JSON(http.StatusOK, res)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}