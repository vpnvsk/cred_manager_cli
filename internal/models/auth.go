package models

type User struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type JWTToken struct {
	Token string `json:"token"`
}
