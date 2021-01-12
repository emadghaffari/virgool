package model

// Notification struct
type Notification struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	KEY     string `json:"key"`
	Type    string `json:"type"`
}

// Data struct
type Data struct {
	User
	Token string `json:"token"`
}
