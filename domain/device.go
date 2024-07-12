package domain

const (
	TableDevice = "device"
)

type Device struct {
	Id     string `gorm:"type:uuId;primary_key;default:uuId_generate_v4();auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
	UserId string `gorm:"type:uuId;not null" json:"user_id"`

	PrivateKey      string `gorm:"type:varchar(255);not null" json:"private_key"`
	PublicKey       string `gorm:"type:varchar(255);not null" json:"public_key"`
	Endpoint        string `json:"endpoint"`
	AllowedIPs      string `json:"allowed_ips"`
	LatestHandshake string `json:"latest_handshake"`
	Transfer        string `json:"transfer"`
}
