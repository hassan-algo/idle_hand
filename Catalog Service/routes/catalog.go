package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type CatalogRoutes struct {
}

func NewCatalogRoutes() *CatalogRoutes {
	return &CatalogRoutes{}
}

func (r *CatalogRoutes) Connect(endPoint string, catalogHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, auth.Authenticate(catalogHandler.GET))
	echo.POST(endPoint, auth.Authenticate(catalogHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(catalogHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(catalogHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(catalogHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(catalogHandler.MULTIPOST))
	return nil
}

