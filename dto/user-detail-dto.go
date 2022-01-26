package dto

type User struct {
	User_id  int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Nik      string `json:"nik"`
	Role_id  string `json:"role_id"`
}