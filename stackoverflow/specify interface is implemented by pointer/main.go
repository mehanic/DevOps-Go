package main

type UserClaims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}
func (c *UserClaims) GetRegisteredClaims() *jwt.RegisteredClaims {
    return &c.RegisteredClaims
}
UserClaimManager := NewManager[UserClaims](...)
