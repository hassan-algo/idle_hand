package business

import (
	"example.com/db"
	"example.com/structs"
)

type CatalogBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewCatalogBusiness() *CatalogBusiness {
	return &CatalogBusiness{}
}

func (b *CatalogBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *CatalogBusiness) GET(data interface{}) (interface{}, error) {
	catalogs := structs.Catalogs{
		MyCatalogs: []structs.Catalog{},
	}
	err := b.dbCon.Con.Find(&catalogs.MyCatalogs).Error
	if err != nil {
		return nil, err
	}
	return catalogs, nil
}
func (b *CatalogBusiness) GETBYID(data interface{}) (interface{}, error) {
	catalog_guid, _ := data.(string)
	catalog := structs.Catalog{}
	err := b.dbCon.Con.Where("catalog_guid = ?", catalog_guid).First(&catalog).Error
	if err != nil {
		return nil, err
	}
	return catalog, nil
}
func (b *CatalogBusiness) POST(data interface{}) (interface{}, error) {
	catalog, _ := data.(structs.Catalog)
	err := b.dbCon.Con.Create(&catalog).Error
	if err != nil {
		return nil, err
	}
	return catalog, nil
}
func (b *CatalogBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	catalog, _ := data.(structs.Catalog)
	return catalog, nil
}
func (b *CatalogBusiness) PUT(data interface{}) (interface{}, error) {
	catalog, _ := data.(structs.Catalog)
	err := b.dbCon.Con.Model(&catalog).Where("catalog_guid = ?", catalog.CatalogGUID).Updates(catalog).Error
	if err != nil {
		return nil, err
	}
	return catalog, nil
}
func (b *CatalogBusiness) DELETE(data interface{}) (interface{}, error) {
	catalog, _ := data.(structs.Catalog)
	err := b.dbCon.Con.Model(&catalog).Where("catalog_guid = ?", catalog.CatalogGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return catalog, nil
}
