package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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

// @Summary Get business_staff
// @Description
// @Produce json
// @Success 200 {object} structs.BusinessStaff "business_staff"
// @Router /business_staff [get]
// @Security ApiKeyAuth
// @Tags business_staff
func (p *BusinessStaffHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessStaffHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	business_staff_guid := uuid.New().String()
	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: business_staff_guid,
		StaffGUID:         data["staff_guid"].(string),
		BusinessGUID:      data["business_guid"].(string),
		Profession:        data["profession"].(string),
		ArrivalTime:       data["arrival_time"].(string),
		LeaveTime:         data["leave_time"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessStaffHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: data["business_staff_guid"].(string),
		StaffGUID:         data["staff_guid"].(string),
		BusinessGUID:      data["business_guid"].(string),
		Profession:        data["profession"].(string),
		ArrivalTime:       data["arrival_time"].(string),
		LeaveTime:         data["leave_time"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *BusinessStaffHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessStaff{
		BusinessStaffGUID: data["business_staff_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessStaffHandlers) GETBYID(ctx echo.Context) error {
	business_staff_guid := ctx.Param("business_staff_guid")
	mydata, err := p.apiBusiness.GETBYID(business_staff_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessStaffHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessStaff")
}
