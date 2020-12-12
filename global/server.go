package global

import (
	"gin-example/user"
	"github.com/gin-gonic/gin"
)

func InitializeRouters() {

	r := gin.Default()

	api := r.Group("/api")
	user.RegistUserEndpoints(api.Group(""))


}

