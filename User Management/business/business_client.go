package business

import (
	"example.com/db"
	"example.com/structs"
)

type BusinessClientBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBusinessClientBusiness() *BusinessClientBusiness {
	return &BusinessClientBusiness{}
}

func (b *BusinessClientBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BusinessClientBusiness) GET(data interface{}) (interface{}, error) {
	business_clients := structs.BusinessClients{
		MyBusinessClients: []structs.BusinessClient{},
	}
	err := b.dbCon.Con.Find(&business_clients.MyBusinessClients).Error
	if err != nil {
		return nil, err
	}
	return business_clients, nil
}
func (b *BusinessClientBusiness) GETBYID(data interface{}) (interface{}, error) {
	business_client_guid, _ := data.(string)
	business_client := structs.BusinessClient{}
	err := b.dbCon.Con.Where("business_client_guid = ?", business_client_guid).First(&business_client).Error
	if err != nil {
		return nil, err
	}
	return business_client, nil
}
func (b *BusinessClientBusiness) POST(data interface{}) (interface{}, error) {
	business_client, _ := data.(structs.BusinessClient)
	err := b.dbCon.Con.Create(&business_client).Error
	if err != nil {
		return nil, err
	}
	return business_client, nil
}
func (b *BusinessClientBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	business_client, _ := data.(structs.BusinessClient)

	return business_client, nil
}
func (b *BusinessClientBusiness) PUT(data interface{}) (interface{}, error) {
	business_client, _ := data.(structs.BusinessClient)
	err := b.dbCon.Con.Model(&business_client).Where("business_client_guid = ?", business_client.BusinessClientGUID).Updates(business_client).Error
	if err != nil {
		return nil, err
	}
	return business_client, nil
}
func (b *BusinessClientBusiness) DELETE(data interface{}) (interface{}, error) {
	business_client, _ := data.(structs.BusinessClient)
	err := b.dbCon.Con.Model(&business_client).Where("business_client_guid = ?", business_client.BusinessClientGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return business_client, nil
}
