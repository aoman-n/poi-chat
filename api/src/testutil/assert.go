package testutil

import (
	"errors"
	"testing"

	"github.com/laster18/poi/api/src/util/validator"
	"github.com/stretchr/testify/assert"
)

func AssertValidationErr(t *testing.T, err error, fieldName, expectErrMsg string) {
	var validationErr *validator.ErrValidation
	if !errors.As(err, &validationErr) {
		t.Errorf("expected = validationErr, got = %T\n", err)
	}
	fieldErrs := validationErr.GetErrFields()
	actualErrMsg := fieldErrs[fieldName]
	assert.Equal(t, expectErrMsg, actualErrMsg)
}

func AssertNoValidationErr(t *testing.T, err error, fieldName string) {
	t.Helper()

	if err == nil {
		return
	}

	var validationErr *validator.ErrValidation
	if !errors.As(err, &validationErr) {
		t.Errorf("expected = validationErr, got = %T\n", err)
	}
	fieldErrs := validationErr.GetErrFields()

	assert.Equal(t, "", fieldErrs[fieldName])
}
