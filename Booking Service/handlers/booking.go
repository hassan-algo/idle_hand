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

// validateBooking validates the required fields of a booking
func validateBooking(booking structs.Booking) error {
	if booking.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if booking.ClientGUID == "" {
		return errors.New("client_guid is required")
	}
	if booking.ServiceGUID == "" {
		return errors.New("service_guid is required")
	}
	if booking.BookingDate == "" {
		return errors.New("booking_date is required")
	}
	if booking.BookingTime == "" {
		return errors.New("booking_time is required")
	}
	if booking.BookingStatus == "" {
		return errors.New("booking_status is required")
	}
	if booking.CustomerRelation == "" {
		return errors.New("customer_relation is required")
	}
	if booking.AssignStaffGUID == "" {
		return errors.New("assign_staff_guid is required")
	}
	if booking.PaymentMethod == "" {
		return errors.New("payment_method is required")
	}
	return nil
}

type BookingHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBookingHandler() *BookingHandlers {
	return &BookingHandlers{}
}

func (h *BookingHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get all bookings
// @Description Get all bookings
// @Produce json
// @Success 200 {object} structs.Booking "bookings"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking [get]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch bookings", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Bookings fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Create a new booking
// @Description Create a new booking with the provided details
// @Accept json
// @Produce json
// @Param booking body structs.Booking true "Booking Object"
// @Success 200 {object} structs.Booking "created booking"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking [post]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	clientGUID, ok := data["client_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid client_guid format", nil)
	}

	serviceGUID, ok := data["service_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid service_guid format", nil)
	}

	bookingDate, ok := data["booking_date"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_date format", nil)
	}

	bookingTime, ok := data["booking_time"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_time format", nil)
	}

	bookingStatus, ok := data["booking_status"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_status format", nil)
	}

	customerRelation, ok := data["customer_relation"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid customer_relation format", nil)
	}

	assignStaffGUID, ok := data["assign_staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid assign_staff_guid format", nil)
	}

	paymentMethod, ok := data["payment_method"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid payment_method format", nil)
	}

	booking_guid := uuid.New().String()
	my_struct := structs.Booking{
		BookingGUID:      booking_guid,
		BusinessGUID:     businessGUID,
		ClientGUID:       clientGUID,
		ServiceGUID:      serviceGUID,
		BookingDate:      bookingDate,
		BookingTime:      bookingTime,
		BookingStatus:    bookingStatus,
		CustomerRelation: customerRelation,
		AssignStaffGUID:  assignStaffGUID,
		PaymentMethod:    paymentMethod,
	}

	// Validate the booking
	if err := validateBooking(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create booking", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Booking created successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Update a booking
// @Description Update an existing booking with the provided details
// @Accept json
// @Produce json
// @Param booking body structs.Booking true "Booking Object"
// @Success 200 {object} structs.Booking "updated booking"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking [put]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if booking_guid exists
	bookingGUID, ok := data["booking_guid"].(string)
	if !ok || bookingGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "booking_guid is required", nil)
	}

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	clientGUID, ok := data["client_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid client_guid format", nil)
	}

	serviceGUID, ok := data["service_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid service_guid format", nil)
	}

	bookingDate, ok := data["booking_date"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_date format", nil)
	}

	bookingTime, ok := data["booking_time"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_time format", nil)
	}

	bookingStatus, ok := data["booking_status"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid booking_status format", nil)
	}

	customerRelation, ok := data["customer_relation"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid customer_relation format", nil)
	}

	assignStaffGUID, ok := data["assign_staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid assign_staff_guid format", nil)
	}

	paymentMethod, ok := data["payment_method"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid payment_method format", nil)
	}

	my_struct := structs.Booking{
		BookingGUID:      bookingGUID,
		BusinessGUID:     businessGUID,
		ClientGUID:       clientGUID,
		ServiceGUID:      serviceGUID,
		BookingDate:      bookingDate,
		BookingTime:      bookingTime,
		BookingStatus:    bookingStatus,
		CustomerRelation: customerRelation,
		AssignStaffGUID:  assignStaffGUID,
		PaymentMethod:    paymentMethod,
	}

	// Validate the booking
	if err := validateBooking(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update booking", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Booking updated successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Delete a booking
// @Description Delete a booking by its GUID
// @Accept json
// @Produce json
// @Param booking body structs.Booking true "Booking Object with GUID"
// @Success 200 {object} structs.Booking "deleted booking"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking [delete]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	bookingGUID, ok := data["booking_guid"].(string)
	if !ok || bookingGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "booking_guid is required", nil)
	}

	my_struct := structs.Booking{
		BookingGUID: bookingGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete booking", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Booking deleted successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Get booking by ID
// @Description Get a specific booking by its GUID
// @Produce json
// @Param booking_guid path string true "Booking GUID"
// @Success 200 {object} structs.Booking "booking"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking/{booking_guid} [get]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) GETBYID(ctx echo.Context) error {
	booking_guid := ctx.Param("booking_guid")
	if booking_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "booking_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(booking_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch booking", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Booking fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Multiple bookings creation
// @Description Create multiple bookings at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST Booking"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /booking/multi [post]
// @Security BearerAuth
// @Tags booking
func (p *BookingHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple bookings creation is not implemented yet", nil)
}
