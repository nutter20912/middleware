package model

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	User    *User  `json:"user"`
}

type Comment struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}
