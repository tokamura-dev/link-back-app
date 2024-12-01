package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApiRouter() *gin.Engine {
	r := gin.Default()

	linkbackapp := r.Group("linkbackapp")
	{
		api := linkbackapp.Group("api")
		{
			v1 := api.Group("v1")
			{
				users := v1.Group("users")
				{
					// ユーザー情報全件取得API
					users.GET("/", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
					// ユーザーID指定のユーザー情報取得API
					users.GET("/:userid", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
					// ユーザー情報作成API
					users.POST("/", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
					// ユーザー情報更新API
					users.PUT("/:userid", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
					// ユーザー情報削除API
					users.DELETE("/:userid", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
				}
			}
		}
	}
	return r
}
