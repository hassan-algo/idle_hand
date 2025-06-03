package handlers

import (
	"errors"
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// handleError creates a standardized error response
func handleError(ctx echo.Context, status int, message string, err error) error {
	response := ErrorResponse{
		Status:  status,
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
	}
	return ctx.JSON(status, response)
}

// validateBusinessDetails validates the required fields of a business details
func validateBusinessDetails(business structs.BusinessDetails) error {
	if business.BusinessName == "" {
		return errors.New("business_name is required")
	}
	if business.IndustryType == "" {
		return errors.New("industry_type is required")
	}
	if business.BusinessPhoneNumber == "" {
		return errors.New("business_phone_number is required")
	}
	if business.BusinessEmail == "" {
		return errors.New("business_email is required")
	}
	if business.BusinessAddress1 == "" {
		return errors.New("business_address1 is required")
	}
	if business.BusinessCity == "" {
		return errors.New("business_city is required")
	}
	if business.BusinessZipCode == "" {
		return errors.New("business_zip_code is required")
	}
	return nil
}

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

// @Summary Get all business details
// @Description Get all business details
// @Produce json
// @Success 200 {object} structs.BusinessDetails "business_details"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_details [get]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business details", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new business details
// @Description Create a new business details with the provided information
// @Accept json
// @Produce json
// @Param business_details body structs.BusinessDetails true "Business Details Object"
// @Success 200 {object} structs.BusinessDetails "created business details"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_details [post]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessName, ok := data["business_name"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_name format", nil)
	}

	industryType, ok := data["industry_type"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid industry_type format", nil)
	}

	businessDescription, ok := data["business_description"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_description format", nil)
	}

	businessPhoneNumber, ok := data["business_phone_number"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_phone_number format", nil)
	}

	businessLogo, ok := data["business_logo"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_logo format", nil)
	}

	businessPhoto, ok := data["business_photo"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_photo format", nil)
	}

	businessEmail, ok := data["business_email"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_email format", nil)
	}

	businessAddress1, ok := data["business_address1"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_address1 format", nil)
	}

	businessAddress2, ok := data["business_address2"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_address2 format", nil)
	}

	businessCity, ok := data["business_city"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_city format", nil)
	}

	businessZipCode, ok := data["business_zip_code"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_zip_code format", nil)
	}

	businessSocialMedia, ok := data["business_social_media"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_social_media format", nil)
	}

	businessPaymentMethod, ok := data["business_payment_method"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_payment_method format", nil)
	}

	businessLicense, ok := data["business_license"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_license format", nil)
	}

	businessLocation, ok := data["business_location"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_location format", nil)
	}

	businessGUID := uuid.New().String()
	my_struct := structs.BusinessDetails{
		BusinessName:          businessName,
		IndustryType:          industryType,
		BusinessDescription:   businessDescription,
		BusinessPhoneNumber:   businessPhoneNumber,
		BusinessLogo:          businessLogo,
		BusinessPhoto:         businessPhoto,
		BusinessEmail:         businessEmail,
		BusinessAddress1:      businessAddress1,
		BusinessAddress2:      businessAddress2,
		BusinessCity:          businessCity,
		BusinessZipCode:       businessZipCode,
		BusinessSocialMedia:   businessSocialMedia,
		BusinessPaymentMethod: businessPaymentMethod,
		BusinessLicense:       businessLicense,
		BusinessLocation:      businessLocation,
		BusinessGUID:          businessGUID,
	}

	// Validate the business details
	if err := validateBusinessDetails(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create business details", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update business details
// @Description Update an existing business details
// @Accept json
// @Produce json
// @Param business_details body structs.BusinessDetails true "Business Details Object"
// @Success 200 {object} structs.BusinessDetails "updated business details"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_details [put]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if business_guid exists
	businessGUID, ok := data["business_guid"].(string)
	if !ok || businessGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_guid is required", nil)
	}

	my_struct := structs.BusinessDetails{
		BusinessGUID: businessGUID,
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update business details", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete business details
// @Description Delete a business details by its GUID
// @Accept json
// @Produce json
// @Param business_details body structs.BusinessDetails true "Business Details Object with GUID"
// @Success 200 {object} structs.BusinessDetails "deleted business details"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_details [delete]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessGUID, ok := data["business_guid"].(string)
	if !ok || businessGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_guid is required", nil)
	}

	my_struct := structs.BusinessDetails{
		BusinessGUID: businessGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete business details", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get business details by ID
// @Description Get a specific business details by its GUID
// @Produce json
// @Param business_guid path string true "Business GUID"
// @Success 200 {object} structs.BusinessDetails "business details"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_details/{business_guid} [get]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) GETBYID(ctx echo.Context) error {
	businessGUID := ctx.Param("business_guid")
	if businessGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(businessGUID)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business details", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple business details creation
// @Description Create multiple business details at once
// @Accept json
// @Produce json
// @Success 501 {object} ErrorResponse "Not Implemented"
// @Router /business_details/multi [post]
// @Security ApiKeyAuth
// @Tags business_details
func (p *BusinessDetailsHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple business details creation is not implemented yet", nil)
}
