package main

import (
	"net/http"

	"github.com/CAATHARSIS/absolute_cinema/common/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	database.Init(dbUrl)

	router.LoadHTMLGlob("internal/templates/*")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	router.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"port":  port,
				"dbUrl": dbUrl,
			})
	})

	router.Run(port)
}
