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

// validateRequestReview validates the required fields of a request review
func validateRequestReview(review structs.RequestReview) error {
	if review.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if review.UserGUID == "" {
		return errors.New("user_guid is required")
	}
	return nil
}

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

// @Summary Get all request reviews
// @Description Get all request reviews
// @Produce json
// @Success 200 {object} structs.RequestReview "request_review"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /request_review [get]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch request reviews", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new request review
// @Description Create a new request review with the provided information
// @Accept json
// @Produce json
// @Param request_review body structs.RequestReview true "Request Review Object"
// @Success 200 {object} structs.RequestReview "created request review"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /request_review [post]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) POST(ctx echo.Context) error {
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

	request_review_guid := uuid.New().String()
	my_struct := structs.RequestReview{
		RequestReviewsGUID: request_review_guid,
		BusinessGUID:       businessGUID,
		UserGUID:           userGUID,
		ReviewRequested:    1,
	}

	// Validate the request review
	if err := validateRequestReview(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create request review", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update request review
// @Description Update an existing request review
// @Accept json
// @Produce json
// @Param request_review body structs.RequestReview true "Request Review Object"
// @Success 200 {object} structs.RequestReview "updated request review"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /request_review [put]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	requestReviewsGUID, ok := data["request_reviews_guid"].(string)
	if !ok || requestReviewsGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "request_reviews_guid is required", nil)
	}

	my_struct := structs.RequestReview{
		RequestReviewsGUID: requestReviewsGUID,
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update request review", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete request review
// @Description Delete a request review by its GUID
// @Accept json
// @Produce json
// @Param request_review body structs.RequestReview true "Request Review Object with GUID"
// @Success 200 {object} structs.RequestReview "deleted request review"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /request_review [delete]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	requestReviewsGUID, ok := data["request_reviews_guid"].(string)
	if !ok || requestReviewsGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "request_reviews_guid is required", nil)
	}

	my_struct := structs.RequestReview{
		RequestReviewsGUID: requestReviewsGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete request review", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get request review by ID
// @Description Get a specific request review by its GUID
// @Produce json
// @Param request_review_guid path string true "Request Review GUID"
// @Success 200 {object} structs.RequestReview "request review"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /request_review/{request_review_guid} [get]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) GETBYID(ctx echo.Context) error {
	request_review_guid := ctx.Param("request_review_guid")
	if request_review_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "request_review_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(request_review_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch request review", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple request reviews creation
// @Description Create multiple request reviews at once
// @Accept json
// @Produce json
// @Success 501 {object} ErrorResponse "Not Implemented"
// @Router /request_review/multi [post]
// @Security BearerAuth
// @Tags request_review
func (p *RequestReviewHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple request reviews creation is not implemented yet", nil)
}
