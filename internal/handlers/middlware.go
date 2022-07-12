package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) UserIdentity(ctx *gin.Context) {

	header := ctx.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "Missing auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid auth header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "Error ParseToken")
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int64, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int64)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "invalid user id")
		return 0, errors.New("invalid user id")
	}
	return idInt, nil
}
