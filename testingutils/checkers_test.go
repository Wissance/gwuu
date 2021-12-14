package testingutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckStringsSuccessfulWithoutOrder(t *testing.T) {
	arr1 := []string{"", ""}
	arr2 := []string{"", ""}
	checkResult, _ := CheckStrings(t, arr1, arr2, false, true)
	assert.True(t, checkResult)

	arr1 = []string{"English", "Русский"}
	arr2 = []string{"English", "Русский"}
	checkResult, _ = CheckStrings(t, arr1, arr2, false, true)
	assert.True(t, checkResult)

	arr1 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 = []string{"TextFore", "ТекстДва", "TextOne", "TextThree"}
	checkResult, _ = CheckStrings(t, arr1, arr2, false, true)
	assert.True(t, checkResult)
}

func TestCheckStringsFailsWithoutOrder(t *testing.T) {
	arr1 := []string{"", ""}
	arr2 := []string{" ", " "}
	checkResult, msg := CheckStrings(t, arr1, arr2, false, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, msg)

	arr1 = []string{"english", "Русский"}
	arr2 = []string{"English", "русский"}
	checkResult, msg = CheckStrings(t, arr1, arr2, false, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, msg)

	arr1 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 = []string{"TextFore", "ТекстДва", "TextOne"}
	checkResult, msg = CheckStrings(t, arr1, arr2, false, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, msg)
}
