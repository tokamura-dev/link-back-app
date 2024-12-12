package handler

import (
	"link-back-app/database"
	"link-back-app/domain/repository"
	"link-back-app/handler/rest"
	"link-back-app/usecase"

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
					repo := repository.NewUsersRepository(database.Connect())
					uc := usecase.NewUsersUsecase(repo)
					handler := rest.NewUsersHandler(uc)

					// ユーザーID指定のユーザー情報取得API
					users.GET("/get", handler.GetOneByEmployeeIdUsersHandler)
					// ユーザー情報全件取得API
					users.GET("/all_get", handler.GetAllUsersHandler)
					// ユーザー情報作成API
					users.POST("/", handler.RegisterUsersHandler)
					// ユーザー情報更新API
					users.PUT("/", handler.UpdateUsersHandler)
					// ユーザー情報論理削除API
					users.DELETE("/logical_delete/:passparam", handler.LogicalDeleteUsersHandler)
					// ユーザー情報削除API
					users.DELETE("/:passparam", handler.DeleteUsersHandler)
				}
			}
		}
	}
	return r
}
