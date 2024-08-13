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
