package model

//structure for data table
type Comment struct {
	CommentId      string `json:"commentId"`
	CommentDesc    string `json:"commentDesc"`
	CommentPostId  string `json:"commentPostId"`
	CommentUser    string `json:"commentUser"`
	CommentCreated string `json:"commentCreated"`
}

