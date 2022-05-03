package domain

type User struct {
	ID                     uint16 `json:"id" db:"id" redis:"id" validate:"omitempty"`
	Name                   string `json:"name" db:"name" redis:"name" validate:"required,lte=6"`
	Password               string `json:"password,omitempty" db:"password" redis:"password" validate:"omitempty,required,gte=6"`
	Email                  string `json:"email" db:"email" redis:"email" validate:"omitempty,required,email"`
	EmailVerified          uint8  `json:"email_verified" db:"email_verified" redis:"email_verified"`
	RegisteredDate         int32  `json:"registered_date" db:"registered_date" redis:"registered_date" validate:"omitempty"`
	RegisterIP             string `json:"register_ip" db:"register_ip" redis:"register_ip" validate:"omitempty"`
	UcpLoginIP             string `json:"ucp_login_ip" db:"ucp_login_ip" redis:"ucp_login_ip"`
	LoginIP                string `json:"login_ip" db:"login_ip" redis:"login_ip"`
	Admin                  uint8  `json:"admin" db:"admin" redis:"admin"`
	AdminDivision          uint8  `json:"admin_division" db:"admin_division" redis:"admin_division"`
	Helper                 uint8  `json:"helper" db:"helper" redis:"helper"`
	LastLogin              int32  `json:"lastlogin" db:"lastlogin" redis:"lastlogin"`
	EvoPoints              uint16 `json:"evo_points" db:"evo_points" redis:"evo_points"`
	EvoLevel               uint8  `json:"evo_level" db:"evo_level" redis:"evo_level"`
	EvoActive              int32  `json:"evo_active" db:"evo_active" redis:"evo_active"`
	EvoNameChange          uint8  `json:"evo_namechange" db:"evo_namechange" redis:"evo_namechange"`
	EvoNumberChange        uint8  `json:"evo_numberchange" db:"evo_numberchange" redis:"evo_numberchange"`
	EvoPlateChange         uint8  `json:"evo_platechange" db:"evo_platechange" redis:"evo_platechange"`
	EvoUCPChange           uint8  `json:"evo_ucpchange" db:"evo_ucpchange" redis:"evo_ucpchange"`
	Status                 uint8  `json:"status" db:"status" redis:"status"`
	DelayCharacterDeletion int32  `json:"delay_character_deletion" db:"delay_character_deletion" redis:"delay_character_deletion"`
	PasswordForgot         uint8  `json:"password_forgot,omitempty" db:"password_forgot" redis:"password_forgot"`
}

type UserWithToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
