package validate

import "errors"

type Validator struct {
	err error
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) NotEmptyString(value string) bool {
	if v.err != nil {
		return false
	}

	if len(value) == 0 {
		v.err = errors.New("String is empty")
		return false
	}

	return true
}

func (v *Validator) Valid() bool {
	return v.err != nil
}

func (v *Validator) Error() string {
	return v.err.Error()
}
