package model

import "time"

type AlbumMetadata struct {
	IsDownloaded bool   `json:"isDownloaded"`
	IsInLibrary  bool   `json:"isInLibrary"`
	PlayCount    int    `json:"playCount"`
	LastPlayed   string `json:"lastPlayed"`
	IsPinned     bool   `json:"isPinned"`
}

type TrackMetadata struct {
	IsDownloaded bool      `json:"isDownloaded"`
	IsInLibrary  bool      `json:"isInLibrary"`
	PlayCount    int       `json:"playCount,omitempty"`
	LastPlayed   string    `json:"lastPlayed,omitempty"`
	FileLocation string    `json:"fileLocation,omitempty"`
	DownloadDate time.Time `json:"downloadDate,omitempty"`
}
