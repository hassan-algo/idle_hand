package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/labstack/echo/v4"
)

type PasswordManagementHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewPasswordManagementHandler() *PasswordManagementHandlers {
	return &PasswordManagementHandlers{}
}

func (h *PasswordManagementHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get password_management
// @Description
// @Produce json
// @Success 200 {object} structs.PasswordManagement "password_management"
// @Router /password_management [get]
// @Security ApiKeyAuth
// @Tags password_management
func (p *PasswordManagementHandlers) GET(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET PasswordManagement")
}

func (p *PasswordManagementHandlers) POST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "POST PasswordManagement")
}

func (p *PasswordManagementHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	password_management := structs.PasswordManagement{
		UserGuid:        data["user_guid"].(string),
		NewPassword:     data["new_password"].(string),
		ConfirmPassword: data["confirm_password"].(string),
	}
	message, err := p.apiBusiness.PUT(password_management)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, message)
}
func (p *PasswordManagementHandlers) DELETE(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "DELETE PasswordManagement")
}

func (p *PasswordManagementHandlers) GETBYID(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "GETBYID PasswordManagement")
}

func (p *PasswordManagementHandlers) MULTIPOST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	password_management := structs.PasswordManagement{
		Email: data["email"].(string),

	}
	message, err := p.apiBusiness.MULTIPOST(password_management)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, message)
}
