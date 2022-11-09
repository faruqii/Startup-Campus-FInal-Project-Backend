package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UserRegister struct {
	Name     string `json:"name" validate:"required string"`
	Email    string `json:"email" validate:"required email"`
	Phone    string `json:"phone" validate:"required string"`
	Password string `json:"password" validate:"required string"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required email"`
	Password string `json:"password" validate:"required string"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Type  string `json:"type"`
}

type UserToken struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	Token string `json:"token"`
}
