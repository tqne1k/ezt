package domain

const (
	TableDepartment = "department"
)

type Department struct {
	Id          string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedDate string `gorm:"type:varchar(255);not null" json:"created_date"`

	OrganizationId string `json:"organization_id"`
	DepartmentId   string `json:"department_id"`

	NetworkId string `json:"network_id"`
}
