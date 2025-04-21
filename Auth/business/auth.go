package business

import (
	"example.com/db"
	"example.com/structs"
)

type AuthBusiness struct {
	dbCon *db.DatabaseConnection
}

func NewAuthBusiness() *AuthBusiness {
	return &AuthBusiness{}
}

func (b *AuthBusiness) Connect(dbConnection *db.DatabaseConnection) error {
	b.dbCon = dbConnection
	return nil
}

func (b *AuthBusiness) GET(data interface{}) (interface{}, error) {
	auths := structs.Auths{
		MyAuths: []structs.Auth{
		},
	}
	return auths, nil
}
func (b *AuthBusiness) GETBYID(data interface{}) (interface{}, error) {
	auth, _ := data.(structs.Auth)
	return auth, nil
}
func (b *AuthBusiness) POST(data interface{}) (interface{}, error) {
	auth, _ := data.(structs.Auth)
	return auth, nil
}
func (b *AuthBusiness) MULTIPOST(data interface{}) (interface{}, error) {
	auth, _ := data.(structs.Auth)
	return auth, nil
}
func (b *AuthBusiness) PUT(data interface{}) (interface{}, error) {
	auth, _ := data.(structs.Auth)
	return auth, nil
}
func (b *AuthBusiness) DELETE(data interface{}) (interface{}, error) {
	auth, _ := data.(structs.Auth)
	return auth, nil
}
