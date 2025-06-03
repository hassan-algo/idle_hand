package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type PasswordManagementRoutes struct {
}

func NewPasswordManagementRoutes() *PasswordManagementRoutes {
	return &PasswordManagementRoutes{}
}

func (r *PasswordManagementRoutes) Connect(endPoint string, password_managementHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, password_managementHandler.GET)
	echo.POST(endPoint, password_managementHandler.POST)
	echo.PUT(endPoint, password_managementHandler.PUT)
	echo.DELETE(endPoint, password_managementHandler.DELETE)
	echo.GET(endPoint+"/:id", password_managementHandler.GETBYID)
	echo.POST(endPoint+"/forgot_password", password_managementHandler.MULTIPOST)
	return nil
}

