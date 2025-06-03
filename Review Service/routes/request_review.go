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

	echo.GET(endPoint, request_reviewHandler.GET)
	echo.POST(endPoint, request_reviewHandler.POST)
	echo.PUT(endPoint, request_reviewHandler.PUT)
	echo.DELETE(endPoint, request_reviewHandler.DELETE)
	echo.GET(endPoint+"/:id", request_reviewHandler.GETBYID)
	echo.POST(endPoint+"/:multi", request_reviewHandler.MULTIPOST)
	return nil
}

