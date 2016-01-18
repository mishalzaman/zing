package validate

import (
	"errors"
	"strings"
)

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

	if len(strings.TrimSpace(value)) == 0 {
		v.err = errors.New("String is empty")
		return false
	}

	return true
}

func (v *Validator) NoSpaces(value string) bool {
	if v.err != nil {
		return false
	}

	if strings.Contains(value, " ") {
		v.err = errors.New("String has spaces")
		return false
	}

	return true
}

func (v *Validator) NotValid() bool {
	return v.err != nil
}

func (v *Validator) Error() string {
	return v.err.Error()
}
