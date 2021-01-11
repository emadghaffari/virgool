package model

// Notification struct
type Notification struct {
	Data    User   `json:"data"`
	Message string `json:"message"`
	Token   string `json:"token"`
	KEY     string `json:"key"`
}
