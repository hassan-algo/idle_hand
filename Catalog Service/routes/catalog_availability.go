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

	echo.GET(endPoint, auth.Authenticate(catalog_availabilityHandler.GET))
	echo.POST(endPoint, auth.Authenticate(catalog_availabilityHandler.POST))
	echo.PUT(endPoint, auth.Authenticate(catalog_availabilityHandler.PUT))
	echo.DELETE(endPoint, auth.Authenticate(catalog_availabilityHandler.DELETE))
	echo.GET(endPoint+"/:id", auth.Authenticate(catalog_availabilityHandler.GETBYID))
	echo.POST(endPoint+"/:multi", auth.Authenticate(catalog_availabilityHandler.MULTIPOST))
	return nil
}

