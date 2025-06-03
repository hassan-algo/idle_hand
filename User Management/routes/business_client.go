package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BusinessClientRoutes struct {
}

func NewBusinessClientRoutes() *BusinessClientRoutes {
	return &BusinessClientRoutes{}
}

func (r *BusinessClientRoutes) Connect(endPoint string, business_clientHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, business_clientHandler.GET)
	echo.POST(endPoint, business_clientHandler.POST)
	echo.PUT(endPoint, business_clientHandler.PUT)
	echo.DELETE(endPoint, business_clientHandler.DELETE)
	echo.GET(endPoint+"/:id", business_clientHandler.GETBYID)
	echo.POST(endPoint+"/:multi", business_clientHandler.MULTIPOST)
	return nil
}

