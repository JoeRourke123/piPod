package model

type Album struct {
	Id          string  `json:"id"`
	Uri         string  `json:"uri"`
	Name        string  `json:"name"`
	Artist      string  `json:"artist"`
	Tracks      []Track `json:"tracks,omitempty"`
	ReleaseDate string  `json:"releaseDate"`
	Duration    int     `json:"duration"`
}
