package business

import (
	"example.com/db"
	"example.com/structs"
)

type BusinessStaffBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBusinessStaffBusiness() *BusinessStaffBusiness {
	return &BusinessStaffBusiness{}
}

func (b *BusinessStaffBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BusinessStaffBusiness) GET(data interface{}) (interface{}, error) {
	business_staffs := structs.BusinessStaffs{
		MyBusinessStaffs: []structs.BusinessStaff{},
	}
	err := b.dbCon.Con.Find(&business_staffs.MyBusinessStaffs).Error
	if err != nil {
		return nil, err
	}
	return business_staffs, nil
}
func (b *BusinessStaffBusiness) GETBYID(data interface{}) (interface{}, error) {
	business_staff_guid, _ := data.(string)
	business_staff := structs.BusinessStaff{}
	err := b.dbCon.Con.Where("business_staff_guid = ?", business_staff_guid).First(&business_staff).Error
	if err != nil {
		return nil, err
	}
	return business_staff, nil
}
func (b *BusinessStaffBusiness) POST(data interface{}) (interface{}, error) {
	business_staff, _ := data.(structs.BusinessStaff)
	err := b.dbCon.Con.Create(&business_staff).Error
	if err != nil {
		return nil, err
	}
	return business_staff, nil
}
func (b *BusinessStaffBusiness) MULTIPOST(data interface{}) (interface{}, error) {
		business_staff, _ := data.(structs.BusinessStaff)
	return business_staff, nil
}
func (b *BusinessStaffBusiness) PUT(data interface{}) (interface{}, error) {
	business_staff, _ := data.(structs.BusinessStaff)
	err := b.dbCon.Con.Model(&business_staff).Where("business_staff_guid = ?", business_staff.BusinessStaffGUID).Updates(business_staff).Error
	if err != nil {
		return nil, err
	}
	return business_staff, nil
}
func (b *BusinessStaffBusiness) DELETE(data interface{}) (interface{}, error) {
	business_staff, _ := data.(structs.BusinessStaff)
	err := b.dbCon.Con.Model(&business_staff).Where("business_staff_guid = ?", business_staff.BusinessStaffGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return business_staff, nil
}
