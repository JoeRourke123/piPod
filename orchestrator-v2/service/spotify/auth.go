package spotify

func GetAuthUrl() string {
	return Auth.AuthURL(AuthState)
}
