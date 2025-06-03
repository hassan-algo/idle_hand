package structs

import "time"

type BusinessDetailss struct {
	MyBusinessDetailss []BusinessDetails `json:"business_detailss,omitempty"`
}

type BusinessDetails struct {
	BusinessName          string    `json:"business_name,omitempty" gorm:"column:business_name"`
	IndustryType          string    `json:"industry_type,omitempty" gorm:"column:industry_type"`
	BusinessDescription   string    `json:"business_description,omitempty" gorm:"column:business_description"`
	BusinessPhoneNumber   string    `json:"business_phone_number,omitempty" gorm:"column:business_phone_number"`
	BusinessLogo          string    `json:"business_logo,omitempty" gorm:"column:business_logo"`
	BusinessPhoto         string    `json:"business_photo,omitempty" gorm:"column:business_photo"`
	BusinessEmail         string    `json:"business_email,omitempty" gorm:"column:business_email"`
	BusinessAddress1      string    `json:"business_address1,omitempty" gorm:"column:business_address1"`
	BusinessAddress2      string    `json:"business_address2,omitempty" gorm:"column:business_address_2"`
	BusinessCity          string    `json:"business_city,omitempty" gorm:"column:business_city"`
	BusinessZipCode       string    `json:"business_zip_code,omitempty" gorm:"column:business_zip_code"`
	BusinessSocialMedia   string    `json:"business_social_media,omitempty" gorm:"column:business_social_media"`
	BusinessPaymentMethod string    `json:"business_payment_method,omitempty" gorm:"column:business_payment_method"`
	BusinessGUID          string    `json:"business_guid,omitempty" gorm:"column:business_guid"`
	CreatedAtRaw          time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	BusinessOTP           string    `json:"business_otp,omitempty" gorm:"column:business_otp"`
	BusinessLicense       string    `json:"business_license,omitempty" gorm:"column:business_license"`
	BusinessLocation      string    `json:"business_location,omitempty" gorm:"column:business_location"`
	IsDeleted             int       `json:"is_deleted,omitempty" gorm:"column:is_deleted"`
}

func (BusinessDetails) TableName() string {
	return "tbl_business_details"
}
