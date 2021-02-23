package model

// Query struct
type Query struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}