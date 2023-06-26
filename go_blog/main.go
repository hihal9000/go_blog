package main

import (
	"go_blog/model"
	"go_blog/routes"
)

func main() {
	// 引用数据库

	model.InitDb()
	routes.InitRouter()
}
