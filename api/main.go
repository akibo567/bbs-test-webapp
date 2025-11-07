package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Ping2 struct {
	TEST string
}

func main() {
	port := getenv("PORT", "8080")
	dbURL := getenv("DATABASE_URL", "")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/ping2", func(c *gin.Context) {
		var from_front Ping2
		var res_mes string = "pong"

		if c.ShouldBind(&from_front) == nil {
			res_mes += from_front.TEST
		}
		c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})
	})

	router.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.String(http.StatusInternalServerError, "db not ready")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	// /api/... に合わせるなら、/apiグループを定義

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	router.Run(":" + port)

}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
