package structs

type BusinessClients struct {
	MyBusinessClients []BusinessClient `json:"business_clients,omitempty"`
}

type BusinessClient struct {
	ClientGUID             string `json:"client_guid,omitempty" gorm:"column:client_guid"`
	BusinessGUID           string `json:"business_guid,omitempty" gorm:"column:business_guid"`
	PreferredContactMethod string `json:"preferred_contact_method,omitempty" gorm:"column:prefferred_contact_method"`
	Source                 string `json:"source,omitempty" gorm:"column:source"`
	Note                   string `json:"note,omitempty" gorm:"column:note"`
	BusinessClientGUID     string `json:"business_client_guid,omitempty" gorm:"column:business_client_guid"`
}

func (BusinessClient) TableName() string {
	return "tbl_business_client"
}
