package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

// @Summary Get catalog_availability
// @Description
// @Produce json
// @Success 200 {object} structs.CatalogAvailability "catalog_availability"
// @Router /catalog_availability [get]
// @Security ApiKeyAuth
// @Tags catalog_availability
func (p *CatalogAvailabilityHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogAvailabilityHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	catalog_availability_guid := uuid.New().String()
	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: catalog_availability_guid,
		CatalogGUID:             data["catalog_guid"].(string),
		BusinessGUID:            data["business_guid"].(string),
		DayOfWeek:               data["day_of_week"].(string),
		HoursPerDay:             data["hours_per_day"].(int),
		SlotsPerDay:             data["slots_per_day"].(string),
		NumberOfBreaks:          data["number_of_breaks"].(int),
		TimePerBreak:            data["time_per_break"].(string),
		BufferPerAppointment:    data["buffer_per_appointment"].(int),
		AcceptSameDayBooking:    data["accept_same_day_booking"].(int),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogAvailabilityHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: data["catalog_availability_guid"].(string),
		CatalogGUID:             data["catalog_guid"].(string),
		BusinessGUID:            data["business_guid"].(string),
		DayOfWeek:               data["day_of_week"].(string),
		HoursPerDay:             data["hours_per_day"].(int),
		SlotsPerDay:             data["slots_per_day"].(string),
		NumberOfBreaks:          data["number_of_breaks"].(int),
		TimePerBreak:            data["time_per_break"].(string),
		BufferPerAppointment:    data["buffer_per_appointment"].(int),
		AcceptSameDayBooking:    data["accept_same_day_booking"].(int),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *CatalogAvailabilityHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.CatalogAvailability{
		CatalogAvailabilityGUID: data["catalog_availability_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogAvailabilityHandlers) GETBYID(ctx echo.Context) error {
	catalog_availability_guid := ctx.Param("catalog_availability_guid")
	mydata, err := p.apiBusiness.GETBYID(catalog_availability_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *CatalogAvailabilityHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST CatalogAvailability")
}
