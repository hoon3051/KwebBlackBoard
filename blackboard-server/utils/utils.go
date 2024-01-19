package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
)

// GetUser 함수는 Gin 컨텍스트에서 사용자 정보를 추출하고 검증합니다.
func GetUser(c *gin.Context) (*models.User, int, error) {
	userInterface, exists := c.Get("user")

	// 사용자 정보가 없는 경우, 오류를 반환합니다.
	if !exists {
		return nil, http.StatusUnauthorized, errors.New("Unauthorized")
	}

	user, ok := userInterface.(models.User)

	// 타입 변환이 실패한 경우, 오류를 반환합니다.
	if !ok {
		return nil, http.StatusInternalServerError, errors.New("Server error")
	}

	return &user, http.StatusOK, nil
}

// GetUintParam 함수는 Gin 컨텍스트에서 파라미터를 추출하고, uint 타입으로 변환합니다.
func GetUintParam(c *gin.Context, paramName string) (uint, error) {
	paramValue := c.Param(paramName)
	paramInt, err := strconv.ParseUint(paramValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(paramInt), nil
}
