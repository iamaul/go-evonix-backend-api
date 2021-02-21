package utils

// ErrorMsg is self-explanatory
type ErrorMsg string

const (
	// TokenErr will be shown when authorizing user token whether its exist or is expired
	TokenErr = "Token not found or has expired, authorization denied."
	// UsernameExists will be shown when the user create a new account
	UsernameExists = "Username already exists"
	// EmailExists will be shown when the user create a new account
	EmailExists = "Email already exists"
	// ContentTypeErr shown when requesting API with invalid format content type
	ContentTypeErr = "Unable to make a request due to policy"
	// InvalidSigningMethodErr will be shown when the signing method is invalid
	InvalidSigningMethodErr = "Invalid signing method"
	// InvalidRefreshTokenErr will be shown when the token claims is invalid
	InvalidRefreshTokenErr = "Invalid refresh token"
)
