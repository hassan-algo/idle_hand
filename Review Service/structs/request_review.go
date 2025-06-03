package structs

type RequestReviews struct {
	MyRequestReviews []RequestReview `json:"request_reviews,omitempty"`
}

type RequestReview struct {
	RequestReviewsGUID string    `json:"request_reviews_guid,omitempty" gorm:"column:request_reviews_guid"`
	BusinessGUID       string    `json:"business_guid,omitempty" gorm:"column:business_guid"`
	UserGUID           string    `json:"user_guid,omitempty" gorm:"column:user_guid"`
	ReviewRequested    int       `json:"review_requested,omitempty" gorm:"column:review_requested"`
}

func (RequestReview) TableName() string {
	return "tbl_request_reviews"
}
