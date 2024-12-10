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
