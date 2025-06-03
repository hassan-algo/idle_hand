package business

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"

	"example.com/db"
	"example.com/extras"
	"example.com/structs"
)

type PasswordManagementBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewPasswordManagementBusiness() *PasswordManagementBusiness {
	return &PasswordManagementBusiness{}
}

func (b *PasswordManagementBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *PasswordManagementBusiness) GET(data interface{}) (interface{}, error) {
	password_managements := structs.PasswordManagements{
		MyPasswordManagements: []structs.PasswordManagement{},
	}
	return password_managements, nil
}
func (b *PasswordManagementBusiness) GETBYID(data interface{}) (interface{}, error) {
	password_management, _ := data.(structs.PasswordManagement)
	return password_management, nil
}
func (b *PasswordManagementBusiness) POST(data interface{}) (interface{}, error) {
	password_management, _ := data.(structs.PasswordManagement)
	return password_management, nil
}
func (b *PasswordManagementBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	password_management, _ := data.(structs.PasswordManagement)
	otp := extras.GenerateSixDigitCode()
	otp_message := strconv.Itoa(otp)
	extras.ForgotPassEmail(password_management.Email, otp_message)
	return "OTP sent to your email", nil
}
func (b *PasswordManagementBusiness) PUT(data interface{}) (interface{}, error) {
	password_management, _ := data.(structs.PasswordManagement)
	// check if the new password and confirm password are the same
	if password_management.NewPassword != password_management.ConfirmPassword {
		return "New password and confirm password are not the same", errors.New("new password and confirm password are not the same")
	}

	// Fetch the current hashed password from the database
	var currentHashedPassword string
	err := b.dbCon.Con.Table("tbl_users").Where("user_guid = ?", password_management.UserGuid).Select("password").Scan(&currentHashedPassword).Error
	if err != nil {
		return "Old password is incorrect", errors.New("old password is incorrect")
	}

	// Hash the provided old password and compare
	hashOld := md5.Sum([]byte(password_management.OldPassword))
	hashOldPassword := hex.EncodeToString(hashOld[:])
	if currentHashedPassword != hashOldPassword {
		return "Old password is incorrect", errors.New("old password is incorrect")
	}

	// Check if the new password is the same as the old password
	hashNew := md5.Sum([]byte(password_management.NewPassword))
	hashNewPassword := hex.EncodeToString(hashNew[:])
	if currentHashedPassword == hashNewPassword {
		return "New password is the same as the old password", errors.New("new password is the same as the old password")
	}

	// Update the password
	err = b.dbCon.Con.Table("tbl_users").Where("user_guid = ?", password_management.UserGuid).Update("password", hashNewPassword).Error
	if err != nil {
		return nil, err
	}
	return "Password updated successfully", nil
}
func (b *PasswordManagementBusiness) DELETE(data interface{}) (interface{}, error) {
	password_management, _ := data.(structs.PasswordManagement)
	return password_management, nil
}
