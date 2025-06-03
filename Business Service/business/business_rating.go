package business

import (
	"example.com/db"
	"example.com/structs"
)

type BusinessRatingBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBusinessRatingBusiness() *BusinessRatingBusiness {
	return &BusinessRatingBusiness{}
}

func (b *BusinessRatingBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BusinessRatingBusiness) GET(data interface{}) (interface{}, error) {
	business_ratings := structs.BusinessRatings{
		MyBusinessRatings: []structs.BusinessRating{
		},
	}
	err := b.dbCon.Con.Find(&business_ratings.MyBusinessRatings).Error
	if err != nil {
		return nil, err
	}
	return business_ratings, nil
}
func (b *BusinessRatingBusiness) GETBYID(data interface{}) (interface{}, error) {
	business_rating_guid, _ := data.(string)
	business_rating := structs.BusinessRating{}
	err := b.dbCon.Con.Where("business_rating_guid = ?", business_rating_guid).First(&business_rating).Error
	if err != nil {
		return nil, err
	}
	return business_rating, nil
}
func (b *BusinessRatingBusiness) POST(data interface{}) (interface{}, error) {
	business_rating, _ := data.(structs.BusinessRating)
	err := b.dbCon.Con.Create(&business_rating).Error
	if err != nil {
		return nil, err
	}
	return business_rating, nil
}
func (b *BusinessRatingBusiness) MULTIPOST(data interface{}) (interface{}, error) {
		business_rating, _ := data.(structs.BusinessRating)
	return business_rating, nil
}
func (b *BusinessRatingBusiness) PUT(data interface{}) (interface{}, error) {
	business_rating, _ := data.(structs.BusinessRating)
	err := b.dbCon.Con.Model(&business_rating).Where("business_rating_guid = ?", business_rating.BusinessRatingGUID).Updates(business_rating).Error
	if err != nil {
		return nil, err
	}
	return business_rating, nil
}
func (b *BusinessRatingBusiness) DELETE(data interface{}) (interface{}, error) {
	business_rating, _ := data.(structs.BusinessRating)
	err := b.dbCon.Con.Model(&business_rating).Where("business_rating_guid = ?", business_rating.BusinessRatingGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return business_rating, nil
}
