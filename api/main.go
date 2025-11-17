package main

import (
	//"database/sql"
	//"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"bbs-test-webapp/util"
	"bbs-test-webapp/server/easychat"
	"bbs-test-webapp/server"
)

func main() {
	port := util.Getenv("PORT", "8080")

    server.InitDB()
    server.InitRouter()

	server.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // or "*"
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	server.Router.GET("/health", func(c *gin.Context) {
		if err := server.DB.Ping(); err != nil {
			c.String(http.StatusInternalServerError, "db not ready")
			return
		}
		c.String(http.StatusOK, "ok")
	})


	server.Router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "heLあo",
		})
	})

	easychat.Chatrouting();



	

	server.Router.Run(":" + port)

}