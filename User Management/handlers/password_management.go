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

// @Summary Get password management status
// @Description Get the current password management status
// @Produce json
// @Success 200 {object} structs.PasswordManagement "Password management status"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management [get]
// @Security ApiKeyAuth
// @Tags password_management
func (p *PasswordManagementHandlers) GET(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET PasswordManagement")
}

// @Summary Create password management
// @Description Create a new password management entry
// @Accept json
// @Produce json
// @Param password_management body structs.PasswordManagement true "Password Management Object"
// @Success 200 {object} structs.PasswordManagement "Created password management"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management [post]
// @Security ApiKeyAuth
// @Tags password_management
func (p *PasswordManagementHandlers) POST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "POST PasswordManagement")
}

// @Summary Update password
// @Description Update user password
// @Accept json
// @Produce json
// @Param password_management body structs.PasswordManagement true "Password Management Object"
// @Success 200 {object} structs.PasswordManagement "Updated password management"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 404 {object} handlers.ErrorResponse "Not Found"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management [put]
// @Security ApiKeyAuth
// @Tags password_management
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

// @Summary Delete password management
// @Description Delete a password management entry
// @Accept json
// @Produce json
// @Param password_management body structs.PasswordManagement true "Password Management Object with GUID"
// @Success 200 {object} structs.PasswordManagement "Deleted password management"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 404 {object} handlers.ErrorResponse "Not Found"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management [delete]
// @Security ApiKeyAuth
// @Tags password_management
func (p *PasswordManagementHandlers) DELETE(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "DELETE PasswordManagement")
}

// @Summary Get password management by ID
// @Description Get a specific password management entry by ID
// @Produce json
// @Param password_management_guid path string true "Password Management GUID"
// @Success 200 {object} structs.PasswordManagement "Password management details"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 404 {object} handlers.ErrorResponse "Not Found"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management/{password_management_guid} [get]
// @Security ApiKeyAuth
// @Tags password_management
func (p *PasswordManagementHandlers) GETBYID(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GETBYID PasswordManagement")
}

// @Summary Reset password by email
// @Description Reset password using email address
// @Accept json
// @Produce json
// @Param password_management body structs.PasswordManagement true "Password Management Object with Email"
// @Success 200 {object} structs.PasswordManagement "Password reset initiated"
// @Failure 400 {object} handlers.ErrorResponse "Bad Request"
// @Failure 500 {object} handlers.ErrorResponse "Internal Server Error"
// @Router /password_management/multi [post]
// @Security ApiKeyAuth
// @Tags password_management
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
