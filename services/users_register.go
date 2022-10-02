package services

func UsersRegister(params *UsersCreateParams) (string, error) {
	user, err := UsersCreateService(params)

	if err != nil {
		panic(err)
	}

	return UsersGenerateTokenService(user)
}
