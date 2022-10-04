package services

type SecretsParams struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name" validate:"required"`
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Website  string `json:"website,omitempty"`
	Note     string `json:"note,omitempty"`
	Pkey     string `json:"pkey,omitempty"`
}
