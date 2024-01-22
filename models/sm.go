package models

type Secret struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Jwtrsing string `json:"jwtrsing"`
	Database string `json:"database"`
}
