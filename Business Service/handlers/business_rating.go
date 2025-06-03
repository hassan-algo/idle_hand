package handlers

import (
	"net/http"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BusinessRatingHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewBusinessRatingHandler() *BusinessRatingHandlers {
	return &BusinessRatingHandlers{}
}

func (h *BusinessRatingHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get business_rating
// @Description
// @Produce json
// @Success 200 {object} structs.BusinessRating "business_rating"
// @Router /business_rating [get]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessRatingHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	business_rating_guid := uuid.New().String()
	my_struct := structs.BusinessRating{
		BusinessRatingGUID: business_rating_guid,
		BusinessGUID:       data["business_guid"].(string),
		UserGUID:           data["user_guid"].(string),
		ReviewText:         data["review_text"].(string),
		Rating:             data["rating"].(string),
	}
	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessRatingHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessRating{
		BusinessRatingGUID: data["business_rating_guid"].(string),
		BusinessGUID:       data["business_guid"].(string),
		UserGUID:           data["user_guid"].(string),
		ReviewText:         data["review_text"].(string),
		Rating:             data["rating"].(string),
	}
	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessRatingHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)
	my_struct := structs.BusinessRating{
		BusinessRatingGUID: data["business_rating_guid"].(string),
	}
	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}
func (p *BusinessRatingHandlers) GETBYID(ctx echo.Context) error {
	business_rating_guid := ctx.Param("business_rating_guid")
	mydata, err := p.apiBusiness.GETBYID(business_rating_guid)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, mydata)
}

func (p *BusinessRatingHandlers) MULTIPOST(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "MULTIPOST BusinessRating")
}
