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
	"github.com/google/uuid"
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

	// // First check if email exists
	// checkEmailQuery := "SELECT userguid FROM tbl_users WHERE LOWER(email) = LOWER($1) AND is_deleted = '0'"
	// var existingUserGuid string
	// err := b.dbCon.Con.QueryRow(checkEmailQuery, email).Scan(&existingUserGuid)

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		// Email doesn't exist, create new user
	// 		newUserGuid := uuid.New().String()  // Generate new UUID for user
	// 		loginToken := extras.GetSecretKey() // Generate new login token

	// 		insertQuery := "INSERT INTO tbl_users (userguid, email, password, login_token) VALUES ($1, $2, $3, $4)"
	// 		_, err := b.dbCon.Con.Exec(insertQuery, newUserGuid, email, hashPassword, loginToken)
	// 		if err != nil {
	// 			return structs.Response{Valid: false, Message: "Failed to create user: " + err.Error(), Data: nil}, err
	// 		}

	// 		// Create response for new user
	// 		claims := jwt.MapClaims{
	// 			"exp":      time.Now().Add(730 * time.Hour).Unix(),
	// 			"userguid": newUserGuid,
	// 			"email":    email,
	// 		}

	// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 		signedToken, err := token.SignedString([]byte(loginToken))
	// 		if err != nil {
	// 			return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
	// 		}

	// 		resUser := structs.ResponseUserWithToken{
	// 			Name:       "", // New user has no name yet
	// 			ProfilePic: "", // New user has no profile pic yet
	// 			Email:      email,
	// 			Token:      signedToken + " " + newUserGuid,
	// 		}

	// 		return structs.Response{
	// 			Valid:   true,
	// 			Message: "User created successfully",
	// 			Data:    resUser,
	// 		}, nil
	// 	}
	// 	return structs.Response{Valid: false, Message: "Database error: " + err.Error(), Data: nil}, err
	// }

	// Email exists, proceed with normal login flow
	query := "SELECT user_guid, profile_pic, email, login_token, last_login FROM tbl_users WHERE (LOWER(email) = LOWER($1)) AND password = $2 AND is_deleted = '0'"
	rowsRs, err := b.dbCon.Con.Query(query, email, hashPassword)

	if err != nil {
		return structs.Response{Valid: false, Message: "Auth Failed! " + err.Error(), Data: nil}, err
	}
	defer rowsRs.Close()

	results := make([]structs.Credentials, 0)

	var obj structs.Credentials
	for rowsRs.Next() {
		var lastLogin sql.NullTime
		err := rowsRs.Scan(&obj.UserGuid, &obj.ProfilePic, &obj.Email, &obj.Login_Token, &lastLogin)
		if err != nil {
			return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
		}
		results = append(results, obj)

		if lastLogin.Valid && time.Since(lastLogin.Time) < 730*time.Hour {
			obj.Login_Token, _ = b.getLoginToken(obj.UserGuid)
		} else {
			obj.Login_Token = extras.GetSecretKey()
			updateTokenQuery := "UPDATE tbl_users SET login_token = $1, last_login = NOW() WHERE userguid = $2"
			_, err := b.dbCon.Con.Exec(updateTokenQuery, obj.Login_Token, obj.UserGuid)
			if err != nil {
				return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
			}
		}
	}

	if len(results) < 1 {
		return structs.Response{Valid: false, Message: "Invalid credentials", Data: nil}, errors.New("Invalid credentials")
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
		Valid:   true,
		Message: "User logged in successfully",
		Data:    resUser,
	}

	return res, nil
}
func (b *AuthBusiness) Register(email string, password string) (interface{}, error) {
	if err := b.dbCon.CheckTimeOut(); err != nil {
		log.Println(err.Error())
	}
	hash := md5.Sum([]byte(password))
	hashPassword := hex.EncodeToString(hash[:])

	// First check if email exists
	checkEmailQuery := "SELECT user_guid FROM tbl_users WHERE LOWER(email) = LOWER($1) AND is_deleted = '0'"
	var existingUserGuid string
	err := b.dbCon.Con.QueryRow(checkEmailQuery, email).Scan(&existingUserGuid)

	if err != nil {
		if err == sql.ErrNoRows {
			// Email doesn't exist, create new user
			newUserGuid := uuid.New().String()  // Generate new UUID for user
			loginToken := extras.GetSecretKey() // Generate new login token

			insertQuery := "INSERT INTO tbl_users (user_guid, email, password, login_token) VALUES ($1, $2, $3, $4)"
			_, err := b.dbCon.Con.Exec(insertQuery, newUserGuid, email, hashPassword, loginToken)
			if err != nil {
				return structs.Response{Valid: false, Message: "Failed to create user: " + err.Error(), Data: nil}, err
			}

			// Create response for new user
			claims := jwt.MapClaims{
				"exp":      time.Now().Add(730 * time.Hour).Unix(),
				"userguid": newUserGuid,
				"email":    email,
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedToken, err := token.SignedString([]byte(loginToken))
			if err != nil {
				return structs.Response{Valid: false, Message: err.Error(), Data: nil}, err
			}

			resUser := structs.ResponseUserWithToken{
				Name:       "", // New user has no name yet
				ProfilePic: "", // New user has no profile pic yet
				Email:      email,
				Token:      signedToken + " " + newUserGuid,
			}

			return structs.Response{
				Valid:   true,
				Message: "User created successfully",
				Data:    resUser,
			}, nil
		}
		return structs.Response{Valid: false, Message: "Database error: " + err.Error(), Data: nil}, err
	}

	return structs.Response{Valid: false, Message: "Email already exists", Data: nil}, errors.New("Email already exists")
}

func (b *AuthBusiness) Authenticate(userGuid string, token string) (error, string, string) {

	var (
		JWT_KEY         string
		updatedUserGuid string
		// role            string
	)

	err := b.dbCon.Con.QueryRow("SELECT login_token, user_guid FROM tbl_users WHERE user_guid = $1", userGuid).Scan(&JWT_KEY, &updatedUserGuid)
	if err == sql.ErrNoRows {
		// extras.LogThisWithActor(i.e, "Can't get any rows", "") //
		return fmt.Errorf("auth failed 2: %v", err), "", ""
	} else if err != nil {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return fmt.Errorf("server error: %v", err), "", ""
	}

	if JWT_KEY != "" {
		claims := &jwt.StandardClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		})
		if err != nil {
			// extras.LogThisWithActor(i.e, err.Error(), "") //

			return fmt.Errorf("auth failed 3: %v", err), "", ""
		}

		// extras.LogThisWithActor(i.e, "", "Candidate")
		return nil, updatedUserGuid, "" //
	} else {
		// extras.LogThisWithActor(i.e, err.Error(), "")
		return fmt.Errorf("auth failed 5: %v", err), "", ""
	}
}

// GetLoginToken retrieves the current login token for a user from the database.
// It takes the user's unique identifier (userguid) as an argument and returns the login token and any potential errors encountered.
func (b *AuthBusiness) getLoginToken(userGuid string) (string, error) {
	// Define the SQL query to fetch the login token for the given userguid
	query := `SELECT login_token FROM tbl_users WHERE user_guid = $1`
	// Prepare and execute the query
	row := b.dbCon.Con.QueryRow(query, userGuid)

	var loginToken string
	err := row.Scan(&loginToken)
	if err != nil {
		// If an error occurs (e.g., no matching record), return an empty string and the error
		return "", err
	}
	// Return the fetched login token and nil error if the operation is successful
	return loginToken, nil
}
