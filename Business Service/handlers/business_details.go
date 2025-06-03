package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BusinessDetailsHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBusinessDetailsHandler() *BusinessDetailsHandlers {
	return &BusinessDetailsHandlers{}
}

func (h *BusinessDetailsHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get business_details
// @Description
// @Produce json
// @Success 200 {object} structs.BusinessDetails "business_details"
// @Router /business_details [get]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) GET(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "GET BusinessDetails")
}

func (p *BusinessDetailsHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	businessGUID := uuid.New().String()
	my_struct := structs.BusinessDetails{
		BusinessName:          data["business_name"].(string),
		IndustryType:          data["industry_type"].(string),
		BusinessDescription:   data["business_description"].(string),
		BusinessPhoneNumber:   data["business_phone_number"].(string),
		BusinessLogo:          data["business_logo"].(string),
		BusinessPhoto:         data["business_photo"].(string),
		BusinessEmail:         data["business_email"].(string),
		BusinessAddress1:      data["business_address1"].(string),
		BusinessAddress2:      data["business_address2"].(string),
		BusinessCity:          data["business_city"].(string),
		BusinessZipCode:       data["business_zip_code"].(string),
		BusinessSocialMedia:   data["business_social_media"].(string),
		BusinessPaymentMethod: data["business_payment_method"].(string),
		BusinessLicense:       data["business_license"].(string),
		BusinessLocation:      data["business_location"].(string),
		BusinessGUID:          businessGUID,
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessDetailsHandlers) PUT(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "PUT BusinessDetails")
}
func (p *BusinessDetailsHandlers) DELETE(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "DELETE BusinessDetails")
}

func (p *BusinessDetailsHandlers) GETBYID(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "GETBYID BusinessDetails")
}

func (p *BusinessDetailsHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessDetails")
}
