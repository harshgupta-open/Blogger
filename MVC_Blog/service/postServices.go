package services

import (
	"context"

	db "github.com/blogger-api/db/initializer_db"
	datab "github.com/blogger-api/db/sqlc"

	// "github.com/blogger-api/model"
)

func GetallPostResult() ([]datab.Post,error) {
	//rows, err := db.DB.Query(`SELECT * FROM "Posts"`)

	ro := datab.New(db.DB)
	rows, err := ro.QueryGetAllPost(context.Background())

	if err != nil {
		return nil, err
	}

	// posts := make([]model.Post, 0)
	// for _,r:=range(rows) {
	// 	p := model.Post{}
	// 	err = rows.Scan(&p.PostId, &p.Title, &p.User, &p.Created, &p.Description)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, err

	// 	}
	//	posts = append(posts, p)
	//}
	return rows, nil

}
