package users

type LoginRequestCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken    string `json:"access_token"`
	ExpirationTime string `json:"expiration_time"`
}
