package business

import (
	"example.com/db"
	"example.com/structs"
)

type BookingBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewBookingBusiness() *BookingBusiness {
	return &BookingBusiness{}
}

func (b *BookingBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *BookingBusiness) GET(data interface{}) (interface{}, error) {
	bookings := structs.Bookings{
		MyBookings: []structs.Booking{},
	}
	err := b.dbCon.Con.Find(&bookings.MyBookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (b *BookingBusiness) GETBYID(data interface{}) (interface{}, error) {
	booking_guid, _ := data.(string)
	booking := structs.Booking{}
	err := b.dbCon.Con.Where("booking_guid = ?", booking_guid).First(&booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}
func (b *BookingBusiness) POST(data interface{}) (interface{}, error) {
	booking, _ := data.(structs.Booking)
	err := b.dbCon.Con.Create(&booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}
func (b *BookingBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	booking, _ := data.(structs.Booking)
	return booking, nil
}
func (b *BookingBusiness) PUT(data interface{}) (interface{}, error) {
	booking, _ := data.(structs.Booking)
	err := b.dbCon.Con.Model(&booking).Where("booking_guid = ?", booking.BookingGUID).Updates(booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}

// soft delete
func (b *BookingBusiness) DELETE(data interface{}) (interface{}, error) {
	booking, _ := data.(structs.Booking)
	err := b.dbCon.Con.Model(&booking).Where("booking_guid = ?", booking.BookingGUID).Update("is_deleted", 1).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}
