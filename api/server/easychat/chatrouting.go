package easychat

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	//"bbs-test-webapp/util"
	"bbs-test-webapp/server"

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


func Chatrouting(){

	server.Router.POST("/get_chat_messages", func(c *gin.Context) {
		rows, err := server.DB.Query(`SELECT id,name,message FROM test_kakikomi LIMIT 5`)
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

	server.Router.POST("/post_message", func(c *gin.Context) {

		var from_front Receive_Chat_Message
		var res_mes string = "pong"

		if c.ShouldBind(&from_front) == nil {
			res_mes += from_front.Message
		}

		/*c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})*/

		res, err := server.DB.Exec(`
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
}