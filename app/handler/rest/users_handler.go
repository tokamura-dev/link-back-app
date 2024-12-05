package rest

import (
	"link-back-app/models"
	"link-back-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersHandler interface {
	GetAllUsersHandler(c *gin.Context)
	RegisterUsersHandler(c *gin.Context)
}

func NewUsersHandler(usecase usecase.UsersUsecase) UsersHandler {
	return &usersHandlerImpl{
		usecase: usecase,
	}
}

type usersHandlerImpl struct {
	usecase usecase.UsersUsecase
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersHandlerImpl) GetAllUsersHandler(c *gin.Context) {
	datas, err := u.usecase.GetAllUsersUsecase()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"datas": datas,
	})
}

/*
 * ユーザー情報登録処理
 */
func (u *usersHandlerImpl) RegisterUsersHandler(c *gin.Context) {
	var requestUsers models.RequestUsers
	if err := validator.New().Struct(&requestUsers); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
	}
}
