package services

type UsersSignParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersCreateParams struct {
	UsersSignParams
	Name string `json:"name"`
}
