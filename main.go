package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


//glogal variable for database pointer
var db *sql.DB


//initializer for initializing database
func init() {
	var err error

	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")

}


//database details of my data
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "H@rsh@130"
	dbname   = "Blogger2"
)


//structure for post table data
type post struct {
	PostId      string `json:"postId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"user"`
	Created     string `json:"created"`
}


//structure for comment table data
type comment struct {
	CommentId      string `json:"commentId"`
	CommentDesc    string `json:"commentDesc"`
	CommentPostId  string `json:"commentPostId"`
	CommentUser    string `json:"commentUser"`
	CommentCreated string `json:"commentCreated"`
}


//getting all posts from post table through api
func allPost(c *gin.Context) {
	rows, err := db.Query(`SELECT * FROM "Posts"`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	posts := make([]post, 0)
	for rows.Next() {
		p := post{}
		err = rows.Scan(&p.PostId, &p.Title, &p.User, &p.Created, &p.Description)
		fmt.Println(p)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, err)
			return

		}
		posts = append(posts, p)
	}
	c.IndentedJSON(http.StatusOK, posts)

}


//getting all comment of a post with certain post id through api
func commentById(c *gin.Context) {
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

// this is a helper function which returns all comment with certain post id
func getCommentById(id string) (*[]comment, error) {
	query := `SELECT * FROM "Comments" where "commentPostId" = $1 `
	rows, err := db.Query(query, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	coms := make([]comment, 0)
	for rows.Next() {
		c := comment{}
		err = rows.Scan(&c.CommentId, &c.CommentPostId, &c.CommentDesc, &c.CommentUser, &c.CommentCreated)
		fmt.Println(c)
		if err != nil {
			log.Println(err)
			return nil, err

		}
		coms = append(coms, c)
	}
	return &coms, nil

}

//this a helper function which return post with certain post id
func postById(id string) (*post, error) {

	if _, err := strconv.Atoi(id); err != nil {
		return nil, fmt.Errorf("id does not containt integer")
	}
	query := `SELECT * FROM "Posts" where "postId" = $1 `
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("post not found")
	}
	var p post
	for rows.Next() {
		err = rows.Scan(&p.PostId, &p.Title, &p.User, &p.Created, &p.Description)
		fmt.Println(p)
		if err != nil {
			log.Println(err)
			return nil, err

		}
	}

	return &p, nil

}


//this  function will add post to database through api
func addPost(c *gin.Context) {
	var p post

	if err := c.BindJSON(&p); err != nil {
		return
	}
	insert := `insert into "Posts" ("postId","title","user","created","description") values ($1,$2,$3,$4,$5)`
	d := time.Now().Format("2006-01-02 15:04:05")
	p.Created = d
	_, err := db.Exec(insert, p.PostId, p.Title, p.User, d, p.Description)
	if err != nil {
		log.Println("error in query")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in query", "error": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, p)
}


//adding comment to a post through api
func addComment(c *gin.Context) {
	var co comment
	if err := c.BindJSON(&co); err != nil {

		c.IndentedJSON(http.StatusBadGateway, err)
		return
	}
	if _, err := strconv.Atoi(co.CommentPostId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id does not contain any integer"})
	}
	po, err := postById(co.CommentPostId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "post not found", "error": err})
		return
	}
	insert := `insert into "Comments" ("commentId","commentPostId","commentDesc","commentUser","commentCreated") values ($1,$2,$3,$4,$5)`
	d := time.Now().Format("2006-01-02 15:04:05")
	co.CommentCreated = d
	log.Println(co)
	_, err = db.Exec(insert, co.CommentId, co.CommentPostId, co.CommentDesc, co.CommentUser, d)
	if err != nil {
		log.Println("error in query")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in query", "error": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, po)

	c.IndentedJSON(http.StatusCreated, co)
}


//this is home page api function just for testing
func homepageHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "We are good to go."})
}


//main function with all router path and gin
func main() {
	router := gin.Default()
	router.GET("/", homepageHandler)
	router.GET("/post", allPost)
	router.GET("/post/:id", commentById)
	router.POST("/post", addPost)
	router.POST("/comment", addComment)
	router.Run("localhost:8080")
	defer db.Close()

}


//a function that check error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
