package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BusinessStaffRoutes struct {
}

func NewBusinessStaffRoutes() *BusinessStaffRoutes {
	return &BusinessStaffRoutes{}
}

func (r *BusinessStaffRoutes) Connect(endPoint string, business_staffHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, business_staffHandler.GET)
	echo.POST(endPoint, business_staffHandler.POST)
	echo.PUT(endPoint, business_staffHandler.PUT)
	echo.DELETE(endPoint, business_staffHandler.DELETE)
	echo.GET(endPoint+"/:id", business_staffHandler.GETBYID)
	echo.POST(endPoint+"/:multi", business_staffHandler.MULTIPOST)
	return nil
}

