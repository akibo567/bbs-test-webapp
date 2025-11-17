package bbs

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	//"bbs-test-webapp/util"
	"bbs-test-webapp/server"

)

type Receive_Thread struct {
	Name string
	Title string
}

type Send_Thread struct {
	ID      int    `json:"id"`
	Title string `json:"title"`
	Name string `json:"name"`
}


func BBSrouting(){

	server.Router.POST("/get_threads", func(c *gin.Context) {
		rows, err := server.DB.Query(`SELECT id,title FROM threads 
		ORDER BY id DESC LIMIT 10`)
		if err != nil {
			log.Fatal(err)
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
				Title: title,
				Name: "PLACEHOLDER",
			})

		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"threads": bbs_threads,
		})
	})

	server.Router.POST("/post_threads", func(c *gin.Context) {

		var from_front Receive_Thread
		/*c.JSON(http.StatusOK, gin.H{
			"message": res_mes,
		})*/
		if c.ShouldBind(&from_front) != nil {
			
		}

		res, err := server.DB.Exec(`
			INSERT INTO threads (title, created_at,updated_at) VALUES ($1,NOW(),NOW())
		`,from_front.Title)
		if err != nil {
			//log.Printf(err.Error())
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