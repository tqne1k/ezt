package domain

const (
	TableDepartmentNetwork = "department_network"
)

type DepartmentNetwork struct {
	DepartmentId string `json:"department_id"`
	NetworkId    string `json:"network_id"`
	CreatedDate  string `gorm:"type:varchar(255);not null" json:"created_date"`
}
