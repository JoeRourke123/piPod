package model

import "time"

type Playlist struct {
	Id          string        `json:"id"`
	Uri         string        `json:"uri"`
	Name        string        `json:"name"`
	Owner       string        `json:"owner"`
	Tracks      []Track       `json:"tracks,omitempty"`
	Duration    string        `json:"duration"`
	Created     *time.Time    `json:"created"`
	AddedAt     *time.Time    `json:"addedAt"`
	CoverArtUrl string        `json:"coverArtUrl,omitempty"`
	Metadata    AlbumMetadata `json:"metadata,omitempty"`
}
