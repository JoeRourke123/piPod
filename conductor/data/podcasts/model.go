package podcasts

type CastosResponse struct {
	Success bool               `json:"success"`
	Data    []CastosSearchItem `json:"data"`
}

type CastosSearchItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
