package initializer

import (
	"fmt"

	db "github.com/blogger-api/db/initializer_db"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func RouterIntialize() {
	fmt.Println("Starting Blogger API")
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	//LoadEnvVariable()
	db.ConnectDataBase()
	mapUrls()
	router.Run(":8080")

}
