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

// validateCatalogAvailability validates the required fields of a catalog availability
func validateCatalogAvailability(availability structs.CatalogAvailability) error {
	if availability.CatalogGUID == "" {
		return errors.New("catalog_guid is required")
	}
	if availability.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if availability.DayOfWeek == "" {
		return errors.New("day_of_week is required")
	}
	if availability.HoursPerDay < 0 {
		return errors.New("hours_per_day cannot be negative")
	}
	if availability.NumberOfBreaks < 0 {
		return errors.New("number_of_breaks cannot be negative")
	}
	if availability.BufferPerAppointment < 0 {
		return errors.New("buffer_per_appointment cannot be negative")
	}
	if availability.AcceptSameDayBooking < 0 || availability.AcceptSameDayBooking > 1 {
		return errors.New("accept_same_day_booking must be 0 or 1")
	}
	return nil
}

type CatalogAvailabilityHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewCatalogAvailabilityHandler() *CatalogAvailabilityHandlers {
	return &CatalogAvailabilityHandlers{}
}

func (h *CatalogAvailabilityHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get all catalog availabilities
// @Description Get all catalog availabilities
// @Produce json
// @Success 200 {object} structs.CatalogAvailability "catalog_availability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability [get]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch catalog availabilities", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new catalog availability
// @Description Create a new catalog availability with the provided details
// @Accept json
// @Produce json
// @Param catalog_availability body structs.CatalogAvailability true "Catalog Availability Object"
// @Success 200 {object} structs.CatalogAvailability "created catalog availability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability [post]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	dayOfWeek, ok := data["day_of_week"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid day_of_week format", nil)
	}

	hoursPerDay, ok := data["hours_per_day"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid hours_per_day format", nil)
	}

	slotsPerDay, ok := data["slots_per_day"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid slots_per_day format", nil)
	}

	numberOfBreaks, ok := data["number_of_breaks"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid number_of_breaks format", nil)
	}

	timePerBreak, ok := data["time_per_break"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid time_per_break format", nil)
	}

	bufferPerAppointment, ok := data["buffer_per_appointment"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid buffer_per_appointment format", nil)
	}

	acceptSameDayBooking, ok := data["accept_same_day_booking"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid accept_same_day_booking format", nil)
	}

	catalog_availability_guid := uuid.New().String()
	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: catalog_availability_guid,
		CatalogGUID:             catalogGUID,
		BusinessGUID:            businessGUID,
		DayOfWeek:               dayOfWeek,
		HoursPerDay:             hoursPerDay,
		SlotsPerDay:             slotsPerDay,
		NumberOfBreaks:          numberOfBreaks,
		TimePerBreak:            timePerBreak,
		BufferPerAppointment:    bufferPerAppointment,
		AcceptSameDayBooking:    acceptSameDayBooking,
	}

	// Validate the catalog availability
	if err := validateCatalogAvailability(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create catalog availability", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update a catalog availability
// @Description Update an existing catalog availability with the provided details
// @Accept json
// @Produce json
// @Param catalog_availability body structs.CatalogAvailability true "Catalog Availability Object"
// @Success 200 {object} structs.CatalogAvailability "updated catalog availability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability [put]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if catalog_availability_guid exists
	catalogAvailabilityGUID, ok := data["catalog_availability_guid"].(string)
	if !ok || catalogAvailabilityGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_availability_guid is required", nil)
	}

	// Type assertion with error handling
	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	dayOfWeek, ok := data["day_of_week"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid day_of_week format", nil)
	}

	hoursPerDay, ok := data["hours_per_day"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid hours_per_day format", nil)
	}

	slotsPerDay, ok := data["slots_per_day"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid slots_per_day format", nil)
	}

	numberOfBreaks, ok := data["number_of_breaks"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid number_of_breaks format", nil)
	}

	timePerBreak, ok := data["time_per_break"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid time_per_break format", nil)
	}

	bufferPerAppointment, ok := data["buffer_per_appointment"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid buffer_per_appointment format", nil)
	}

	acceptSameDayBooking, ok := data["accept_same_day_booking"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid accept_same_day_booking format", nil)
	}

	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: catalogAvailabilityGUID,
		CatalogGUID:             catalogGUID,
		BusinessGUID:            businessGUID,
		DayOfWeek:               dayOfWeek,
		HoursPerDay:             hoursPerDay,
		SlotsPerDay:             slotsPerDay,
		NumberOfBreaks:          numberOfBreaks,
		TimePerBreak:            timePerBreak,
		BufferPerAppointment:    bufferPerAppointment,
		AcceptSameDayBooking:    acceptSameDayBooking,
	}

	// Validate the catalog availability
	if err := validateCatalogAvailability(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update catalog availability", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete a catalog availability
// @Description Delete a catalog availability by its GUID
// @Accept json
// @Produce json
// @Param catalog_availability body structs.CatalogAvailability true "Catalog Availability Object with GUID"
// @Success 200 {object} structs.CatalogAvailability "deleted catalog availability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability [delete]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	catalogAvailabilityGUID, ok := data["catalog_availability_guid"].(string)
	if !ok || catalogAvailabilityGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_availability_guid is required", nil)
	}

	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: catalogAvailabilityGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete catalog availability", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get catalog availability by ID
// @Description Get a specific catalog availability by its GUID
// @Produce json
// @Param catalog_availability_guid path string true "Catalog Availability GUID"
// @Success 200 {object} structs.CatalogAvailability "catalog availability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability/{catalog_availability_guid} [get]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) GETBYID(ctx echo.Context) error {
	catalog_availability_guid := ctx.Param("catalog_availability_guid")
	if catalog_availability_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_availability_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(catalog_availability_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch catalog availability", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple catalog availabilities creation
// @Description Create multiple catalog availabilities at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST CatalogAvailability"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog_availability/multi [post]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple catalog availabilities creation is not implemented yet", nil)
}
