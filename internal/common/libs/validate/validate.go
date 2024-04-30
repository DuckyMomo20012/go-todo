package validate

import "github.com/go-playground/validator/v10"

func Validate(obj any) error {
	validate := validator.New()

	if err := validate.Struct(obj); err != nil {
		return err
	}

	return nil
}
