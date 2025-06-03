package structs

type BusinessPoliciess struct {
	MyBusinessPoliciess []BusinessPolicies `json:"business_policiess,omitempty"`
}

type BusinessPolicies struct {
	BusinessPoliciesGUID      string  `json:"business_policies_guid,omitempty" gorm:"column:business_policies_guid"`
	BusinessGUID              string  `json:"business_guid,omitempty" gorm:"column:business_guid"`
	CatalogGUID               string  `json:"catalog_guid,omitempty" gorm:"column:catalog_guid"`
	CancellationHours         float64 `json:"cancellation_hours,omitempty" gorm:"column:cancellation_hours"`
	CancellationAmount        float64 `json:"cancellation_amount,omitempty" gorm:"column:cancellation_amount"`
	NoShowFee                 float64 `json:"no_show_fee,omitempty" gorm:"column:no_show_fee"`
	BookingDepositePercentage float64 `json:"booking_deposite_percentage,omitempty" gorm:"column:booking_deposite_percentage"`
	BookingTerms              string  `json:"booking_terms,omitempty" gorm:"column:booking_terms"`
	BookingPolices            string  `json:"booking_polices,omitempty" gorm:"column:booking_polices"`
}

// TableName sets the custom table name for GORM
func (BusinessPolicies) TableName() string {
	return "tbl_business_policies"
}
