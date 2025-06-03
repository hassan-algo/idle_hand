package structs

type CatalogAvailabilitys struct {
	MyCatalogAvailabilitys []CatalogAvailability `json:"catalog_availabilitys,omitempty"`
}

type CatalogAvailability struct {
	CatalogAvailabilityGUID string  `json:"catalog_availability_guid,omitempty" gorm:"column:catalog_availability_guid"`
	CatalogGUID             string  `json:"catalog_guid,omitempty" gorm:"column:catalog_guid"`
	BusinessGUID            string  `json:"business_guid,omitempty" gorm:"column:business_guid"`
	DayOfWeek               string  `json:"day_of_week,omitempty" gorm:"column:day_of_week"`
	HoursPerDay             float64 `json:"hours_per_day,omitempty" gorm:"column:hours_per_day"`
	SlotsPerDay             string  `json:"slots_per_day,omitempty" gorm:"column:slots_per_day"`
	NumberOfBreaks          float64 `json:"number_of_breaks,omitempty" gorm:"column:number_of_breaks"`
	TimePerBreak            string  `json:"time_per_break,omitempty" gorm:"column:time_per_break"`
	BufferPerAppointment    float64 `json:"buffer_per_appointment,omitempty" gorm:"column:buffer_per_appointment"`
	AcceptSameDayBooking    float64 `json:"accept_same_day_booking,omitempty" gorm:"column:accept_same_day_booking"`
}

// TableName sets the custom table name for GORM
func (CatalogAvailability) TableName() string {
	return "tbl_catalog_availability"
}
