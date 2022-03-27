package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/guidoxie/myblog/docs"
	v1 "github.com/guidoxie/myblog/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	{
		tag := v1.NewTag()
		apiV1.POST("/tags", tag.Create)            // 新增
		apiV1.DELETE("/tags/:id", tag.Delete)      // 删除
		apiV1.PUT("/tags/:id", tag.Update)         // 更新
		apiV1.PATCH("/tags/:id/state", tag.Update) // 修改状态
		apiV1.GET("/tags", tag.List)               // 查询

		article := v1.NewArticle()
		apiV1.POST("/articles", article.Create)            // 新增
		apiV1.DELETE("/articles/:id", article.Delete)      // 删除
		apiV1.PUT("/articles/:id", article.Update)         // 更新
		apiV1.PATCH("/articles/:id/state", article.Update) // 修改状态
		apiV1.GET("/articles/:id", article.Get)            // 查询
		apiV1.GET("/articles", article.List)               // 查询
	}
	return r
}
