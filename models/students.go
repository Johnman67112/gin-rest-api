package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

func StudentValidator(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
