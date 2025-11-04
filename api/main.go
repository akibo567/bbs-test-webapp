package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	/*
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			if err := db.Ping(); err != nil {
				http.Error(w, "db not ready", 500)
				return
			}
			fmt.Fprintln(w, "ok")
		})

		// 実API: /api/... に合わせる
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, `{"message":"hello!!!"}`)
		})

		log.Printf("listening on :%s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	*/
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
