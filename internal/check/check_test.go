package check

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	err := errors.New("This is an Error")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function should be panic")
		}
	}()

	Err(err)
}

func TestErrWithMsg(t *testing.T) {
	err := errors.New("This is an Error")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function should be panic")
		}
	}()

	ErrWithMsg(err, "Unexpected Error")
}

func TestCompareErr(t *testing.T) {
	errInvalid := errors.New("This is an Error")
	errValid := errors.New("bucket already exists")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("function should be panic")
		}
	}()

	CompareErr(errInvalid)

	CompareErr(errValid)
}
