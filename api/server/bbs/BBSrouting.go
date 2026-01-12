package bbs

import (
	"log"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"

	// "bbs-test-webapp/util"
	"bbs-test-webapp/server"
)

// フロントから受け取るスレッド書き込みの形式
type Receive_Thread struct {
	Name    string
	Title   string
	Message string
}

// バックエンドから送るスレッド一覧JSONの形式
type Send_Thread struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// フロントから受け取るメッセージ書き込みの形式
type Receive_Message struct {
	Name         string
	ThreadID     int
	Message_Text string
}

// バックエンドから送るスレッドのメッセージ形式
type Send_Message struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Message_Text string `json:"message_text"`
}

// メッセージ一覧取得時のスレッド情報取得の形式
type Receive_Thread_Info struct {
	ID int
}

func BBSrouting() {
	const messageLength = 17
	const nameMaxLength = 20
	const titleMaxLength = 20
	trimNewlines := func(message string) string {
		trimmed := strings.ReplaceAll(message, "\n", "")
		return strings.ReplaceAll(trimmed, "\r", "")
	}
	isNonEmptyWithin := func(value string, max int) bool {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		return utf8.RuneCountInString(trimmed) <= max
	}

	server.Router.POST("/bbs/get_threads", func(c *gin.Context) {
		rows, err := server.DB.Query(`SELECT id,title FROM threads 
		ORDER BY id DESC LIMIT 10`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var bbs_threads []Send_Thread

		for rows.Next() {
			var id int
			var title string

			if err := rows.Scan(&id, &title); err != nil {
				log.Fatal(err)
			}
			bbs_threads = append(bbs_threads, Send_Thread{
				ID:      id,
				Title:   title,
				Name:    "PLACEHOLDER",
				Message: "本文なし",
			})

		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"threads": bbs_threads,
		})
	})

	server.Router.POST("/bbs/post_thread", func(c *gin.Context) {

		var from_front Receive_Thread
		/*c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})*/
		if err := c.ShouldBind(&from_front); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		if utf8.RuneCountInString(trimNewlines(from_front.Message)) != messageLength {
			c.JSON(http.StatusBadRequest, gin.H{"error": "message must be 17 characters"})
			return
		}
		if !isNonEmptyWithin(from_front.Name, nameMaxLength) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name is required and must be 20 characters or less"})
			return
		}
		if !isNonEmptyWithin(from_front.Title, titleMaxLength) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title is required and must be 20 characters or less"})
			return
		}

		var insertedThreadID int

		err := server.DB.QueryRow(`
			INSERT INTO threads (title, created_at, updated_at)
			VALUES ($1, NOW(), NOW())
			RETURNING id
		`, from_front.Title).Scan(&insertedThreadID)

		if err != nil {
			// log.Printf(err.Error())
			// エラー内容を返して原因を掴む
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err2 := server.DB.Exec(`
			INSERT INTO thread_messages (name,message_text,thread_id,created_at, updated_at)
			VALUES ($1,$2,$3, NOW(), NOW());
		`, from_front.Name, from_front.Message, insertedThreadID)
		if err2 != nil {
			// log.Printf(err.Error())
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

	server.Router.POST("/bbs/get_messages", func(c *gin.Context) {

		var from_front Receive_Thread_Info
		if c.ShouldBind(&from_front) != nil {

		}

		// スレッド名を取得
		var threadTitle string
		err := server.DB.QueryRow(`SELECT title FROM threads WHERE id = $1`, from_front.ID).Scan(&threadTitle)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rows, err2 := server.DB.Query(`SELECT id,message_text,name FROM thread_messages
		WHERE thread_messages.thread_id = $1
		ORDER BY id ASC LIMIT 10`, from_front.ID)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
			return
		}

		var thread_messages []Send_Message

		for rows.Next() {
			var id int
			var message_text string
			var name string

			if err := rows.Scan(&id, &message_text, &name); err != nil {
				log.Fatal(err)
			}

			thread_messages = append(thread_messages, Send_Message{
				ID:           id,
				Name:         name,
				Message_Text: message_text,
			})

		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": thread_messages,
			"title":    threadTitle,
		})
	})

	server.Router.POST("/bbs/post_message", func(c *gin.Context) {

		var from_front Receive_Message
		/*c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})*/
		if err := c.ShouldBind(&from_front); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		if utf8.RuneCountInString(trimNewlines(from_front.Message_Text)) != messageLength {
			c.JSON(http.StatusBadRequest, gin.H{"error": "message must be 17 characters"})
			return
		}
		if !isNonEmptyWithin(from_front.Name, nameMaxLength) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name is required and must be 20 characters or less"})
			return
		}

		res, err := server.DB.Exec(`
			INSERT INTO thread_messages (name,message_text,thread_id,created_at, updated_at)
			VALUES ($1,$2,$3, NOW(), NOW());
		`, from_front.Name, from_front.Message_Text, from_front.ThreadID)
		if err != nil {
			// log.Printf(err.Error())
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
