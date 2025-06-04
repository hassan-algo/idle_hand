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

	echo.GET(endPoint, auth.Authenticate(business_policiesHandler.GET))
	echo.POST(endPoint, auth.Authenticate(business_policiesHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(business_policiesHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(business_policiesHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(business_policiesHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(business_policiesHandler.MULTIPOST))
	return nil
}

