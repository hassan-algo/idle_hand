package apis

import (
	"example.com/db"
	"github.com/labstack/echo/v4"
)

type APIBusiness interface {
	Connect(*db.DatabaseConnection) error
	GET(interface{}) (interface{}, error)
	POST(interface{}) (interface{}, error)
	MULTIPOST(interface{}) (interface{}, error)
	PUT(interface{}) (interface{}, error)
	GETBYID(interface{}) (interface{}, error)
	DELETE(interface{}) (interface{}, error)
}

type APIHandler interface {
	Connect(APIBusiness) error
	GET(echo.Context) error
	POST(echo.Context) error
	MULTIPOST(echo.Context) error
	PUT(echo.Context) error
	DELETE(echo.Context) error
	GETBYID(echo.Context) error
}

type APIRouter interface {
	Connect(string, APIHandler, *echo.Echo, AuthHandler) error
}

type AuthBusiness interface {
	Connect(*db.DatabaseConnection) error
	Authenticate(string, string) (error, string, string)
	Authentication(string, string) (interface{}, error)
	Register(string, string) (interface{}, error)
}

type AuthHandler interface {
	Connect(AuthBusiness) error
	Authentication(echo.Context) error
	Register(echo.Context) error
	Authenticate(func(ec echo.Context) error, ...string) func(ec echo.Context) error
}

type AuthRouter interface {
	Connect(string, AuthHandler, *echo.Echo) error
}
