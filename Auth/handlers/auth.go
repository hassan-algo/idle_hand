package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/structs"
	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewAuthHandler() *AuthHandlers {
	return &AuthHandlers{}
}

func (h *AuthHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get auth
// @Description
// @Produce json
// @Success 200 {object} structs.Auth "auth"
// @Router /auth [get]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) GET(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Get Auth", Valid: true, Data: nil})
}

// @Summary Post auth
// @Description
// @Produce json
// @Param  product body structs.Auth true "Post auth"
// @Success 200 {object} structs.Response
// @Router /auth [post]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) POST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Post Auth", Valid: true, Data: nil})
}

// @Summary Put auth
// @Description
// @Produce json
// @Param  product body structs.Auth true "Put auth"
// @Success 200 {object} structs.Response
// @Router /auth [put]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) PUT(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Put Auth", Valid: true, Data: nil})
}

// @Summary Delete auth
// @Description
// @Produce json
// @Param  product body structs.Auth true "Delete auth"
// @Success 200 {object} structs.Response
// @Router /auth [delete]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) DELETE(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Delete Auth", Valid: true, Data: nil})
}

// @Summary Get auth by id
// @Description
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} structs.Auth "auth"
// @Router /auth/{id} [get]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) GETBYID(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Get auth by id", Valid: true, Data: nil})
}

// @Summary Multipost auth
// @Description
// @Produce json
// @Param  product body structs.Auths true "Multipost auth"
// @Success 200 {object} structs.Response
// @Router /auth/multi [post]
// @Security ApiKeyAuth
// @Tags auth
func (h *AuthHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, structs.Response{Message: "Multipost Auth", Valid: true, Data: nil})
}

