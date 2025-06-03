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

	echo.GET(endPoint, offersHandler.GET)
	echo.POST(endPoint, offersHandler.POST)
	echo.PUT(endPoint, offersHandler.PUT)
	echo.DELETE(endPoint, offersHandler.DELETE)
	echo.GET(endPoint+"/:id", offersHandler.GETBYID)
	echo.POST(endPoint+"/:multi", offersHandler.MULTIPOST)
	return nil
}

