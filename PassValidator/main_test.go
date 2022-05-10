package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ShouldGetErr_WhenPassLenIsLessThanEight(t *testing.T) {
	pass := "134567"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, TOO_SHORT, err.Error())
}

func TestValidate_ShouldGetErr_WhenNotContainsAtLeastTwoNumbers(t *testing.T) {
	pass := "abcdefgh1"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, TOO_FEW_DIGITS, err.Error())
}

func TestValidate_ShouldGetErr_WhenIsLessThanEightAndNotContainsAtLeastTwoNumbers(t *testing.T) {
	pass := "ab2cd"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, fmt.Sprintf("%v\n%v", TOO_SHORT, TOO_FEW_DIGITS), err.Error())
}