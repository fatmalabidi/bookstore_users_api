package users

type User struct {
	ID          int64  `json:"id"`
	FirstNAme   string `json:"first_name"`
	LastNAme    string `json:"last_name"`
	Email       string `json:"email"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	DateCreated string `json:"date_created"`
}
