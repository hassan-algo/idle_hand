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

// validateBusinessClient validates the required fields of a business client
func validateBusinessClient(client structs.BusinessClient) error {
	if client.ClientGUID == "" {
		return errors.New("client_guid is required")
	}
	if client.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	return nil
}

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

// @Summary Get all business clients
// @Description Get all business clients
// @Produce json
// @Success 200 {object} structs.BusinessClient "business_client"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_client [get]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business clients", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new business client
// @Description Create a new business client with the provided information
// @Accept json
// @Produce json
// @Param business_client body structs.BusinessClient true "Business Client Object"
// @Success 200 {object} structs.BusinessClient "created business client"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_client [post]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	clientGUID, ok := data["client_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid client_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	preferredContactMethod, _ := data["preferred_contact_method"].(string)
	source, _ := data["source"].(string)
	note, _ := data["note"].(string)

	business_client_guid := uuid.New().String()
	my_struct := structs.BusinessClient{
		BusinessClientGUID:     business_client_guid,
		ClientGUID:             clientGUID,
		BusinessGUID:           businessGUID,
		PreferredContactMethod: preferredContactMethod,
		Source:                 source,
		Note:                   note,
	}

	// Validate the business client
	if err := validateBusinessClient(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create business client", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update business client
// @Description Update an existing business client
// @Accept json
// @Produce json
// @Param business_client body structs.BusinessClient true "Business Client Object"
// @Success 200 {object} structs.BusinessClient "updated business client"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_client [put]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessClientGUID, ok := data["business_client_guid"].(string)
	if !ok || businessClientGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_client_guid is required", nil)
	}

	clientGUID, ok := data["client_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid client_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	preferredContactMethod, _ := data["preferred_contact_method"].(string)
	source, _ := data["source"].(string)
	note, _ := data["note"].(string)

	my_struct := structs.BusinessClient{
		BusinessClientGUID:     businessClientGUID,
		ClientGUID:             clientGUID,
		BusinessGUID:           businessGUID,
		PreferredContactMethod: preferredContactMethod,
		Source:                 source,
		Note:                   note,
	}

	// Validate the business client
	if err := validateBusinessClient(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update business client", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete business client
// @Description Delete a business client by its GUID
// @Accept json
// @Produce json
// @Param business_client body structs.BusinessClient true "Business Client Object with GUID"
// @Success 200 {object} structs.BusinessClient "deleted business client"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_client [delete]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessClientGUID, ok := data["business_client_guid"].(string)
	if !ok || businessClientGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_client_guid is required", nil)
	}

	my_struct := structs.BusinessClient{
		BusinessClientGUID: businessClientGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete business client", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get business client by ID
// @Description Get a specific business client by its GUID
// @Produce json
// @Param business_client_guid path string true "Business Client GUID"
// @Success 200 {object} structs.BusinessClient "business client"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_client/{business_client_guid} [get]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) GETBYID(ctx echo.Context) error {
	business_client_guid := ctx.Param("business_client_guid")
	if business_client_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "business_client_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(business_client_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business client", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple business clients creation
// @Description Create multiple business clients at once
// @Accept json
// @Produce json
// @Success 501 {object} ErrorResponse "Not Implemented"
// @Router /business_client/multi [post]
// @Security ApiKeyAuth
// @Tags business_client
func (p *BusinessClientHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple business clients creation is not implemented yet", nil)
}
