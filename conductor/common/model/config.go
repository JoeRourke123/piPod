package model

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	AuthUrl     string `json:"authUrl"`
	HasToken    bool   `json:"hasToken"`
}
