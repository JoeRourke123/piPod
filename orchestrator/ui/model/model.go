package model

type ListViewResponse struct {
	Title          string                 `json:"title"`
	ShowStatus     bool                   `json:"show_status"`
	Items          []ListViewItemResponse `json:"items"`
	Icon           string                 `json:"icon,omitempty"`
	AdditionalInfo []AdditionalInfo       `json:"additional_info,omitempty"`
}

type ListViewItemResponse struct {
	Title           string                 `json:"title"`
	Subtitle        string                 `json:"subtitle,omitempty"`
	Path            string                 `json:"path,omitempty"`
	Icon            string                 `json:"icon,omitempty"`
	Disabled        bool                   `json:"disabled,omitempty"`
	BackgroundImage string                 `json:"background_image,omitempty"`
	Actions         []ListViewItemResponse `json:"actions,omitempty"`
	ActionType      string                 `json:"action_type,omitempty"`
	RequestUrl      string                 `json:"request_url,omitempty"`
	ToastMessage    string                 `json:"toast_message,omitempty"`
}

type AuthResponse struct {
	HasToken    bool   `json:"has_token"`
	AuthUrl     string `json:"auth_url"`
	AccessToken string `json:"access_token,omitempty"`
}

type PlayerRequest struct {
	DeviceId        string `json:"deviceId"`
	Action          string `json:"action"`
	SpotifyUri      string `json:"spotifyUri,omitempty"`
	AlbumID         string `json:"albumId,omitempty"`
	PlaybackContext string `json:"playbackContext,omitempty"`
}

type LoadingViewResponse struct {
	Title string `json:"title"`
}

type AdditionalInfo struct {
	Text string `json:"text"`
	Icon string `json:"icon,omitempty"`
	Bold bool   `json:"bold,omitempty"`
}

type OsUpdates struct {
	IsInternetEnabled bool `json:"is_internet_enabled"`
}

type PlayerResponse struct {
	PlayerState string `json:"playerState"`
	SpotifyUri  string `json:"spotifyUri"`
	TrackName   string `json:"trackName"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumId     string `json:"albumId"`
	ImageUrl    string `json:"imageUrl"`
	Duration    int    `json:"duration"`
	PlayerURL   string `json:"playerUrl"`
}
