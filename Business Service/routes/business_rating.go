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

	echo.GET(endPoint, auth.Authenticate(business_ratingHandler.GET))
	echo.POST(endPoint, auth.Authenticate(business_ratingHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(business_ratingHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(business_ratingHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(business_ratingHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(business_ratingHandler.MULTIPOST))
	return nil
}

