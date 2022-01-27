package entity

//User represents users table in the database
type User struct {
	User_id     int    `json:"user_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Position_id string `json:"position"`
	Nik         string `json:"nik"`
	Role_id     string `json:"role_id"`
}

type UserDetailModel struct {
	User_id       int    `json:"user_id"`
	Name          string `json:"name"`
	Position_id   int    `json:"position_id"`
	Nik           string `json:"nik"`
	Role_id       int    `json:"role_id"`
	Role_name     string `json:"role"`
	Position_name string `json:"position"`
}
