package model

import "time"

type Track struct {
	Id          string        `json:"id"`
	Uri         string        `json:"uri"`
	Name        string        `json:"name"`
	Artist      string        `json:"artist"`
	Album       Album         `json:"album"`
	Duration    int           `json:"duration"`
	Metadata    TrackMetadata `json:"metadata"`
	ReleaseDate *time.Time    `json:"release_date,omitempty"`
	ISRC        string        `json:"isrc"`
}
