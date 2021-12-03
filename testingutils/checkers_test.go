package testingutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckStringsSuccessfulWithoutOrder(t *testing.T) {
	arr1 := []string{"", ""}
	arr2 := []string{"", ""}
	checkResult := CheckStrings(t, arr1, arr2, false)
	assert.True(t, checkResult)

	arr1 = []string{"English", "Русский"}
	arr2 = []string{"English", "Русский"}
	checkResult = CheckStrings(t, arr1, arr2, false)
	assert.True(t, checkResult)

	arr1 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 = []string{"TextFore", "ТекстДва", "TextOne", "TextThree"}
	checkResult = CheckStrings(t, arr1, arr2, false)
	assert.True(t, checkResult)
}

func TestCheckStringsFailsWithoutOrder(t *testing.T) {
	arr1 := []string{"", ""}
	arr2 := []string{" ", " "}
	checkResult := CheckStrings(t, arr1, arr2, false)
	assert.False(t, checkResult)

	arr1 = []string{"english", "Русский"}
	arr2 = []string{"English", "русский"}
	checkResult = CheckStrings(t, arr1, arr2, false)
	assert.True(t, checkResult)

	arr1 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 = []string{"TextFore", "ТекстДва", "TextOne"}
	checkResult = CheckStrings(t, arr1, arr2, false)
	assert.True(t, checkResult)
}
