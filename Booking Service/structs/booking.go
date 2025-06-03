package structs

type Bookings struct {
	MyBookings []Booking `json:"bookings,omitempty"`
}

type Booking struct {
	BookingGUID      string `json:"booking_guid,omitempty" gorm:"column:booking_guid"`
	BusinessGUID     string `json:"business_guid,omitempty" gorm:"column:business_guid"`
	ClientGUID       string `json:"client_guid,omitempty" gorm:"column:client_guid"`
	ServiceGUID      string `json:"service_guid,omitempty" gorm:"column:service_guid"`
	BookingDate      string `json:"booking_date,omitempty" gorm:"column:booking_date"`
	BookingTime      string `json:"booking_time,omitempty" gorm:"column:booking_time"`
	BookingStatus    string `json:"booking_status,omitempty" gorm:"column:booking_status"`
	CustomerRelation string `json:"customer_relation,omitempty" gorm:"column:customer_relation"`
	AssignStaffGUID  string `json:"assign_staff_guid,omitempty" gorm:"column:assign_staff_guid"`
	PaymentMethod    string `json:"payment_method,omitempty" gorm:"column:payment_method"`
}

func (Booking) TableName() string {
	return "tbl_booking"
}
