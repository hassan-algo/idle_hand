package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RequestReviewHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewRequestReviewHandler() *RequestReviewHandlers {
	return &RequestReviewHandlers{}
}

func (h *RequestReviewHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get request_review
// @Description
// @Produce json
// @Success 200 {object} structs.RequestReview "request_review"
// @Router /request_review [get]
// @Security ApiKeyAuth
// @Tags request_review
func (p *RequestReviewHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *RequestReviewHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	request_review_guid := uuid.New().String()
	my_struct := structs.RequestReview{
		RequestReviewsGUID: request_review_guid,
		BusinessGUID:       data["business_guid"].(string),
		UserGUID:           data["user_guid"].(string),
		ReviewRequested:    1,
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, mydata)
}

func (p *RequestReviewHandlers) PUT(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "PUT RequestReview")
}
func (p *RequestReviewHandlers) DELETE(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "DELETE RequestReview")
}

func (p *RequestReviewHandlers) GETBYID(ctx echo.Context) error {
	request_review_guid := ctx.Param("request_review_guid")
	mydata, err := p.apiBusiness.GETBYID(request_review_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *RequestReviewHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST RequestReview")
}
