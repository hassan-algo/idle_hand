package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BookingRoutes struct {
}

func NewBookingRoutes() *BookingRoutes {
	return &BookingRoutes{}
}

func (r *BookingRoutes) Connect(endPoint string, bookingHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, auth.Authenticate(bookingHandler.GET))
	echo.POST(endPoint, auth.Authenticate(bookingHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(bookingHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(bookingHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(bookingHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(bookingHandler.MULTIPOST))
	return nil
}

