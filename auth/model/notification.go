package model

// Notification struct
type Notification struct {
	User    User   `json:"user"`
	Message string `json:"message"`
	JWT     string `json:"jwt"`
	KEY     string `json:"key"`
}
