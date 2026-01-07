package Entity

import (
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Student struct {
	gorm.Model 
	Name  string `valid:"required~Name is required"`
	Price float64 `valid:"required~Price must be greater than 0,range(1|999999)~Price must be greater than 0"`
	Stock int `valid:"range(0|999999)~Stock cannot be negative"`
	Description string `valid:"required~Description must be at least 10 characters,stringlength(10|100)~Description must be at least 10 characters"`
}

func (s *Student) Validation() error {
	ok, err := govalidator.ValidateStruct(s)

	if !ok {
	return err
	}
	return nil
}