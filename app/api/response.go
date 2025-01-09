package api

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponseStruct struct {
	SearchCount  int         `json:"searchCount"`
	SearchResult interface{} `json:"searchResult"`
}

type ErrorResponseStruct struct {
	Message string `json:"message"`
}

type AuthResponseStruct struct {
	AuthResult bool   `json:"authResult"`
	JwtToken   string `json:"jwtToken"`
	Message    string `json:"message"`
}

/**
 * 成功時レスポンス
 **/
func SuccessResponse(context *gin.Context, statusCode int, datas []interface{}) {
	context.JSON(statusCode, SuccessResponseStruct{
		SearchCount:  len(datas),
		SearchResult: datas,
	})
}

/**
 * エラー時レスポンス
 **/
func ErrorrResponse(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, ErrorResponseStruct{
		Message: message,
	})
}

/**
 * 認証時レスポンス
 **/
func AuthResponse(context *gin.Context, statusCode int, authResult bool, jwtToken string, message string) {
	context.JSON(statusCode, AuthResponseStruct{
		AuthResult: authResult,
		JwtToken:   jwtToken,
		Message:    message,
	})
}
