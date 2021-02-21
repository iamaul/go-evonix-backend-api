package utils

type Response struct {
	Code         int         `json:"code,omitempty"`
	Result       interface{} `json:"result,omitempty"`
	AccessToken  string      `json:"access_token,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	Message      string      `json:"message,omitempty"`
	Success      bool        `json:"success,omitempty"`
}
