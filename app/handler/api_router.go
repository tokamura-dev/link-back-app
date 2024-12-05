package handler

import (
	"link-back-app/database"
	"link-back-app/domain/repository"
	"link-back-app/handler/rest"
	"link-back-app/usecase"

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
					repo := repository.NewUsersRepository(database.Connect())
					uc := usecase.NewUsersUsecase(repo)
					handler := rest.NewUsersHandler(uc)

					users.GET("/", handler.GetAllUsersHandler)
					// ユーザーID指定のユーザー情報取得API
					users.GET("/:userid", func(c *gin.Context) {
						c.JSON(http.StatusOK, gin.H{})
					})
					// ユーザー情報作成API
					users.POST("/", handler.RegisterUsersHandler)

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
