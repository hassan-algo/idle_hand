package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CatalogHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewCatalogHandler() *CatalogHandlers {
	return &CatalogHandlers{}
}

func (h *CatalogHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get catalog
// @Description
// @Produce json
// @Success 200 {object} structs.Catalog "catalog"
// @Router /catalog [get]
// @Security ApiKeyAuth
// @Tags catalog
func (p *CatalogHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	catalog_guid := uuid.New().String()
	my_struct := structs.Catalog{
		CatalogGUID:        catalog_guid,
		BusinessGUID:       data["business_guid"].(string),
		CatalogCategory:    data["catalog_category"].(string),
		CatalogName:        data["catalog_name"].(string),
		CatalogDescription: data["catalog_description"].(string),
		CatalogPrice:       data["catalog_price"].(string),
		CatalogOffering:    data["catalog_offering"].(string),
		CatalogPhoto:       data["catalog_photo"].(string),
		AssignedStaffGUID:  data["assigned_staff_guid"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	my_struct := structs.Catalog{
		CatalogGUID:        data["catalog_guid"].(string),
		BusinessGUID:       data["business_guid"].(string),
		CatalogCategory:    data["catalog_category"].(string),
		CatalogName:        data["catalog_name"].(string),
		CatalogDescription: data["catalog_description"].(string),
		CatalogPrice:       data["catalog_price"].(string),
		CatalogOffering:    data["catalog_offering"].(string),
		CatalogPhoto:       data["catalog_photo"].(string),
		AssignedStaffGUID:  data["assigned_staff_guid"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogHandlers) GETBYID(ctx echo.Context) error {
	catalog_guid := ctx.Param("catalog_guid")
	mydata, err := p.apiBusiness.GETBYID(catalog_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST Catalog")
}

func (p *CatalogHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.Catalog{
		CatalogGUID: data["catalog_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
