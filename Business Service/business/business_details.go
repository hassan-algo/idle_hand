package business

import (
	"example.com/db"
	"example.com/structs"
)

type BusinessDetailsBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBusinessDetailsBusiness() *BusinessDetailsBusiness {
	return &BusinessDetailsBusiness{}
}

func (b *BusinessDetailsBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BusinessDetailsBusiness) GET(data interface{}) (interface{}, error) {

	business_detailss := structs.BusinessDetailss{
		MyBusinessDetailss: []structs.BusinessDetails{},
	}
	err := b.dbCon.Con.Find(&business_detailss.MyBusinessDetailss).Error
	if err != nil {
		return nil, err
	}
	return business_detailss, nil
}
func (b *BusinessDetailsBusiness) GETBYID(data interface{}) (interface{}, error) {
	business_details, _ := data.(structs.BusinessDetails)
	err := b.dbCon.Con.Where("business_guid = ?", business_details.BusinessGUID).First(&business_details).Error
	if err != nil {
		return nil, err
	}
	return business_details, nil
}
func (b *BusinessDetailsBusiness) POST(data interface{}) (interface{}, error) {
	business_details, _ := data.(structs.BusinessDetails)
	err := b.dbCon.Con.Create(&business_details).Error
	if err != nil {
		return nil, err
	}
	return business_details, nil
}
func (b *BusinessDetailsBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	business_details, _ := data.(structs.BusinessDetails)
	return business_details, nil
}
func (b *BusinessDetailsBusiness) PUT(data interface{}) (interface{}, error) {
	business_details, _ := data.(structs.BusinessDetails)
	err := b.dbCon.Con.Model(&business_details).Where("business_guid = ?", business_details.BusinessGUID).Updates(business_details).Error
	if err != nil {
		return nil, err
	}
	return business_details, nil
}
func (b *BusinessDetailsBusiness) DELETE(data interface{}) (interface{}, error) {
	business_details, _ := data.(structs.BusinessDetails)
	err := b.dbCon.Con.Model(&business_details).Where("business_guid = ?", business_details.BusinessGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return business_details, nil
}
