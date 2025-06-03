package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type OffersHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewOffersHandler() *OffersHandlers {
	return &OffersHandlers{}
}

func (h *OffersHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get offers
// @Description
// @Produce json
// @Success 200 {object} structs.Offers "offers"
// @Router /offers [get]
// @Security ApiKeyAuth
// @Tags offers
func (p *OffersHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *OffersHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	offer_guid := uuid.New().String()
	my_struct := structs.Offers{
		OfferGUID:    offer_guid,
		OfferTitle:   data["offer_title"].(string),
		OfferMessage: data["offer_message"].(string),
		CustomerType: data["customer_type"].(string),
		Medium:       data["medium"].(int),
		BusinessGUID: data["business_guid"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *OffersHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.Offers{
		OfferGUID:    data["offer_guid"].(string),
		OfferTitle:   data["offer_title"].(string),
		OfferMessage: data["offer_message"].(string),
		CustomerType: data["customer_type"].(string),
		Medium:       data["medium"].(int),
		BusinessGUID: data["business_guid"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *OffersHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.Offers{
		OfferGUID: data["offer_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *OffersHandlers) GETBYID(ctx echo.Context) error {
	offer_guid := ctx.Param("offer_guid")
	mydata, err := p.apiBusiness.GETBYID(offer_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *OffersHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST Offers")
}
