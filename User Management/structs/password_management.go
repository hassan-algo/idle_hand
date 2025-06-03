package structs

type PasswordManagements struct {
	MyPasswordManagements []PasswordManagement `json:"password_managements,omitempty"`
}

type PasswordManagement struct {
	PasswordManagementId   string `json:"password_management_id,omitempty"`
	PasswordManagementName string `json:"password_management_name,omitempty"`
	NewPassword            string `json:"new_password,omitempty"`
	ConfirmPassword        string `json:"confirm_password,omitempty"`
	OTP                    string `json:"otp,omitempty"`
	Email                  string `json:"email,omitempty"`
	UserGuid               string `json:"user_guid,omitempty"`
	OldPassword            string `json:"old_password,omitempty"`
}

func (PasswordManagement) TableName() string {
	return "tbl_users"
}
