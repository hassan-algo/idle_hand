package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type OffersRoutes struct {
}

func NewOffersRoutes() *OffersRoutes {
	return &OffersRoutes{}
}

func (r *OffersRoutes) Connect(endPoint string, offersHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, auth.Authenticate(offersHandler.GET))
	echo.POST(endPoint, auth.Authenticate(offersHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(offersHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(offersHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(offersHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(offersHandler.MULTIPOST))
	return nil
}

