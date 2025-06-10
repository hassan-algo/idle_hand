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

// validateCatalog validates the required fields of a catalog
func validateCatalog(catalog structs.Catalog) error {
	if catalog.BusinessGUID == "" {
		return errors.New("business_guid is required")
	}
	if catalog.CatalogCategory == "" {
		return errors.New("catalog_category is required")
	}
	if catalog.CatalogName == "" {
		return errors.New("catalog_name is required")
	}
	if catalog.CatalogDescription == "" {
		return errors.New("catalog_description is required")
	}
	if catalog.CatalogPrice == "" {
		return errors.New("catalog_price is required")
	}
	if catalog.CatalogOffering == "" {
		return errors.New("catalog_offering is required")
	}
	return nil
}

type CatalogHandlers struct {
	apiBusiness apis.APIBusiness
}

func NewCatalogHandler() *CatalogHandlers {
	return &CatalogHandlers{}
}

func (h *CatalogHandlers) Connect(business apis.APIBusiness) error {
	h.apiBusiness = business
	return nil
}

// @Summary Get all catalogs
// @Description Get all catalogs
// @Produce json
// @Success 200 {object} structs.Catalog "catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog [get]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) GET(ctx echo.Context) error {
	mydata, err := p.apiBusiness.GET(nil)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch catalogs", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Catalogs fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Create a new catalog
// @Description Create a new catalog with the provided details
// @Accept json
// @Produce json
// @Param catalog body structs.Catalog true "Catalog Object"
// @Success 200 {object} structs.Catalog "created catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog [post]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) POST(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	catalogCategory, ok := data["catalog_category"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_category format", nil)
	}

	catalogName, ok := data["catalog_name"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_name format", nil)
	}

	catalogDescription, ok := data["catalog_description"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_description format", nil)
	}

	catalogPrice, ok := data["catalog_price"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_price format", nil)
	}

	catalogOffering, ok := data["catalog_offering"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_offering format", nil)
	}

	catalogPhoto, ok := data["catalog_photo"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_photo format", nil)
	}

	assignedStaffGUID, ok := data["assigned_staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid assigned_staff_guid format", nil)
	}

	catalog_guid := uuid.New().String()
	my_struct := structs.Catalog{
		CatalogGUID:        catalog_guid,
		BusinessGUID:       businessGUID,
		CatalogCategory:    catalogCategory,
		CatalogName:        catalogName,
		CatalogDescription: catalogDescription,
		CatalogPrice:       catalogPrice,
		CatalogOffering:    catalogOffering,
		CatalogPhoto:       catalogPhoto,
		AssignedStaffGUID:  assignedStaffGUID,
	}

	// Validate the catalog
	if err := validateCatalog(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.POST(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to create catalog", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Catalog created successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Update a catalog
// @Description Update an existing catalog with the provided details
// @Accept json
// @Produce json
// @Param catalog body structs.Catalog true "Catalog Object"
// @Success 200 {object} structs.Catalog "updated catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog [put]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) PUT(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	// Check if catalog_guid exists
	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok || catalogGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_guid is required", nil)
	}

	// Type assertion with error handling
	businessGUID, ok := data["business_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid business_guid format", nil)
	}

	catalogCategory, ok := data["catalog_category"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_category format", nil)
	}

	catalogName, ok := data["catalog_name"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_name format", nil)
	}

	catalogDescription, ok := data["catalog_description"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_description format", nil)
	}

	catalogPrice, ok := data["catalog_price"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_price format", nil)
	}

	catalogOffering, ok := data["catalog_offering"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_offering format", nil)
	}

	catalogPhoto, ok := data["catalog_photo"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid catalog_photo format", nil)
	}

	assignedStaffGUID, ok := data["assigned_staff_guid"].(string)
	if !ok {
		return handleError(ctx, http.StatusBadRequest, "Invalid assigned_staff_guid format", nil)
	}

	my_struct := structs.Catalog{
		CatalogGUID:        catalogGUID,
		BusinessGUID:       businessGUID,
		CatalogCategory:    catalogCategory,
		CatalogName:        catalogName,
		CatalogDescription: catalogDescription,
		CatalogPrice:       catalogPrice,
		CatalogOffering:    catalogOffering,
		CatalogPhoto:       catalogPhoto,
		AssignedStaffGUID:  assignedStaffGUID,
	}

	// Validate the catalog
	if err := validateCatalog(my_struct); err != nil {
		return handleError(ctx, http.StatusBadRequest, "Validation failed", err)
	}

	mydata, err := p.apiBusiness.PUT(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to update catalog", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Catalog updated successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Get catalog by ID
// @Description Get a specific catalog by its GUID
// @Produce json
// @Param catalog_guid path string true "Catalog GUID"
// @Success 200 {object} structs.Catalog "catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog/{catalog_guid} [get]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) GETBYID(ctx echo.Context) error {
	catalog_guid := ctx.Param("catalog_guid")
	if catalog_guid == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_guid is required", nil)
	}

	mydata, err := p.apiBusiness.GETBYID(catalog_guid)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to fetch catalog", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Catalog fetched successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}

// @Summary Multiple catalogs creation
// @Description Create multiple catalogs at once
// @Accept json
// @Produce json
// @Success 200 {string} string "MULTIPOST Catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog/multi [post]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) MULTIPOST(ctx echo.Context) error {
	return handleError(ctx, http.StatusNotImplemented, "Multiple catalogs creation is not implemented yet", nil)
}

// @Summary Delete a catalog
// @Description Delete a catalog by its GUID
// @Accept json
// @Produce json
// @Param catalog body structs.Catalog true "Catalog Object with GUID"
// @Success 200 {object} structs.Catalog "deleted catalog"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 404 {object} ErrorResponse "Not Found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /catalog [delete]
// @Security BearerAuth
// @Tags catalog
func (p *CatalogHandlers) DELETE(ctx echo.Context) error {
	data := extras.GetJSONRawBody(ctx)

	catalogGUID, ok := data["catalog_guid"].(string)
	if !ok || catalogGUID == "" {
		return handleError(ctx, http.StatusBadRequest, "catalog_guid is required", nil)
	}

	my_struct := structs.Catalog{
		CatalogGUID: catalogGUID,
	}

	mydata, err := p.apiBusiness.DELETE(my_struct)
	if err != nil {
		return handleError(ctx, http.StatusInternalServerError, "Failed to delete catalog", err)
	}
	return ctx.JSON(http.StatusOK, structs.Response{
		Valid:       true,
		Message:     "Catalog deleted successfully",
		Data:        mydata,
		Status_code: http.StatusOK,
	})
}
