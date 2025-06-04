package handlers

import (
	"fmt"
	"net/http"
	"strings"

	// "errors"
	// "example.com/business"
	"example.com/apis"
	"example.com/extras"
	"example.com/structs"
	"github.com/labstack/echo/v4"
	// "github.com/dgrijalva/jwt-go"
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
			res := structs.Response{
				Valid:   false,
				Message: "authentication not found!",
				Data:    nil,
			}
			return ec.JSON(http.StatusUnauthorized, res)
		}
		// fmt.Println("authHeader", authHeader)
		splitHeader := strings.Split(authHeader, " ")
		// fmt.Println("splitHeader", splitHeader)

		if len(splitHeader) < 3 {
			res := structs.Response{
				Valid:   false,
				Message: "invalid authentication token!",
				Data:    nil,
			}
			return ec.JSON(http.StatusUnauthorized, res)
		}

		userGuid := splitHeader[2]
		token := splitHeader[1]
		fmt.Println("userGuid", userGuid)
		fmt.Println("token", token)

		err, userGuid, _ := h.authBusiness.Authenticate(userGuid, token)

		if err != nil {

			res := structs.Response{
				Valid:   false,
				Message: err.Error(),
				Data:    nil,
			}
			return ec.JSON(http.StatusUnauthorized, res)
		}

		// if !extras.Contains(role, returnedRole) {
		// 	res := structs.Response{
		// 		Valid:   false,
		// 		Message: "permissions not found!",
		// 		Data:    nil,
		// 	}

		// 	return ec.JSON(http.StatusUnauthorized, res)
		// }

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
		// if err := h.CheckAuth(ec); err != nil {
		// 	res := structs.Response{
		// 		Valid:   false,
		// 		Message: err.Error(),
		// 		Data:    nil,
		// 	}
		// 	return ec.JSON(http.StatusForbidden, res)
		// }
		return f(ec)
	}
}

func (h *AuthHandlers) Decorate(f func(ec echo.Context) error) func(ec echo.Context) error {
	return func(ec echo.Context) error {
		return f(ec)
	}
}
