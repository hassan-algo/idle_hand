package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

// @Summary Get booking
// @Description
// @Produce json
// @Success 200 {object} structs.Booking "booking"
// @Router /booking [get]
// @Security ApiKeyAuth
// @Tags booking
func (p *BookingHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BookingHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	booking_guid := uuid.New().String()
	my_struct := structs.Booking{
		BookingGUID:      booking_guid,
		BusinessGUID:     data["business_guid"].(string),
		ClientGUID:       data["client_guid"].(string),
		ServiceGUID:      data["service_guid"].(string),
		BookingDate:      data["booking_date"].(string),
		BookingTime:      data["booking_time"].(string),
		BookingStatus:    data["booking_status"].(string),
		CustomerRelation: data["customer_relation"].(string),
		AssignStaffGUID:  data["assign_staff_guid"].(string),
		PaymentMethod:    data["payment_method"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BookingHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	my_struct := structs.Booking{
		BookingGUID:      data["booking_guid"].(string),
		BusinessGUID:     data["business_guid"].(string),
		ClientGUID:       data["client_guid"].(string),
		ServiceGUID:      data["service_guid"].(string),
		BookingDate:      data["booking_date"].(string),
		BookingTime:      data["booking_time"].(string),
		BookingStatus:    data["booking_status"].(string),
		CustomerRelation: data["customer_relation"].(string),
		AssignStaffGUID:  data["assign_staff_guid"].(string),
		PaymentMethod:    data["payment_method"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *BookingHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.Booking{
		BookingGUID: data["booking_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BookingHandlers) GETBYID(ctx echo.Context) error {
	booking_guid := ctx.Param("booking_guid")
	mydata, err := p.apiBusiness.GETBYID(booking_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BookingHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST Booking")
}
