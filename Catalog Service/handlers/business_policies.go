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

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// ValidationError represents validation error details
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// validateBusinessPolicy validates the required fields of a business policy
func validateBusinessPolicy(policy structs.BusinessPolicies) error {
	if policy.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if policy.CatalogGUID == "" {
		return errors.New("catalog_guid is required")
	}
	if policy.CancellationHours < 0 {
		return errors.New("cancellation_hours cannot be negative")
	}
	if policy.CancellationAmount < 0 {
		return errors.New("cancellation_amount cannot be negative")
	}
	if policy.NoShowFee < 0 {
		return errors.New("no_show_fee cannot be negative")
	}
	if policy.BookingDepositePercentage < 0 || policy.BookingDepositePercentage > 100 {
		return errors.New("booking_deposite_percentage must be between 0 and 100")
	}
	return nil
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
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies [get]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) GET(ctx echo.Context) error {

	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business policies", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Business policies fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Create a new business policy
// @Description Create a new business policy with the provided details
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object"
// @Success 200 {object} structs.BusinessPolicies "created business policy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies [post]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_guid format", nil)
	}

	cancellationHours, ok := data["cancellation_hours"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid cancellation_hours format", nil)
	}

	cancellationAmount, ok := data["cancellation_amount"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid cancellation_amount format", nil)
	}

	noShowFee, ok := data["no_show_fee"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid no_show_fee format", nil)
	}

	bookingDepositePercentage, ok := data["booking_deposite_percentage"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_deposite_percentage format", nil)
	}

	bookingTerms, ok := data["booking_terms"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_terms format", nil)
	}

	bookingPolices, ok := data["booking_polices"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_polices format", nil)
	}

	business_policies_guid := uuid.New().String()
	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      business_policies_guid,
		BusinessGUID:              businessGUID,
		CatalogGUID:               catalogGUID,
		CancellationHours:         cancellationHours,
		CancellationAmount:        cancellationAmount,
		NoShowFee:                 noShowFee,
		BookingDepositePercentage: bookingDepositePercentage,
		BookingTerms:              bookingTerms,
		BookingPolices:            bookingPolices,
	}

	// Validate the business policy
	if err := validateBusinessPolicy(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create business policy", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Business policy created successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Update a business policy
// @Description Update an existing business policy with the provided details
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object"
// @Success 200 {object} structs.BusinessPolicies "updated business policy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies [put]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if business_policies_guid exists
	businessPoliciesGUID, ok := data["business_policies_guid"].(string)
	if !ok || businessPoliciesGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_policies_guid is required", nil)
	}

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_guid format", nil)
	}

	cancellationHours, ok := data["cancellation_hours"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid cancellation_hours format", nil)
	}

	cancellationAmount, ok := data["cancellation_amount"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid cancellation_amount format", nil)
	}

	noShowFee, ok := data["no_show_fee"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid no_show_fee format", nil)
	}

	bookingDepositePercentage, ok := data["booking_deposite_percentage"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_deposite_percentage format", nil)
	}

	bookingTerms, ok := data["booking_terms"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_terms format", nil)
	}

	bookingPolices, ok := data["booking_polices"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_polices format", nil)
	}

	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID:      businessPoliciesGUID,
		BusinessGUID:              businessGUID,
		CatalogGUID:               catalogGUID,
		CancellationHours:         cancellationHours,
		CancellationAmount:        cancellationAmount,
		NoShowFee:                 noShowFee,
		BookingDepositePercentage: bookingDepositePercentage,
		BookingTerms:              bookingTerms,
		BookingPolices:            bookingPolices,
	}

	// Validate the business policy
	if err := validateBusinessPolicy(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update business policy", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Business policy updated successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Delete a business policy
// @Description Delete a business policy by its GUID
// @Accept json
// @Produce json
// @Param business_policy body structs.BusinessPolicies true "Business Policy Object with GUID"
// @Success 200 {object} structs.BusinessPolicies "deleted business policy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies [delete]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessPoliciesGUID, ok := data["business_policies_guid"].(string)
	if !ok || businessPoliciesGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_policies_guid is required", nil)
	}

	my_struct := structs.BusinessPolicies{
		BusinessPoliciesGUID: businessPoliciesGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete business policy", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Business policy deleted successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Get business policy by ID
// @Description Get a specific business policy by its GUID
// @Produce json
// @Param business_policies_guid path string true "Business Policy GUID"
// @Success 200 {object} structs.BusinessPolicies "business policy"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies/{business_policies_guid} [get]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) GETBYID(ctx echo.Context) error {
	business_policies_guid := ctx.Param("business_policies_guid")
	if business_policies_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "business_policies_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(business_policies_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business policy", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Business policy fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Multiple business policies creation
// @Description Create multiple business policies at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST BusinessPolicies"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_policies/multi [post]
// @Security BearerAuth
// @Tags business_policies
func (p *BusinessPoliciesHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple business policies creation is not implemented yet", nil)
}
