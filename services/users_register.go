package services

func UsersRegister(params *UsersCreateParams) (string, error) {
	user, err := UsersCreateService(params)

	if err != nil {
		return "", err
	}

	return UsersGenerateTokenService(user)
}
