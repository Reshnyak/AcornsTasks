package models

const (
	ErrorUserNotFound = "User not found"
)

type UserID string

// User model definition
type User struct {
	ID    UserID `json:"id"`
	Login string `json:"login"`
	Token string `json:"access_token"`
}
