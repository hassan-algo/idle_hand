package business

import (
	"example.com/db"
	"example.com/structs"
)

type OffersBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewOffersBusiness() *OffersBusiness {
	return &OffersBusiness{}
}

func (b *OffersBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *OffersBusiness) GET(data interface{}) (interface{}, error) {
	offerss := structs.Offerss{
		MyOfferss: []structs.Offers{},
	}
	err := b.dbCon.Con.Find(&offerss.MyOfferss).Error
	if err != nil {
		return nil, err
	}
	return offerss, nil
}
func (b *OffersBusiness) GETBYID(data interface{}) (interface{}, error) {
	offer_guid, _ := data.(string)
	offers := structs.Offers{}
	err := b.dbCon.Con.Where("offer_guid = ?", offer_guid).First(&offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}
func (b *OffersBusiness) POST(data interface{}) (interface{}, error) {
	offers, _ := data.(structs.Offers)
	err := b.dbCon.Con.Create(&offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}
func (b *OffersBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	offers, _ := data.(structs.Offers)
	return offers, nil
}
func (b *OffersBusiness) PUT(data interface{}) (interface{}, error) {
	offers, _ := data.(structs.Offers)
	err := b.dbCon.Con.Model(&offers).Where("offer_guid = ?", offers.OfferGUID).Updates(offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}
func (b *OffersBusiness) DELETE(data interface{}) (interface{}, error) {
	offers, _ := data.(structs.Offers)
	err := b.dbCon.Con.Model(&offers).Where("offer_guid = ?", offers.OfferGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}
