package domain

const (
	TableNetwork = "network"
)

type Network struct {
	Id          string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedDate string `gorm:"type:varchar(255);not null" json:"created_date"`

	PublicKey      string `json:"public_key"`
	PrivateKey     string `json:"private_key"`
	NetworkAddress string `json:"network_address"`
	ListeningPort  string `json:"listening_port"`
}

type NetworkResponse struct {
	Id          string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreatedDate string `gorm:"type:varchar(255);not null" json:"created_date"`

	PublicKey      string `json:"public_key"`
	PrivateKey     string `json:"private_key"`
	NetworkAddress string `json:"network_address"`
	ListeningPort  string `json:"listening_port"`

	Devices []Device `json:"devices"`
}
