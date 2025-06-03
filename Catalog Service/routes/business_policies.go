package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type BusinessPoliciesRoutes struct {
}

func NewBusinessPoliciesRoutes() *BusinessPoliciesRoutes {
	return &BusinessPoliciesRoutes{}
}

func (r *BusinessPoliciesRoutes) Connect(endPoint string, business_policiesHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, business_policiesHandler.GET)
	echo.POST(endPoint, business_policiesHandler.POST)
	echo.PUT(endPoint, business_policiesHandler.PUT)
	echo.DELETE(endPoint, business_policiesHandler.DELETE)
	echo.GET(endPoint+"/:id", business_policiesHandler.GETBYID)
	echo.POST(endPoint+"/:multi", business_policiesHandler.MULTIPOST)
	return nil
}

