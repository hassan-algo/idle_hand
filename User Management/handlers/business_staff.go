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

// validateBusinessStaff validates the required fields of a business staff
func validateBusinessStaff(staff structs.BusinessStaff) error {
	if staff.StaffGUID == "" {
		return errors.New("staff_guid is required")
	}
	if staff.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	return nil
}

type BusinessStaffHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBusinessStaffHandler() *BusinessStaffHandlers {
	return &BusinessStaffHandlers{}
}

func (h *BusinessStaffHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get all business staff
// @Description Get all business staff members
// @Produce json
// @Success 200 {object} structs.BusinessStaff "business staff"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_staff [get]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business staff", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new business staff
// @Description Create a new business staff member with the provided information
// @Accept json
// @Produce json
// @Param business_staff body structs.BusinessStaff true "Business Staff Object"
// @Success 200 {object} structs.BusinessStaff "created business staff"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_staff [post]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	staffGUID, ok := data["staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid staff_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	profession, _ := data["profession"].(string)
	arrivalTime, _ := data["arrival_time"].(string)
	leaveTime, _ := data["leave_time"].(string)

	business_staff_guid := uuid.New().String()
	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: business_staff_guid,
		StaffGUID:         staffGUID,
		BusinessGUID:      businessGUID,
		Profession:        profession,
		ArrivalTime:       arrivalTime,
		LeaveTime:         leaveTime,
	}

	// Validate the business staff
	if err := validateBusinessStaff(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create business staff", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update business staff
// @Description Update an existing business staff member
// @Accept json
// @Produce json
// @Param business_staff body structs.BusinessStaff true "Business Staff Object"
// @Success 200 {object} structs.BusinessStaff "updated business staff"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_staff [put]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessStaffGUID, ok := data["business_staff_guid"].(string)
	if !ok || businessStaffGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_staff_guid is required", nil)
	}

	staffGUID, ok := data["staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid staff_guid format", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	profession, _ := data["profession"].(string)
	arrivalTime, _ := data["arrival_time"].(string)
	leaveTime, _ := data["leave_time"].(string)

	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: businessStaffGUID,
		StaffGUID:         staffGUID,
		BusinessGUID:      businessGUID,
		Profession:        profession,
		ArrivalTime:       arrivalTime,
		LeaveTime:         leaveTime,
	}

	// Validate the business staff
	if err := validateBusinessStaff(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update business staff", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete business staff
// @Description Delete a business staff member by their GUID
// @Accept json
// @Produce json
// @Param business_staff body structs.BusinessStaff true "Business Staff Object with GUID"
// @Success 200 {object} structs.BusinessStaff "deleted business staff"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_staff [delete]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessStaffGUID, ok := data["business_staff_guid"].(string)
	if !ok || businessStaffGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_staff_guid is required", nil)
	}

	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: businessStaffGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete business staff", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get business staff by ID
// @Description Get a specific business staff member by their GUID
// @Produce json
// @Param business_staff_guid path string true "Business Staff GUID"
// @Success 200 {object} structs.BusinessStaff "business staff"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_staff/{business_staff_guid} [get]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) GETBYID(ctx echo.Context) error {
	business_staff_guid := ctx.Param("business_staff_guid")
	if business_staff_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "business_staff_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(business_staff_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business staff", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple business staff creation
// @Description Create multiple business staff members at once
// @Accept json
// @Produce json
// @Success 501 {object} ErrorResponse "Not Implemented"
// @Router /business_staff/multi [post]
// @Security BearerAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple business staff creation is not implemented yet", nil)
}
