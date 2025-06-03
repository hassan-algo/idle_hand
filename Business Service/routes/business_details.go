package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BusinessDetailsRoutes struct {
}

func NewBusinessDetailsRoutes() *BusinessDetailsRoutes {
	return &BusinessDetailsRoutes{}
}

func (r *BusinessDetailsRoutes) Connect(endPoint string, business_detailsHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, business_detailsHandler.GET)
	echo.POST(endPoint, business_detailsHandler.POST)
	echo.PUT(endPoint, business_detailsHandler.PUT)
	echo.DELETE(endPoint, business_detailsHandler.DELETE)
	echo.GET(endPoint+"/:id", business_detailsHandler.GETBYID)
	echo.POST(endPoint+"/:multi", business_detailsHandler.MULTIPOST)
	return nil
}

