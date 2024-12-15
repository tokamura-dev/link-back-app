package rest

import (
	"link-back-app/api"
	usersmodel "link-back-app/models/users_model"
	"link-back-app/usecase"
	stringutil "link-back-app/utils/string_util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	GetOneByEmployeeIdUsersHandler(context *gin.Context)
	GetAllUsersHandler(context *gin.Context)
	RegisterUsersHandler(context *gin.Context)
	UpdateUsersHandler(context *gin.Context)
	LogicalDeleteUsersHandler(context *gin.Context)
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
 * ユーザーID1件取得(社員ID指定)
 **/
func (u *usersHandlerImpl) GetOneByEmployeeIdUsersHandler(context *gin.Context) {
	var requestOneByEmployeeIdUsers usersmodel.RequestOneByEmployeeIdUsers
	if err := context.ShouldBindQuery(&requestOneByEmployeeIdUsers); err != nil {
		api.ErrorrResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	data, err := u.usecase.GetOneByEmployeeIdUsersUsecase(requestOneByEmployeeIdUsers.EmployeeId)
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	api.SuccessResponse(context, http.StatusOK, data)
}

/**
 * ユーザー情報全件取得処理
 */
func (u *usersHandlerImpl) GetAllUsersHandler(context *gin.Context) {
	// ユーザー情報を全件取得
	datas, err := u.usecase.GetAllUsersUsecase()
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	api.SuccessResponse(context, http.StatusOK, datas)
}

/*
 * ユーザー情報登録処理
 */
func (u *usersHandlerImpl) RegisterUsersHandler(context *gin.Context) {
	var requestCreateUsers usersmodel.RequestCreateUsers
	if err := context.ShouldBindJSON(&requestCreateUsers); err != nil {
		api.ErrorrResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	// ユーザー情報登録処理
	err := u.usecase.RegisterUsersUsecase(requestCreateUsers)
	if err != nil {
		api.ErrorrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	// レスポンス処理
	context.JSON(http.StatusCreated, gin.H{})
}

/**
 * ユーザー情報更新処理
 **/
func (u *usersHandlerImpl) UpdateUsersHandler(context *gin.Context) {
	var users usersmodel.Users
	if err := context.ShouldBindJSON(&users); err != nil {
		api.ErrorrResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	// ユーザー情報更新処理
	err := u.usecase.UpdateUsersUsecase(users)
	if err != nil {
		api.ErrorrResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	// レスポンス処理
	context.JSON(http.StatusNoContent, gin.H{})
}

/**
 * ユーザー情報論理削除処理
 **/
func (u *usersHandlerImpl) LogicalDeleteUsersHandler(context *gin.Context) {
	passparam := strings.Split(context.Param("passparam"), "_")
	if !stringutil.ValidPassParam(passparam, usersmodel.USERS_PRIMARY_KEY_COUNT) {
		api.ErrorrResponse(context, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	deleteTargetEmployeeId := passparam[0]

	// 社員IDに紐づくユーザー情報を論理削除
	err := u.usecase.LogicalDeleteUsersUsecase(deleteTargetEmployeeId)
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	context.JSON(http.StatusNoContent, gin.H{})
}

/**
 * ユーザー情報削除処理
 **/
func (u *usersHandlerImpl) DeleteUsersHandler(context *gin.Context) {
	passparam := strings.Split(context.Param("passparam"), "_")
	if !stringutil.ValidPassParam(passparam, usersmodel.USERS_PRIMARY_KEY_COUNT) {
		api.ErrorrResponse(context, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	deleteTargetEmployeeId := passparam[0]

	// 社員IDに紐づくユーザー情報を物理削除
	err := u.usecase.DeleteUsersUsecase(deleteTargetEmployeeId)
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	context.JSON(http.StatusNoContent, gin.H{})
}
