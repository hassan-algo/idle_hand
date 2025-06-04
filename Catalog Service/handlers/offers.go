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

// validateOffer validates the required fields of an offer
func validateOffer(offer structs.Offers) error {
	if offer.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if offer.OfferTitle == "" {
		return errors.New("offer_title is required")
	}
	if offer.OfferMessage == "" {
		return errors.New("offer_message is required")
	}
	if offer.CustomerType == "" {
		return errors.New("customer_type is required")
	}
	if offer.Medium < 0 {
		return errors.New("medium cannot be negative")
	}
	return nil
}

type OffersHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewOffersHandler() *OffersHandlers {
	return &OffersHandlers{}
}

func (h *OffersHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get all offers
// @Description Get all offers
// @Produce json
// @Success 200 {object} structs.Offers "offers"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers [get]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch offers", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Create a new offer
// @Description Create a new offer with the provided details
// @Accept json
// @Produce json
// @Param offer body structs.Offers true "Offer Object"
// @Success 200 {object} structs.Offers "created offer"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers [post]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	offerTitle, ok := data["offer_title"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid offer_title format", nil)
	}

	offerMessage, ok := data["offer_message"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid offer_message format", nil)
	}

	customerType, ok := data["customer_type"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid customer_type format", nil)
	}

	medium, ok := data["medium"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid medium format", nil)
	}

	offer_guid := uuid.New().String()
	my_struct := structs.Offers{
		OfferGUID:    offer_guid,
		OfferTitle:   offerTitle,
		OfferMessage: offerMessage,
		CustomerType: customerType,
		Medium:       int(medium),
		BusinessGUID: businessGUID,
	}

	// Validate the offer
	if err := validateOffer(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create offer", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Update an offer
// @Description Update an existing offer with the provided details
// @Accept json
// @Produce json
// @Param offer body structs.Offers true "Offer Object"
// @Success 200 {object} structs.Offers "updated offer"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers [put]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if offer_guid exists
	offerGUID, ok := data["offer_guid"].(string)
	if !ok || offerGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "offer_guid is required", nil)
	}

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	offerTitle, ok := data["offer_title"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid offer_title format", nil)
	}

	offerMessage, ok := data["offer_message"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid offer_message format", nil)
	}

	customerType, ok := data["customer_type"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid customer_type format", nil)
	}

	medium, ok := data["medium"].(float64)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid medium format", nil)
	}

	my_struct := structs.Offers{
		OfferGUID:    offerGUID,
		OfferTitle:   offerTitle,
		OfferMessage: offerMessage,
		CustomerType: customerType,
		Medium:       int(medium),
		BusinessGUID: businessGUID,
	}

	// Validate the offer
	if err := validateOffer(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update offer", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Delete an offer
// @Description Delete an offer by its GUID
// @Accept json
// @Produce json
// @Param offer body structs.Offers true "Offer Object with GUID"
// @Success 200 {object} structs.Offers "deleted offer"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers [delete]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	offerGUID, ok := data["offer_guid"].(string)
	if !ok || offerGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "offer_guid is required", nil)
	}

	my_struct := structs.Offers{
		OfferGUID: offerGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete offer", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Get offer by ID
// @Description Get a specific offer by its GUID
// @Produce json
// @Param offer_guid path string true "Offer GUID"
// @Success 200 {object} structs.Offers "offer"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers/{offer_guid} [get]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) GETBYID(ctx echo.Context) error {
	offer_guid := ctx.Param("offer_guid")
	if offer_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "offer_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(offer_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch offer", err)
	}
	return ctx.JSON(http.StatusOK, mydata)
}

// @Summary Multiple offers creation
// @Description Create multiple offers at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST Offers"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /offers/multi [post]
// @Security BearerAuth
// @Tags offers
func (p *OffersHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple offers creation is not implemented yet", nil)
}
