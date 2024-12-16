package model

type PlayerResponse struct {
	PlayerState string `json:"playerState"`
	SpotifyUri  string `json:"spotifyUri"`
	TrackName   string `json:"trackName"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumId     string `json:"albumId"`
	ImageUrl    string `json:"imageUrl"`
	Duration    int    `json:"duration"`
	PlayerUrl   string `json:"playerUrl"`
}

type PlayerRequest struct {
	DeviceId        string `json:"deviceId"`
	Action          string `json:"action"`
	SpotifyUri      string `json:"spotifyUri,omitempty"`
	AlbumId         string `json:"albumId,omitempty"`
	PlaybackContext string `json:"playbackContext,omitempty"`
}
