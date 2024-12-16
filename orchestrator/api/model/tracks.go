package model

type Track struct {
	Id     string `json:"id"`
	Uri    string `json:"uri"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  Album  `json:"album"`
}
