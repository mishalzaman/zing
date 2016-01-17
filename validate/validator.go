package validate

import "errors"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) NotEmptyString(value string) error {
	if len(value) > 0 {
		return nil
	}

	return errors.New("String is empty")
}
