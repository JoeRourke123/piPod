package model

type ListViewResponse struct {
	Title        string                 `json:"title"`
	ShowStatus   bool                   `json:"show_status"`
	Items        []ListViewItemResponse `json:"items"`
	FallbackIcon string                 `json:"fallback_icon,omitempty"`
}

type ListViewItemResponse struct {
	Title        string                 `json:"title"`
	Subtitle     string                 `json:"subtitle,omitempty"`
	Path         string                 `json:"path,omitempty"`
	Icon         string                 `json:"icon,omitempty"`
	Actions      []ListViewItemResponse `json:"actions,omitempty"`
	ActionType   string                 `json:"action_type,omitempty"`
	RequestUrl   string                 `json:"request_url,omitempty"`
	ToastMessage string                 `json:"toast_message,omitempty"`
}

type AuthResponse struct {
	HasToken    bool   `json:"has_token"`
	AuthUrl     string `json:"auth_url"`
	AccessToken string `json:"access_token,omitempty"`
}

type PlayerRequest struct {
	DeviceId        string `json:"device_id"`
	Action          string `json:"action"`
	SpotifyUri      string `json:"spotify_uri,omitempty"`
	PlaybackContext string `json:"playback_context,omitempty"`
}

type LoadingViewResponse struct {
	Title string `json:"title"`
}
