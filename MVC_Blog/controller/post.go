package controller

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/blogger-api/db/initializer_db"
	datab "github.com/blogger-api/db/sqlc"

	// "github.com/blogger-api/model"
	services "github.com/blogger-api/service"
	"github.com/gin-gonic/gin"
)

func HomepageHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "we are on home page"})
}

func GetAllPost(c *gin.Context) {
	posts, err := services.GetallPostResult()
	if err != nil {
		log.Println(err)

	}
	c.IndentedJSON(http.StatusOK, posts)

}

func AddPost(c *gin.Context) {
	var p datab.InsertPostParams

	if err := c.ShouldBindJSON(&p); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in bindin json"})
	}
	ro := datab.New(db.DB)

	//insert := `insert into "Posts" ("postId","title","user","created","description") values ($1,$2,$3,$4,$5)`
	// d := time.Now().Format("2006-01-02 15:04:05")
	p.Created = sql.NullTime{time.Now(), true}
	// _, err := db.DB.Exec(insert, p.PostId, p.Title, p.User, d, p.Description)
	err := ro.InsertPost(context.Background(), p)
	if err != nil {
		log.Println("error in query")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in query", "error": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, p)
}
