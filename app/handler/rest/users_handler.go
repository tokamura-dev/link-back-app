package rest

import (
	"fmt"
	"link-back-app/api"
	"link-back-app/models"
	"link-back-app/usecase"
	stringutil "link-back-app/utils/string_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	GetAllUsersHandler(context *gin.Context)
	RegisterUsersHandler(context *gin.Context)
	DeleteUsersHandler(context *gin.Context)
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
func (u *usersHandlerImpl) GetAllUsersHandler(context *gin.Context) {
	// ユーザー情報を全件取得
	datas, err := u.usecase.GetAllUsersUsecase()
	if err != nil {
		api.Response(context, http.StatusInternalServerError, nil, err.Error())
		return
	}
	api.Response(context, http.StatusOK, datas, "")
}

/*
 * ユーザー情報登録処理
 */
func (u *usersHandlerImpl) RegisterUsersHandler(context *gin.Context) {
	var requestUsers models.RequestUsers
	if err := context.ShouldBindJSON(&requestUsers); err != nil {
		api.Response(context, http.StatusBadRequest, nil, err.Error())
		return
	}
	// ユーザー情報登録処理
	err := u.usecase.RegisterUsersUsecase(requestUsers)
	if err != nil {
		api.Response(context, http.StatusInternalServerError, nil, err.Error())
		return
	}
	// レスポンス処理
	context.JSON(http.StatusCreated, gin.H{})
}

/**
 * ユーザー情報削除処理
 **/
func (u *usersHandlerImpl) DeleteUsersHandler(context *gin.Context) {
	deleteTargetEmployeeId := context.Param("employeeid")
	if stringutil.IsEmpty(deleteTargetEmployeeId) {
		api.Response(context, http.StatusBadRequest, nil, "有効な値が指定されていません")
	}
	fmt.Println(deleteTargetEmployeeId)
}
