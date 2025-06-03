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

// @Summary Get all business policies
// @Description Get all business policies
// @Produce json
// @Success 200 {object} structs.BusinessPolicies "business_policies"
// @Failure 400 {string} string "error message"
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

// @Summary Create a new business policy
// @Description Create a new business policy with the provided details
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object"
// @Success 200 {object} structs.BusinessPolicies "created business policy"
// @Failure 400 {string} string "error message"
// @Router /business_policies [post]
// @Security ApiKeyAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	business_policies_guid := uuid.New().String()
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      business_policies_guid,
		BusinessGUID:              data["business_guid"].(string),
		CatalogGUID:               data["catalog_guid"].(string),
		CancellationHours:         data["cancellation_hours"].(float64),
		CancellationAmount:        data["cancellation_amount"].(float64),
		NoShowFee:                 data["no_show_fee"].(float64),
		BookingDepositePercentage: data["booking_deposite_percentage"].(float64),
		BookingTerms:              data["booking_terms"].(string),
		BookingPolices:            data["booking_polices"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update a business policy
// @Description Update an existing business policy with the provided details
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object"
// @Success 200 {object} structs.BusinessPolicies "updated business policy"
// @Failure 400 {string} string "error message"
// @Router /business_policies [put]
// @Security ApiKeyAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      data["business_policies_guid"].(string),
		BusinessGUID:              data["business_guid"].(string),
		CatalogGUID:               data["catalog_guid"].(string),
		CancellationHours:         data["cancellation_hours"].(float64),
		CancellationAmount:        data["cancellation_amount"].(float64),
		NoShowFee:                 data["no_show_fee"].(float64),
		BookingDepositePercentage: data["booking_deposite_percentage"].(float64),
		BookingTerms:              data["booking_terms"].(string),
		BookingPolices:            data["booking_polices"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete a business policy
// @Description Delete a business policy by its GUID
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object with GUID"
// @Success 200 {object} structs.BusinessPolicies "deleted business policy"
// @Failure 400 {string} string "error message"
// @Router /business_policies [delete]
// @Security ApiKeyAuth
// @Tags business_policies
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

// @Summary Get business policy by ID
// @Description Get a specific business policy by its GUID
// @Produce json
// @Param business_policies_guid path string true "Business Policy GUID"
// @Success 200 {object} structs.BusinessPolicies "business policy"
// @Failure 400 {string} string "error message"
// @Router /business_policies/{business_policies_guid} [get]
// @Security ApiKeyAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) GETBYID(ctx echo.Context) error {
	business_policies_guid := ctx.Param("business_policies_guid")
	mydata, err := p.apiBusiness.GETBYID(business_policies_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple business policies creation
// @Description Create multiple business policies at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST BusinessPolicies"
// @Router /business_policies/multi [post]
// @Security ApiKeyAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessPolicies")
}
