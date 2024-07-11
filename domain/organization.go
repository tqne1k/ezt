package domain

const (
	TableOrganization = "user"
)

type Organization struct {
	Id          string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedDate string `gorm:"type:varchar(255);not null" json:"created_date"`
}
