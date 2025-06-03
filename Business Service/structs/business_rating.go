package structs

type BusinessRatings struct {
	MyBusinessRatings []BusinessRating `json:"business_ratings,omitempty"`
}

type BusinessRating struct {
	BusinessGUID       string `json:"business_guid,omitempty" gorm:"column:business_guid"`
	UserGUID           string `json:"user_guid,omitempty" gorm:"column:user_guid"`
	BusinessRatingGUID string `json:"business_rating_guid,omitempty" gorm:"column:business_rating_guid"`
	ReviewText         string `json:"review_text,omitempty" gorm:"column:review_text"`
	Rating             string `json:"rating,omitempty" gorm:"column:rating"`
}

// Optional: specify the table name explicitly
func (BusinessRating) TableName() string {
	return "tbl_business_rating"
}
