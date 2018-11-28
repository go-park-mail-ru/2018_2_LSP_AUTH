package handlers

import "github.com/thedevsaddam/govalidator"

var authValidationRules = govalidator.MapData{
	"email":    []string{"required", "between:4,25", "email"},
	"password": []string{"required", "alpha_space"},
}
