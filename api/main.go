package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Receive_Chat_Message struct {
	Name    string
	Message string
}

type Send_Chat_Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // or "*"
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	/*router.POST("/ping2", func(c *gin.Context) {
		var from_front Ping2
		var res_mes string = "pong"

		if c.ShouldBind(&from_front) == nil {
			res_mes += from_front.TEST
		}

		c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})
	})*/

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

	router.POST("/get_chat_messages", func(c *gin.Context) {
		rows, err := db.Query(`SELECT id,name,message FROM test_kakikomi LIMIT 5`)
		if err != nil {
			log.Fatal(err)
		}

		
		var chat_messages []Send_Chat_Message

		for rows.Next() {
			var id int
			var name string
			var message string

			if err := rows.Scan(&id, &name, &message); err != nil {
				log.Fatal(err)
			}
			chat_messages = append(chat_messages, Send_Chat_Message{
				ID:      id,
				Name:    name,
				Message: message,
			})

		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": chat_messages,
		})
	})

	router.POST("/post_message", func(c *gin.Context) {

		var from_front Receive_Chat_Message
		var res_mes string = "pong"

		if c.ShouldBind(&from_front) == nil {
			res_mes += from_front.Message
		}

		/*c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})*/

		res, err := db.Exec(`
			INSERT INTO public.test_kakikomi (name, message) VALUES ($1, $2)
		`, from_front.Name, from_front.Message)
		if err != nil {
			log.Printf(err.Error())
			// エラー内容を返して原因を掴む
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}


		if n, _ := res.RowsAffected(); n == 0 {
			// トリガやルールのせいで 0 になることもあるが、基本は入らない合図
			log.Printf("insert affected 0 rows")
		}

		c.String(http.StatusOK, "ok")
	})

	router.Run(":" + port)

}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
