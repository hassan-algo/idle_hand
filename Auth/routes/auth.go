package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
}

func NewAuthRoutes() *AuthRoutes {
	return &AuthRoutes{}
}

func (r *AuthRoutes) Connect(endPoint string, authHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {
	// echo.GET(endPoint, auth.Authenticate(authHandler.GET)) // to authenticate the user
	echo.GET(endPoint, authHandler.GET)
	echo.POST(endPoint, authHandler.POST)
	echo.PUT(endPoint, authHandler.PUT)
	echo.DELETE(endPoint, authHandler.DELETE)
	echo.GET(endPoint+"/:id", authHandler.GETBYID)
	echo.POST(endPoint+"/multi", authHandler.MULTIPOST)
	return nil
}

