package model

type ListView struct {
	Title          string         `json:"title"`
	ShowStatus     bool           `json:"show_status"`
	Items          []ListViewItem `json:"items"`
	Icon           string         `json:"icon,omitempty"`
	AdditionalInfo []ListViewInfo `json:"additional_info,omitempty"`
}

type ListViewItem struct {
	Title           string         `json:"title"`
	Subtitle        string         `json:"subtitle,omitempty"`
	Path            string         `json:"path,omitempty"`
	Icon            string         `json:"icon,omitempty"`
	Disabled        bool           `json:"disabled,omitempty"`
	BackgroundImage string         `json:"background_image,omitempty"`
	Actions         []ListViewItem `json:"actions,omitempty"`
	ActionType      string         `json:"action_type,omitempty"`
	RequestUrl      string         `json:"request_url,omitempty"`
	ToastMessage    string         `json:"toast_message,omitempty"`
}

type ListViewInfo struct {
	Text string `json:"text"`
	Icon string `json:"icon,omitempty"`
	Bold bool   `json:"bold,omitempty"`
}
