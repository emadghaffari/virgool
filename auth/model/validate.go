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

type valid struct {
	validate *validator.Validate
}

func (v *valid) New() {
	v.validate = validator.New()
}

func (v *valid) Get() *validator.Validate {
	return v.validate
}
