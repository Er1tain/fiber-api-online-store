package client_models

type Client struct {
	Email    string `gorm:"primary_key" json:"email"`
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
