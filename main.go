package main

import (
	"gin-example/global"
)

func main() {

	global.InitializeDatabase()
	global.InitializeRouters()


}