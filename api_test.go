package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)


//setting up router 
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}


//testing homepage api func
func TestHomepageHandler(t *testing.T) {
	r := SetUpRouter()
	mockResponse := "We are good to go."
	r.GET("/", homepageHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Result().Body)
	x := map[string]string{}
	json.Unmarshal([]byte(responseData), &x)
	// log.Println(string(responseData),"\n",strings.TrimSpace(string(responseData)))
	// log.Println(responseData, "\n", string(bytes.TrimSpace(responseData)), "\n", x, "\n", w.Code)
	log.Println(x)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, x["message"])

}

//testing postById func for id should not be string char
func TestPostById(t *testing.T) {
	s:="5"
	_,err:=postById(s)
	if err!=nil{
		t.Fatalf("error sended by postById func")
	}
	if _,err:=strconv.Atoi(s);err!=nil{
		t.Fatalf("id does not containt integer")
	}

}

// testing allPost api function
func TestAllPost(t *testing.T) {
	r := SetUpRouter()
	r.GET("/post", allPost)
	req, _ := http.NewRequest("GET", "/post", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)


}

//testing commentById api function
func TestCommentByID(t *testing.T) {
	r := SetUpRouter()
	s:="5"
	r.GET("/post/:id", commentById)
	req, _ := http.NewRequest("GET", "/post/"+s, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	_,err:=strconv.Atoi(s)
	assert.Equal(t,err,nil)
	assert.Equal(t, http.StatusOK, w.Code)
	

}


//testing addPost api function
func TestAddPost(t *testing.T) {
	r := SetUpRouter()
	r.POST("/post", addPost)
	p:=post{
		PostId: "11",
		Title: "vgfshfvysh",
		User: "giri",
		Created: time.Now().Format("2006-01-02 15:04:05"),
		Description: "sbaj,hdbajnakjojdij",
	}
	assert.NotEqual(t,p.PostId,"")
	assert.NotEmpty(t,p.Title)
	assert.NotEmpty(t,p.Description)
	jsonValue,_:=json.Marshal(p)

	req, _ := http.NewRequest("POST", "/post",bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

//testing addComment api function
func TestAddCommentt(t *testing.T) {
	r := SetUpRouter()
	r.POST("/post", addComment)
	c:=comment{
		CommentId: "13",
		CommentPostId: "10",
		CommentDesc: "hgshgsfjhdjkMDH",
		CommentUser: "UJJWAL",
		CommentCreated: time.Now().Format("2006-01-02 15:04:05"),
	}
	assert.NotEqual(t,c.CommentId,"")
	assert.NotEmpty(t,c.CommentPostId)
	assert.NotEmpty(t,c.CommentDesc)
	jsonValue,_:=json.Marshal(c)

	req, _ := http.NewRequest("POST", "/comment",bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}
