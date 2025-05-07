package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	authBusiness apis.AuthBusiness
}

func NewAuthHandler() *AuthHandlers {
	return &AuthHandlers{}
}

// @Summary Register new user
// @Description
// @Accept json
// @Param  profile body structs.MyAuth true "Register user"
// @Success 200 {object} structs.Response
// @Router /auth [post]
// @Tags auth
func (h *AuthHandlers) Authentication(ec echo.Context) error {
	body := extras.GetJSONRawBody(ec)
	// extras.LogThisWithActor(i.e, "", body["email"].(string))

	email := body["email"].(string)
	password := body["password"].(string)

	data, err := h.authBusiness.Authentication(email, password)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, data)
	}
	return ec.JSON(http.StatusOK, data)
}

// Authenticate wraps a function with authentication logic.
func (h *AuthHandlers) Authenticate(f func(ec echo.Context) error, role ...string) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		// get headers
		authHeader := ec.Request().Header.Get("Authorization")
		if authHeader == "" {
			return errors.New("auth Failed1")
		}
		splitHeader := strings.Split(authHeader, " ")
		if len(splitHeader) < 3 {
			log.Println(authHeader)
			return errors.New("invalid Token!")
		}
		userGuid := splitHeader[2]
		token := splitHeader[1]

		err, userGuid, returnedRole := h.authBusiness.Authenticate(userGuid, token)

		if err != nil {

			res := structs.Response{
				Valid:   false,
				Message: "UnAuthorized Request1",
				Data:    nil,
			}
			return ec.JSON(http.StatusUnauthorized, res)
		}

		if !extras.Contains(role, returnedRole) {
			res := structs.Response{
				Valid:   false,
				Message: "UnAuthorized Request2",
				Data:    nil,
			}

			return ec.JSON(http.StatusUnauthorized, res)
		}

		ec.Set("user_guid", userGuid)
		// Proceed with the original function if authentication is successful
		return f(ec)
	}
}

// Connect is required to fulfill the apis.AuthHandler interface
func (h *AuthHandlers) Connect(business apis.AuthBusiness) error {
	h.authBusiness = business
	return nil
}

func (h *AuthHandlers) Middleware(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		return f(ec)
	}
}

func (h *AuthHandlers) Decorate(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		return f(ec)
	}
}
