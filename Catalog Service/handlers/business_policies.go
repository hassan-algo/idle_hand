package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BusinessPoliciesHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBusinessPoliciesHandler() *BusinessPoliciesHandlers {
	return &BusinessPoliciesHandlers{}
}

func (h *BusinessPoliciesHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get business_policies
// @Description
// @Produce json
// @Success 200 {object} structs.BusinessPolicies "business_policies"
// @Router /business_policies [get]
// @Security ApiKeyAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessPoliciesHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	business_policies_guid := uuid.New().String()
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      business_policies_guid,
		BusinessGUID:              data["business_guid"].(string),
		CatalogGUID:               data["catalog_guid"].(string),
		CancellationHours:         data["cancellation_hours"].(int),
		CancellationAmount:        data["cancellation_amount"].(int),
		NoShowFee:                 data["no_show_fee"].(int),
		BookingDepositePercentage: data["booking_deposite_percentage"].(int),
		BookingTerms:              data["booking_terms"].(string),
		BookingPolices:            data["booking_polices"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessPoliciesHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      data["business_policies_guid"].(string),
		BusinessGUID:              data["business_guid"].(string),
		CatalogGUID:               data["catalog_guid"].(string),
		CancellationHours:         data["cancellation_hours"].(int),
		CancellationAmount:        data["cancellation_amount"].(int),
		NoShowFee:                 data["no_show_fee"].(int),
		BookingDepositePercentage: data["booking_deposite_percentage"].(int),
		BookingTerms:              data["booking_terms"].(string),
		BookingPolices:            data["booking_polices"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *BusinessPoliciesHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID: data["business_policies_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessPoliciesHandlers) GETBYID(ctx echo.Context) error {
	business_policies_guid := ctx.Param("business_policies_guid")
	mydata, err := p.apiBusiness.GETBYID(business_policies_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessPoliciesHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessPolicies")
}
