package business

import (
	"example.com/db"
	"example.com/structs"
)

type CatalogAvailabilityBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewCatalogAvailabilityBusiness() *CatalogAvailabilityBusiness {
	return &CatalogAvailabilityBusiness{}
}

func (b *CatalogAvailabilityBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *CatalogAvailabilityBusiness) GET(data interface{}) (interface{}, error) {
	catalog_availabilitys := structs.CatalogAvailabilitys{
		MyCatalogAvailabilitys: []structs.CatalogAvailability{},
	}
	err := b.dbCon.Con.Find(&catalog_availabilitys.MyCatalogAvailabilitys).Error
	if err != nil {
		return nil, err
	}
	return catalog_availabilitys, nil
}
func (b *CatalogAvailabilityBusiness) GETBYID(data interface{}) (interface{}, error) {
	catalog_availability, _ := data.(structs.CatalogAvailability)
	return catalog_availability, nil
}
func (b *CatalogAvailabilityBusiness) POST(data interface{}) (interface{}, error) {
	catalog_availability, _ := data.(structs.CatalogAvailability)
	err := b.dbCon.Con.Create(&catalog_availability).Error
	if err != nil {
		return nil, err
	}
	return catalog_availability, nil
}
func (b *CatalogAvailabilityBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	catalog_availability, _ := data.(structs.CatalogAvailability)
	return catalog_availability, nil
}
func (b *CatalogAvailabilityBusiness) PUT(data interface{}) (interface{}, error) {
	catalog_availability, _ := data.(structs.CatalogAvailability)
	err := b.dbCon.Con.Model(&catalog_availability).Where("catalog_availability_guid = ?", catalog_availability.CatalogAvailabilityGUID).Updates(catalog_availability).Error
	if err != nil {
		return nil, err
	}
	return catalog_availability, nil
}
func (b *CatalogAvailabilityBusiness) DELETE(data interface{}) (interface{}, error) {
	catalog_availability, _ := data.(structs.CatalogAvailability)
	err := b.dbCon.Con.Model(&catalog_availability).Where("catalog_availability_guid = ?", catalog_availability.CatalogAvailabilityGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return catalog_availability, nil
}
