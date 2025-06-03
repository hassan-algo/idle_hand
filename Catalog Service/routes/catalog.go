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

	echo.GET(endPoint, catalogHandler.GET)
	echo.POST(endPoint, catalogHandler.POST)
	echo.PUT(endPoint, catalogHandler.PUT)
	echo.DELETE(endPoint, catalogHandler.DELETE)
	echo.GET(endPoint+"/:id", catalogHandler.GETBYID)
	echo.POST(endPoint+"/:multi", catalogHandler.MULTIPOST)
	return nil
}

