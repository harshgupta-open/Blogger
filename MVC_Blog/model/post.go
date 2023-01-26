package model


//structure for post table
type Post struct {
	PostId      string `json:"postId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"user"`
	Created     string `json:"created"`
}