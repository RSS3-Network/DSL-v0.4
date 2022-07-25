package validatorx

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var _ echo.Validator = &Validator{}

var Default = &Validator{
	validate: validator.New(),
}

type Validator struct {
	validate *validator.Validate
	locker   sync.Mutex
}

func (v *Validator) Validate(i interface{}) error {
	if v.validate == nil {
		v.locker.Lock()
		defer v.locker.Unlock()

		if v.validate == nil {
			v.validate = validator.New()
		}
	}

	return v.validate.Struct(i)
}
