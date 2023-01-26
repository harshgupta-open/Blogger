package initializer

import "github.com/blogger-api/controller"


func mapUrls(){
	router.GET("/",controller.HomepageHandler)
	 router.GET("/post",controller.GetAllPost)
	 router.GET("/post/:id", controller.GetCommentById)
	 router.POST("/post",controller.AddPost)
	 router.POST("/comment", controller.AddComment)
}