package structs

type BusinessStaffs struct {
	MyBusinessStaffs []BusinessStaff `json:"business_staffs,omitempty"`
}
type BusinessStaff struct {
	BusinessStaffGUID string `json:"business_staff_guid,omitempty" gorm:"column:business_staff_guid"`
	StaffGUID         string `json:"staff_guid,omitempty" gorm:"column:staff_guid"`
	BusinessGUID      string `json:"business_guid,omitempty" gorm:"column:business_guid"`
	Profession        string `json:"profession,omitempty" gorm:"column:profession"`
	ArrivalTime       string `json:"arrival_time,omitempty" gorm:"column:arrival_time"`
	LeaveTime         string `json:"leave_time,omitempty" gorm:"column:leave_time"`
}

// Optional: specify table name for GORM
func (BusinessStaff) TableName() string {
	return "tbl_business_staff"
}
