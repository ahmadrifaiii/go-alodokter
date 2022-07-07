package model

type AuthAccess struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Claim struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}
