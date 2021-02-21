package models

// User model schema
type User struct {
	ID int64 `json:"id"`
	// Name                   string `json:"name" validate:"required,min=3,max=20,alphanum"`
	Name string `json:"name"`
	// Password               string `json:"password" validate:"required,min=6"`
	Password string `json:"password,omitempty"`
	// Email                  string `json:"email" validate:"required,email"`
	Email                  string `json:"email"`
	EmailVerified          uint8  `json:"email_verified"`
	RegisteredDate         int64  `json:"registered_date"`
	RegisterIP             string `json:"register_ip"`
	UCPLoginIP             string `json:"ucp_login_ip"`
	LoginIP                string `json:"login_ip"`
	Admin                  uint8  `json:"admin"`
	AdminDivision          uint8  `json:"admin_division"`
	Helper                 uint8  `json:"helper"`
	LastLogin              int64  `json:"last_login"`
	Status                 uint8  `json:"status"`
	DelayCharacterDeletion int64  `json:"delay_character_deletion"`
	Blocked                uint8  `json:"blocked"`
	LastBlockIssuer        string `json:"lastblock_issuer"`
	LastBlockReason        string `json:"lastblock_reason"`
}
