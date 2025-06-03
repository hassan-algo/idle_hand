package structs

type Catalogs struct {
	MyCatalogs []Catalog `json:"catalogs,omitempty"`
}

type Catalog struct {
	CatalogGUID        string `json:"catalog_guid,omitempty" gorm:"column:catalog_guid"`
	BusinessGUID       string `json:"business_guid,omitempty" gorm:"column:business_guid"`
	CatalogCategory    string `json:"catalog_category,omitempty" gorm:"column:catalog_category"`
	CatalogName        string `json:"catalog_name,omitempty" gorm:"column:catalog_name"`
	CatalogDescription string `json:"catalog_description,omitempty" gorm:"column:catalog_description"`
	CatalogPrice       string `json:"catalog_price,omitempty" gorm:"column:catalog_price"`
	CatalogOffering    string `json:"catalog_offering,omitempty" gorm:"column:catalog_offering"`
	CatalogPhoto       string `json:"catalog_photo,omitempty" gorm:"column:catalog_photo"`
	AssignedStaffGUID  string `json:"assigned_staff_guid,omitempty" gorm:"column:assigned_staff_guid"`
}

// TableName sets the custom table name for GORM
func (Catalog) TableName() string {
	return "tbl_catalog"
}
