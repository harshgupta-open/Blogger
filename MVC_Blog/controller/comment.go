package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/blogger-api/db/initializer_db"
	datab "github.com/blogger-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

func GetCommentById(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id does not contain any integer"})
	}
	po, err := postById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "post not found", "error": err})
		return
	}
	coms, err := getCommentById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Post not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, po)
	c.IndentedJSON(http.StatusOK, coms)
}

func getCommentById(id string) (*[]datab.Comment, error) {
	//query := `SELECT * FROM "Comments" where "commentPostId" = $1 `
	ro := datab.New(db.DB)

	//rows, err := db.DB.Query(query, id)
	i, err := strconv.ParseInt(id, 10, 64)

	i32 := int32(i)

	rows, err := ro.QueryGetCommentById(context.Background(), sql.NullInt32{i32, true})
	if err != nil {
		return nil, err
	}
	// defer rows.Close()

	// coms := make([]model.Comment, 0)
	// for rows.Next() {
	// 	c := model.Comment{}
	// 	err = rows.Scan(&c.CommentId, &c.CommentPostId, &c.CommentDesc, &c.CommentUser, &c.CommentCreated)
	// 	fmt.Println(c)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, err

	// 	}
	// 	coms = append(coms, c)
	// }
	return &rows, nil

}

// this a helper function which return post with certain post id
func postById(id string) (*datab.Post, error) {

	// if _, err := strconv.Atoi(id); err != nil {
	// 	return nil, fmt.Errorf("id does not containt integer")
	// }
	ro := datab.New(db.DB)

	//rows, err := db.DB.Query(query, id)
	i, err := strconv.ParseInt(id, 10, 64)

	i32 := int32(i)
	//query := `SELECT * FROM "Posts" where "postId" = $1 `
	//rows, err := db.DB.Query(query, id)

	p, err := ro.QueryGetPostById(context.Background(), i32)
	if err != nil {
		return nil, fmt.Errorf("post not found")
	}
	// var p model.Post
	// for rows.Next() {
	// 	err = rows.Scan(&p.PostId, &p.Title, &p.User, &p.Created, &p.Description)
	// 	fmt.Println(p)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, err

	// 	}
	// }

	return &p, nil

}

func AddComment(c *gin.Context) {
	var co datab.InsertCommentParams
	if err := c.ShouldBindJSON(&co); err != nil {

		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "binding error", "error": err})
		return
	}

	//po, err := postById(string(co.Commentpostid))
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "post not found", "error": err})
	// 	return
	// }
	ro := datab.New(db.DB)

	//insert := `insert into "Comments" ("commentId","commentPostId","commentDesc","commentUser","commentCreated") values ($1,$2,$3,$4,$5)`
	//d := time.Now().Format("2006-01-02 15:04:05")
	co.Commentcreated = time.Now()
	log.Println(co)
	//_, err = db.DB.Exec(insert, co.CommentId, co.CommentPostId, co.CommentDesc, co.CommentUser, d)
	err := ro.InsertComment(context.Background(), co)
	if err != nil {
		log.Println("error in query")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in query", "error": err})
		return
	}
	//c.IndentedJSON(http.StatusCreated, po)

	c.IndentedJSON(http.StatusCreated, co)
}
