package models

type User struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name" validate:"required,min=3,max=20,alphanum"`
	Password               string `json:"password" validate:"required,min=6"`
	Email                  string `json:"email" validate:"required,email"`
	EmailVerified          int8   `json:"email_verified"`
	RegisteredDate         int64  `json:"registered_date"`
	RegisterIP             string `json:"register_ip"`
	UCPLoginIP             string `json:"ucp_login_ip"`
	LoginIP                string `json:"login_ip"`
	Admin                  int8   `json:"admin"`
	AdminDivision          int8   `json:"admin_division"`
	Helper                 int8   `json:"helper"`
	LastLogin              int64  `json:"last_login"`
	Status                 int8   `json:"status"`
	DelayCharacterDeletion int64  `json:"delay_character_deletion"`
	Blocked                int8   `json:"blocked"`
	LastBlockIssuer        string `json:"lastblock_issuer"`
	LastBlockReason        string `json:"lastblock_reason"`
}
