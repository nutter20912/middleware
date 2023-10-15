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

type DepositOrder struct {
	ID        string               `json:"id"`
	UserID    string               `json:"user_id"`
	Status    DepositStatus        `json:"status"`
	Amount    float64              `json:"amount"`
	Memo      string               `json:"memo"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	Events    []*DepositOrderEvent `json:"events"`
}
