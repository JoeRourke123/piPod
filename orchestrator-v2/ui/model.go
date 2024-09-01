package ui

type ListViewResponse struct {
	Title        string                 `json:"title"`
	ShowStatus   bool                   `json:"show_status"`
	Items        []ListViewItemResponse `json:"items"`
	FallbackIcon string                 `json:"fallback_icon,omitempty"`
}

type ListViewItemResponse struct {
	Title string `json:"title"`
	Path  string `json:"path,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

type AuthResponse struct {
	HasToken    bool   `json:"has_token"`
	AuthUrl     string `json:"auth_url"`
	AccessToken string `json:"access_token,omitempty"`
}

type PlayerRequest struct {
	DeviceId   string `json:"device_id"`
	Action     string `json:"action"`
	SpotifyUri string `json:"spotify_uri,omitempty"`
}
