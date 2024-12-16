package tokenutil

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
 * JWTトークンを生成する処理
 **/
func GenerateToken(employeeId string) (string, error) {
	tokenExpirationDate, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_DATE"))
	if err != nil {
		return "", nil
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["employee_id"] = employeeId
	claims["expiration"] = time.Now().Add(time.Hour * time.Duration(tokenExpirationDate)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

/**
 * トークン文字列を取得する処理
 **/
func extractTokenString(context *gin.Context) string {
	bearToken := context.Request.Header.Get("Authorization")
	strArray := strings.Split(bearToken, " ")
	if len(strArray) == 2 {
		return strArray[1]
	}
	return ""
}

/**
 * トークン文字列からトークンにパースする処理
 **/
func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

/**
 * トークンの検証処理
 **/
func TokenValid(context *gin.Context) error {
	tokenString := extractTokenString(context)
	token, err := parseToken(tokenString)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

/**
 * トークンから社員IDを取得する処理
 **/
func ExtractTokenId(context *gin.Context) (string, error) {
	tokenString := extractTokenString(context)
	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		employeeId, ok := claims["employee_id"].(string)
		if !ok {
			return "", nil
		}
		return string(employeeId), nil
	}
	return "", nil
}
