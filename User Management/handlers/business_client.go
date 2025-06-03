package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BusinessClientHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBusinessClientHandler() *BusinessClientHandlers {
	return &BusinessClientHandlers{}
}

func (h *BusinessClientHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get business_client
// @Description
// @Produce json
// @Success 200 {object} structs.BusinessClient "business_client"
// @Router /business_client [get]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessClientHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	business_client_guid := uuid.New().String()
	my_struct := structs.BusinessClient{
		BusinessClientGUID:     business_client_guid,
		ClientGUID:             data["client_guid"].(string),
		BusinessGUID:           data["business_guid"].(string),
		PreferredContactMethod: data["preferred_contact_method"].(string),
		Source:                 data["source"].(string),
		Note:                   data["note"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessClientHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessClient{
		BusinessClientGUID:     data["business_client_guid"].(string),
		ClientGUID:             data["client_guid"].(string),
		BusinessGUID:           data["business_guid"].(string),
		PreferredContactMethod: data["preferred_contact_method"].(string),
		Source:                 data["source"].(string),
		Note:                   data["note"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *BusinessClientHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessClient{
		BusinessClientGUID: data["business_client_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessClientHandlers) GETBYID(ctx echo.Context) error {
	business_client_guid := ctx.Param("business_client_guid")
	mydata, err := p.apiBusiness.GETBYID(business_client_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessClientHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessClient")
}
