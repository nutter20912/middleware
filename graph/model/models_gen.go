// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
