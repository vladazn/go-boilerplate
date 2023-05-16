package authService

type AuthAccountType int

const (
	AuthAccountTypeGoogle AuthAccountType = 1
	AuthAccountTypeApple  AuthAccountType = 2
)

type AuthParams struct {
	AccountType AuthAccountType
	Token       string
	Username    *string
}

type AuthResponse struct {
	AuthToken string
}
