package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ErrValidation struct {
	errFields map[string]string
	msg       string
}

func NewSingleValidationErr(fieldName, msg string) error {
	return &ErrValidation{
		errFields: map[string]string{
			fieldName: msg,
		},
		msg: msg,
	}
}

type Validation interface {
	ValidateStruct(structPtr interface{}, fields ...*validation.FieldRules) error
}

func (v *ErrValidation) Error() string {
	return v.msg
}

func (v *ErrValidation) GetErrFields() map[string]string {
	return v.errFields
}

func ValidateStruct(structPtr interface{}, fields ...*validation.FieldRules) error {
	err := validation.ValidateStruct(structPtr, fields...)

	if e, ok := err.(validation.InternalError); ok {
		// an internal error happened
		panic(e.InternalError())
	} else if errs, ok := err.(validation.Errors); ok {
		fieldErrs := map[string]string{}
		for fieldName, fieldErr := range errs {
			fieldErrs[fieldName] = fieldErr.Error()
		}

		return &ErrValidation{
			msg:       err.Error(),
			errFields: fieldErrs,
		}
	}

	return err
}
