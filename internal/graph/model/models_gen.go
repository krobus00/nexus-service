// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type Register struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
