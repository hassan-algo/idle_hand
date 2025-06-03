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

	echo.GET(endPoint, bookingHandler.GET)
	echo.POST(endPoint, bookingHandler.POST)
	echo.PUT(endPoint, bookingHandler.PUT)
	echo.DELETE(endPoint, bookingHandler.DELETE)
	echo.GET(endPoint+"/:id", bookingHandler.GETBYID)
	echo.POST(endPoint+"/:multi", bookingHandler.MULTIPOST)
	return nil
}

