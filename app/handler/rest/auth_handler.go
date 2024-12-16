package rest

import (
	"link-back-app/api"
	authmodel "link-back-app/models/auth_model"
	"link-back-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignUpHandler(context *gin.Context)
	SignInHandler(context *gin.Context)
}

func NewAuthHandler(usecase usecase.AuthUsecase) AuthHandler {
	return &authHandlerImpl{
		usecase: usecase,
	}
}

type authHandlerImpl struct {
	usecase usecase.AuthUsecase
}

/**
 * サインアップ処理
 **/
func (a *authHandlerImpl) SignUpHandler(context *gin.Context) {
	var requestSignUp authmodel.RequestSignUp
	if err := context.ShouldBindJSON(&requestSignUp); err != nil {
		api.ErrorrResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	// サインアップ処理
	err := a.usecase.SignUpUsacase(requestSignUp)
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	context.JSON(http.StatusCreated, gin.H{})
}

/**
 * サインイン処理
 **/
func (a *authHandlerImpl) SignInHandler(context *gin.Context) {
	var requestSignIn authmodel.RequstSignIn
	if err := context.ShouldBindJSON(&requestSignIn); err != nil {
		api.ErrorrResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	// サインイン処理
	token, err := a.usecase.SignInUsecase(requestSignIn)
	if err != nil {
		if apiError, ok := err.(*api.ApiError); ok {
			api.ErrorrResponse(context, apiError.HttpStatusCode, apiError.ErrorMessage)
		}
		return
	}
	// レスポンス処理
	context.JSON(http.StatusOK, gin.H{"jwt-token": token})
}
