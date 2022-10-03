package services

type UsersSignParams struct {
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,gte=5,lte=33"`
}

type UsersCreateParams struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32,uniq"`
	Password string `json:"password" validate:"required,gte=5,lte=33"`
}
