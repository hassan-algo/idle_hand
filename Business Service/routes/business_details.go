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

	echo.GET(endPoint, auth.Authenticate(business_detailsHandler.GET))
	echo.POST(endPoint, auth.Authenticate(business_detailsHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(business_detailsHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(business_detailsHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(business_detailsHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(business_detailsHandler.MULTIPOST))
	return nil
}

