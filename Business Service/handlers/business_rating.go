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

// validateBusinessRating validates the required fields of a business rating
func validateBusinessRating(rating structs.BusinessRating) error {
	if rating.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if rating.UserGUID == "" {
		return errors.New("user_guid is required")
	}
	if rating.ReviewText == "" {
		return errors.New("review_text is required")
	}
	if rating.Rating == "" {
		return errors.New("rating is required")
	}
	return nil
}

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

// @Summary Get all business ratings
// @Description Get all business ratings
// @Produce json
// @Success 200 {object} structs.BusinessRating "business_rating"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_rating [get]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business ratings", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new business rating
// @Description Create a new business rating with the provided information
// @Accept json
// @Produce json
// @Param business_rating body structs.BusinessRating true "Business Rating Object"
// @Success 200 {object} structs.BusinessRating "created business rating"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_rating [post]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	userGUID, ok := data["user_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid user_guid format", nil)
	}

	reviewText, ok := data["review_text"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid review_text format", nil)
	}

	rating, ok := data["rating"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid rating format", nil)
	}

	business_rating_guid := uuid.New().String()
	my_struct := structs.BusinessRating{
		BusinessRatingGUID: business_rating_guid,
		BusinessGUID:       businessGUID,
		UserGUID:           userGUID,
		ReviewText:         reviewText,
		Rating:             rating,
	}

	// Validate the business rating
	if err := validateBusinessRating(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create business rating", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update business rating
// @Description Update an existing business rating
// @Accept json
// @Produce json
// @Param business_rating body structs.BusinessRating true "Business Rating Object"
// @Success 200 {object} structs.BusinessRating "updated business rating"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_rating [put]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessRatingGUID, ok := data["business_rating_guid"].(string)
	if !ok || businessRatingGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_rating_guid is required", nil)
	}

	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	userGUID, ok := data["user_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid user_guid format", nil)
	}

	reviewText, ok := data["review_text"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid review_text format", nil)
	}

	rating, ok := data["rating"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid rating format", nil)
	}

	my_struct := structs.BusinessRating{
		BusinessRatingGUID: businessRatingGUID,
		BusinessGUID:       businessGUID,
		UserGUID:           userGUID,
		ReviewText:         reviewText,
		Rating:             rating,
	}

	// Validate the business rating
	if err := validateBusinessRating(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update business rating", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete business rating
// @Description Delete a business rating by its GUID
// @Accept json
// @Produce json
// @Param business_rating body structs.BusinessRating true "Business Rating Object with GUID"
// @Success 200 {object} structs.BusinessRating "deleted business rating"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_rating [delete]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	businessRatingGUID, ok := data["business_rating_guid"].(string)
	if !ok || businessRatingGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "business_rating_guid is required", nil)
	}

	my_struct := structs.BusinessRating{
		BusinessRatingGUID: businessRatingGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete business rating", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get business rating by ID
// @Description Get a specific business rating by its GUID
// @Produce json
// @Param business_rating_guid path string true "Business Rating GUID"
// @Success 200 {object} structs.BusinessRating "business rating"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /business_rating/{business_rating_guid} [get]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) GETBYID(ctx echo.Context) error {
	business_rating_guid := ctx.Param("business_rating_guid")
	if business_rating_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "business_rating_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(business_rating_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch business rating", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple business ratings creation
// @Description Create multiple business ratings at once
// @Accept json
// @Produce json
// @Success 501 {object} ErrorResponse "Not Implemented"
// @Router /business_rating/multi [post]
// @Security ApiKeyAuth
// @Tags business_rating
func (p *BusinessRatingHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple business ratings creation is not implemented yet", nil)
}
