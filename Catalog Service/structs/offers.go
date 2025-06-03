package structs

type Offerss struct {
	MyOfferss []Offers `json:"offerss,omitempty"`
}

type Offers struct {
	OfferGUID    string `json:"offer_guid,omitempty" gorm:"column:offer_guid"`
	OfferTitle   string `json:"offer_title,omitempty" gorm:"column:offer_title"`
	OfferMessage string `json:"offer_message,omitempty" gorm:"column:offer_message"`
	CustomerType string `json:"customer_type,omitempty" gorm:"column:customer_type"`
	Medium       int    `json:"medium,omitempty" gorm:"column:medium"`
	BusinessGUID string `json:"business_guid,omitempty" gorm:"column:business_guid"`
}

// TableName sets the custom table name for GORM
func (Offers) TableName() string {
	return "tbl_offers"
}
