package handler

import (
	"net/http"

	"github.com/tepasha/golang_blank_project/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary      Update user
// @Tags         auth
// @Description  Update user
// @ID           update-user
// @Accept       json
// @Produce      json
// @Param        id       path      int  true  "Update user"
// @Param        input    body      models.UpdateUserInput  true  "Update user"
// @Success      200      {object}  statusResponse
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/usersettings/{id} [put]
func (h *Handler) changeUserInfo(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input models.UpdateUserInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateUser(userId, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{"ok"})
}
