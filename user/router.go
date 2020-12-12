package user

import (
	"github.com/gin-gonic/gin"
)

func RegistUserEndpoints(router *gin.RouterGroup) {
	router.GET("/users", findAllUser)
	router.GET("/user/:id", findUser)
	router.POST("/user", saveUser)
	router.PUT("/user/:id", updateUser)
	router.PUT("/user/:id", deleteUser)
}

func findAllUser(context *gin.Context) {

}

func findUser(context *gin.Context) {

}

func saveUser(context *gin.Context) {

}

func updateUser(context *gin.Context) {

}

func deleteUser(context *gin.Context) {

}
