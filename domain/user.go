package domain

const (
	TableUser = "user"
)

type User struct {
	Id          string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedDate string `gorm:"type:varchar(255);not null" json:"created_date"`

	DepartmentId string `json:"department_id"`

	// TODO: Temporary username and password for testing
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}
