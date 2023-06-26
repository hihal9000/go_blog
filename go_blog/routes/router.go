package routes

import (
	v1 "go_blog/api/v1"
	"go_blog/middleware"
	"go_blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		//上传文件
		auth.POST("upload", v1.UpLoad)
	}
	routerv1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetUsers)
		// 分类模块的路由接口
		routerv1.GET("category", v1.GetCate)
		// 文章模块的路由接口
		routerv1.GET("article", v1.GetArt)
		routerv1.GET("article/list/:id", v1.GetCateArt)
		routerv1.GET("article/:id", v1.GetArtInfo)
		routerv1.POST("login", v1.Login)

	}
	_ = r.Run(utils.HttpPort)
}
