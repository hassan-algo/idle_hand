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

func (r *AuthRoutes) Connect(endPoint string, AuthHandler apis.AuthHandler, echo *echo.Echo) error {
	echo.POST(endPoint, AuthHandler.Authentication)

	return nil
}
