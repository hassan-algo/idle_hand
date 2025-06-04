package business

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"time"

	"example.com/db"
	"example.com/extras"
	"example.com/structs"
	"github.com/golang-jwt/jwt"
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
func (b *AuthBusiness) Authentication(email string, password string) (interface{}, error) {
	if err := b.dbCon.CheckTimeOut(); err != nil {
		log.Println(err.Error())
	}
	hash := md5.Sum([]byte(password))
	hashPassword := hex.EncodeToString(hash[:])

	// Extended SQL query to fetch last_login along with other user details
	query := "SELECT user_guid, profile_pic, email, login_token, last_login FROM tbl_users WHERE (LOWER(email) = LOWER($1)) AND password = $2 AND is_deleted = '0'"
	rowsRs, err := b.dbCon.Con.Raw(query, email, hashPassword).Rows()

	if err != nil {
		return structs.Response{Valid: false, Message: "Auth Failed! " + err.Error(), Data: nil}, err
	}
	defer rowsRs.Close()

	results := make([]structs.Credentials, 0)

	var obj structs.Credentials
	for rowsRs.Next() {
		var lastLogin sql.NullTime // Handling possible null values for last_login
		err := rowsRs.Scan(&obj.UserGuid, &obj.ProfilePic, &obj.Email, &obj.Login_Token, &lastLogin)
		if err != nil {
			return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
		}
		results = append(results, obj)
		fmt.Println("user:", results)
		if lastLogin.Valid && time.Since(lastLogin.Time) < 730*time.Hour {
			// If last_login was less than 24 hours ago, use the existing login_token
			obj.Login_Token, _ = b.getLoginToken(obj.UserGuid) // Assuming this function fetches the existing token
		} else {
			// Otherwise, generate a new token and update it in the database
			obj.Login_Token = extras.GetSecretKey()
			updateTokenQuery := "UPDATE tbl_users SET login_token = $1, last_login = NOW() WHERE userguid = $2"
			err := b.dbCon.Con.Exec(updateTokenQuery, obj.Login_Token, obj.UserGuid).Error
			if err != nil {
				return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
			}
		}
	}

	if len(results) < 1 {
		return structs.Response{Valid: false, Message: "Data not received", Data: nil}, errors.New("No Data Found!")
	}

	claims := jwt.MapClaims{
		"exp":      time.Now().Add(730 * time.Hour).Unix(),
		"userguid": obj.UserGuid,
		"email":    obj.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(obj.Login_Token))
	if err != nil {
		return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
	}

	resUser := structs.ResponseUserWithToken{
		UserGuid:   results[0].UserGuid,
		ProfilePic: obj.ProfilePic,
		Email:      obj.Email,
		Token:      signedToken + " " + results[0].UserGuid,
	}

	res := structs.Response{
		Valid: true,
		// Message: obj.UserType,
		Data: resUser,
	}

	return res, nil
}

func (b *AuthBusiness) Authenticate(userGuid string, token string) (error, string, string) {

	var (
		JWT_KEY         string
		updatedUserGuid string
		// role            string
	)

	err := b.dbCon.Con.Raw("SELECT login_token, user_guid FROM tbl_users WHERE user_guid = $1", userGuid).Row().Scan(&JWT_KEY, &updatedUserGuid)
	if err == sql.ErrNoRows {
		// extras.LogThisWithActor(i.e, "Can't get any rows", "") //
		return errors.New("auth Failed2"), "", ""
	} else if err != nil {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return errors.New("server Error"), "", ""
	}

	if JWT_KEY != "" {
		claims := &jwt.StandardClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})
		if err != nil {
			// extras.LogThisWithActor(i.e, err.Error(), "") //

			return errors.New("auth Failed3"), "", ""
		}

		// extras.LogThisWithActor(i.e, "", "Candidate")
		return nil, updatedUserGuid, "" //
	} else {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return errors.New("auth Failed5"), "", ""
	}
}

// GetLoginToken retrieves the current login token for a user from the database.
// It takes the user's unique identifier (userguid) as an argument and returns the login token and any potential errors encountered.
func (b *AuthBusiness) getLoginToken(userGuid string) (string, error) {
	// Define the SQL query to fetch the login token for the given userguid
	query := `SELECT login_token FROM tbl_users WHERE user_guid = $1`
	// Prepare and execute the query
	row := b.dbCon.Con.Raw(query, userGuid).Row()

	var loginToken string
	err := row.Scan(&loginToken)
	if err != nil {
		// If an error occurs (e.g., no matching record), return an empty string and the error
		return "", err
	}
	// Return the fetched login token and nil error if the operation is successful
	return loginToken, nil
}
