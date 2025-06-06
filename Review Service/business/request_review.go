package business

import (
	"example.com/db"
	"example.com/structs"
)

type RequestReviewBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewRequestReviewBusiness() *RequestReviewBusiness {
	return &RequestReviewBusiness{}
}

func (b *RequestReviewBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *RequestReviewBusiness) GET(data interface{}) (interface{}, error) {
	request_reviews := structs.RequestReviews{
		MyRequestReviews: []structs.RequestReview{},
	}
	err := b.dbCon.Con.Find(&request_reviews.MyRequestReviews).Error
	if err != nil {
		return nil, err
	}
	return request_reviews, nil
}
func (b *RequestReviewBusiness) GETBYID(data interface{}) (interface{}, error) {
	request_review_guid, _ := data.(string)
	request_review := structs.RequestReview{}
	err := b.dbCon.Con.Where("request_review_guid = ?", request_review_guid).First(&request_review).Error
	if err != nil {
		return nil, err
	}
	return request_review, nil
}
func (b *RequestReviewBusiness) POST(data interface{}) (interface{}, error) {
	request_review, _ := data.(structs.RequestReview)
	err := b.dbCon.Con.Create(&request_review).Error
	if err != nil {
		return nil, err
	}
	return request_review, nil
}
func (b *RequestReviewBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	request_review, _ := data.(structs.RequestReview)
	return request_review, nil
}
func (b *RequestReviewBusiness) PUT(data interface{}) (interface{}, error) {
	request_review, _ := data.(structs.RequestReview)
	err := b.dbCon.Con.Model(&request_review).Where("request_review_guid = ?", request_review.RequestReviewsGUID).Updates(request_review).Error
	if err != nil {
		return nil, err
	}
	return request_review, nil
}
func (b *RequestReviewBusiness) DELETE(data interface{}) (interface{}, error) {
	request_review, _ := data.(structs.RequestReview)
	err := b.dbCon.Con.Model(&request_review).Where("request_review_guid = ?", request_review.RequestReviewsGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return request_review, nil
}
