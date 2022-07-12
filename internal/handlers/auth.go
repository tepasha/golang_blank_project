package handler

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/tepasha/golang_blank_project/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      SignUp
// @Tags         auth
// @Description  Create account
// @ID           create-account
// @Accept       json
// @Produce      json
// @Param        input    body       models.User  true  "account info"
// @Success      200      {integer}  integer      1
// @Failure      400,404  {object}   errorResponse
// @Failure      500      {object}   errorResponse
// @Failure      default  {object}   errorResponse
// @Router       /api/auth/sing-up [post]
func (h *Handler) singUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.AvatarSrc.Mime != "" && input.AvatarSrc.Data != "" {
		var osFile *os.File
		var fileSrc string

		unbased, _ := base64.StdEncoding.DecodeString(input.AvatarSrc.Data)

		res := bytes.NewReader(unbased)

		filePath := "./src/avatar/" + "avtr_" + uuid.New().String()

		switch input.AvatarSrc.Mime {
		case "image/png":
			pngI, errPng := png.Decode(res)
			if errPng == nil {
				fileSrc = filePath + ".png"
				osFile, _ = os.OpenFile(fileSrc, os.O_WRONLY|os.O_CREATE, 0777)
				png.Encode(osFile, pngI)
			} else {
				newErrorResponse(ctx, http.StatusInternalServerError, errPng.Error())
			}
		case "image/jpeg":
			jpgI, errJpg := jpeg.Decode(res)
			if errJpg == nil {
				fileSrc = filePath + ".jpg"
				osFile, _ = os.OpenFile(fileSrc, os.O_WRONLY|os.O_CREATE, 0777)
				jpeg.Encode(osFile, jpgI, &jpeg.Options{Quality: 75})
			} else {
				newErrorResponse(ctx, http.StatusInternalServerError, errJpg.Error())
			}
		}

		defer func() {
			osFile.Close()
		}()

		input.Avatar = fileSrc
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, idResponse{
		ID: id,
	})

}

// @Summary      SignIn
// @Tags         auth
// @Description  login
// @ID           login
// @Accept       json
// @Produce      json
// @Param        input    body      singInInput  true  "credentials"
// @Success      200      {object}  getUserToken
// @Failure      400,404  {object}  errorResponse
// @Failure      500      {object}  errorResponse
// @Failure      default  {object}  errorResponse
// @Router       /api/auth/sing-in [post]
func (h *Handler) singIn(ctx *gin.Context) {
	var input singInInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Login, input.PassWord)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getUserToken{
		Token: token,
	})
}
