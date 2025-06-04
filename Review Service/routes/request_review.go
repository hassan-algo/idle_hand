package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type RequestReviewRoutes struct {
}

func NewRequestReviewRoutes() *RequestReviewRoutes {
	return &RequestReviewRoutes{}
}

func (r *RequestReviewRoutes) Connect(endPoint string, request_reviewHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, auth.Authenticate(request_reviewHandler.GET))
	echo.POST(endPoint, auth.Authenticate(request_reviewHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(request_reviewHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(request_reviewHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(request_reviewHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(request_reviewHandler.MULTIPOST))
	return nil
}

