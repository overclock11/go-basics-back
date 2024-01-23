package models

type Secret struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Jwtsing  string `json:"jwtsing"`
	Database string `json:"database"`
}
