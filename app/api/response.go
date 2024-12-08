package api

import (
	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	SearchCount  int         `json:"searchCount"`
	SearchResult interface{} `json:"searchResult"`
	Message      string      `json:"message"`
}

func Response(context *gin.Context, statusCode int, datas []interface{}, message string) {
	context.JSON(statusCode, ResponseStruct{
		SearchCount:  len(datas),
		SearchResult: datas,
		Message:      message,
	})
}
