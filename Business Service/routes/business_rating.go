package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BusinessRatingRoutes struct {
}

func NewBusinessRatingRoutes() *BusinessRatingRoutes {
	return &BusinessRatingRoutes{}
}

func (r *BusinessRatingRoutes) Connect(endPoint string, business_ratingHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, business_ratingHandler.GET)
	echo.POST(endPoint, business_ratingHandler.POST)
	echo.PUT(endPoint, business_ratingHandler.PUT)
	echo.DELETE(endPoint, business_ratingHandler.DELETE)
	echo.GET(endPoint+"/:id", business_ratingHandler.GETBYID)
	echo.POST(endPoint+"/:multi", business_ratingHandler.MULTIPOST)
	return nil
}

