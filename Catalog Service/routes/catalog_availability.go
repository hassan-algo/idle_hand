package routes

import (
	"example.com/apis"
	"github.com/labstack/echo/v4"
)

type CatalogAvailabilityRoutes struct {
}

func NewCatalogAvailabilityRoutes() *CatalogAvailabilityRoutes {
	return &CatalogAvailabilityRoutes{}
}

func (r *CatalogAvailabilityRoutes) Connect(endPoint string, catalog_availabilityHandler apis.APIHandler, echo *echo.Echo, auth apis.AuthHandler) error {

	echo.GET(endPoint, catalog_availabilityHandler.GET)
	echo.POST(endPoint, catalog_availabilityHandler.POST)
	echo.PUT(endPoint, catalog_availabilityHandler.PUT)
	echo.DELETE(endPoint, catalog_availabilityHandler.DELETE)
	echo.GET(endPoint+"/:id", catalog_availabilityHandler.GETBYID)
	echo.POST(endPoint+"/:multi", catalog_availabilityHandler.MULTIPOST)
	return nil
}

