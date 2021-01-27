package model

import (
	"github.com/go-playground/validator/v10"
)

var (
	// Validator var
	Validator validate = &valid{}
)

type validate interface {
	New()
	Get() *validator.Validate
}

// valid struct
type valid struct {
	validate *validator.Validate
}

// new validator
func (v *valid) New() {
	v.validate = validator.New()
}

// Get validator
func (v *valid) Get() *validator.Validate {
	return v.validate
}
