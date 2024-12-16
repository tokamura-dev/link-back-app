package handler

import (
	"link-back-app/database"
	"link-back-app/domain/repository"
	"link-back-app/handler/rest"
	jwtauthmiddleware "link-back-app/middlewares/jwt_auth_middleware"
	"link-back-app/usecase"

	"github.com/gin-gonic/gin"
)

func GetApiRouter() *gin.Engine {
	r := gin.Default()

	dbConnect := database.Connect()

	authRepository := repository.NewAuthRepository(dbConnect)
	usersRepository := repository.NewUsersRepository(dbConnect)

	linkbackapp := r.Group("linkbackapp")
	{
		api := linkbackapp.Group("api")
		{
			v1 := api.Group("v1")
			{
				auth := v1.Group("auth")
				{
					authUsecase := usecase.NewAuthUsecase(dbConnect, authRepository, usersRepository)
					authHandler := rest.NewAuthHandler(authUsecase)

					// サインアップ処理
					auth.POST("/signup", authHandler.SignUpHandler)
					// サインイン処理
					auth.POST("/signin", authHandler.SignInHandler)
				}
				business := v1.Group("business")
				business.Use(jwtauthmiddleware.JwtAuthMiddleware())
				{
					users := business.Group("users")
					{
						usersUsecase := usecase.NewUsersUsecase(dbConnect, usersRepository)
						usersHandler := rest.NewUsersHandler(usersUsecase)

						// ユーザーID指定のユーザー情報取得API
						users.GET("/get", usersHandler.GetOneByEmployeeIdUsersHandler)
						// ユーザー情報全件取得API
						users.GET("/all", usersHandler.GetAllUsersHandler)
						// ユーザー情報作成API
						users.POST("/", usersHandler.RegisterUsersHandler)
						// ユーザー情報更新API
						users.PUT("/", usersHandler.UpdateUsersHandler)
						// ユーザー情報論理削除API
						users.DELETE("/logical_delete/:passparam", usersHandler.LogicalDeleteUsersHandler)
						// ユーザー情報削除API
						users.DELETE("/:passparam", usersHandler.DeleteUsersHandler)
					}
				}
			}

		}
	}
	return r
}
