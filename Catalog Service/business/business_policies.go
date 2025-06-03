package business

import (
	"example.com/db"
	"example.com/structs"
)

type BusinessPoliciesBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBusinessPoliciesBusiness() *BusinessPoliciesBusiness {
	return &BusinessPoliciesBusiness{}
}

func (b *BusinessPoliciesBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BusinessPoliciesBusiness) GET(data interface{}) (interface{}, error) {

	business_policiess := structs.BusinessPoliciess{
		MyBusinessPoliciess: []structs.BusinessPolicies{},
	}
	err := b.dbCon.Con.Find(&business_policiess.MyBusinessPoliciess).Error
	if err != nil {
		return nil, err
	}
	return business_policiess, nil
}
func (b *BusinessPoliciesBusiness) GETBYID(data interface{}) (interface{}, error) {
	business_policies_guid, _ := data.(string)
	business_policies := structs.BusinessPolicies{}
	err := b.dbCon.Con.Where("business_policies_guid = ?", business_policies_guid).First(&business_policies).Error
	if err != nil {
		return nil, err
	}
	return business_policies, nil
}
func (b *BusinessPoliciesBusiness) POST(data interface{}) (interface{}, error) {
	business_policies, _ := data.(structs.BusinessPolicies)
	err := b.dbCon.Con.Create(&business_policies).Error
	if err != nil {
		return nil, err
	}
	return business_policies, nil
}
func (b *BusinessPoliciesBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	business_policies, _ := data.(structs.BusinessPolicies)
	return business_policies, nil
}
func (b *BusinessPoliciesBusiness) PUT(data interface{}) (interface{}, error) {
	business_policies, _ := data.(structs.BusinessPolicies)
	err := b.dbCon.Con.Model(&business_policies).Where("business_policies_guid = ?", business_policies.BusinessPoliciesGUID).Updates(business_policies).Error
	if err != nil {
		return nil, err
	}
	return business_policies, nil
}
func (b *BusinessPoliciesBusiness) DELETE(data interface{}) (interface{}, error) {
	business_policies, _ := data.(structs.BusinessPolicies)
	err := b.dbCon.Con.Model(&business_policies).Where("business_policies_guid = ?", business_policies.BusinessPoliciesGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return business_policies, nil
}
