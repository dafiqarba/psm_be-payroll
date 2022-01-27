package dto

//Response model representation for user detail
type UserDetailRes struct {
	User_id     int    `json:"user_id"`
	Name        string `json:"name"`
	Position_id string `json:"position"`
	Nik         string `json:"nik"`
	Role_id     string `json:"role_id"`
}

type UserLoginRes struct {
	User_id  int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"-"`
	Position string `json:"-"`
	Nik      string `json:"-"`
	Role_id  string `json:"role_id"`
}
