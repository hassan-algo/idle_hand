package structs

import "time"

type MyAuth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Auths struct {
	MyAuths []Auth `json:"auths"`
}

type Auth struct {
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	UserGuid   string `json:"user_guid,omitempty"`
	LoginToken string `json:"login_token,omitempty"`
}

type Authenticate struct {
	Token string `json:"token,omitempty"`
}
type Response struct {
	Valid   bool        `json:"valid"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseUserWithToken struct {
	// UserGuid   string `json:"userguid,omitempty"`
	Name       string `json:"name,omitempty"`
	ProfilePic string `json:"profilepic,omitempty"`
	Email      string `json:"email,omitempty"`
	UserPlan   string `json:"userplan,omitempty"`
	Token      string `json:"token,omitempty"`
}

type Credentials struct {
	UserGuid    string    `json:"userguid,omitempty"`
	FullName    string    `json:"fullname,omitempty"`
	ProfilePic  string    `json:"profilepic,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	UserType    string    `json:"usertype,omitempty"`
	UserPlan    string    `json:"userplan,omitempty"`
	Login_Token string    `json:"login_token,omitempty"`
	Last_login  time.Time `json:"last_login,omitempty"`
}
