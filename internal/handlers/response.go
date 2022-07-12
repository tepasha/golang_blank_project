package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type (
	errorResponse struct {
		Message string `json:"message"`
	}

	statusResponse struct {
		Status string `json:"status"`
	}

	idResponse struct {
		ID int64 `json:"id"`
	}

	singInInput struct {
		Login    string `json:"login" binding:"required"`
		PassWord string `json:"pass" binding:"required"`
	}

	getUserToken struct {
		Token string `json:"token"`
	}
)

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
