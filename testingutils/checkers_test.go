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

func TestCheckStringsSuccessfulWithOrder(t *testing.T) {
	arr1 := []string{"", ""}
	arr2 := []string{"", ""}
	checkResult, _ := CheckStrings(t, arr1, arr2, true, true)
	assert.True(t, checkResult)

	arr1 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 = []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
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

func TestCheckStringsFailsWithOrder(t *testing.T) {
	arr1 := []string{"TextOne", "ТекстДва", "TextThree", "TextFore"}
	arr2 := []string{"TextFore", "ТекстДва", "TextOne", "TextThree"}
	checkResult, _ := CheckStrings(t, arr1, arr2, true, false)
	assert.False(t, checkResult)
}

func TestCheckIntegersSuccessfulWithoutOrder(t *testing.T) {
	arr1 := make([]int, 3)
	arr1[0] = 0
	arr1[2] = 0
	arr1[1] = 2

	arr2 := make([]int, 3)
	arr2[0] = 2
	arr2[1] = 0
	arr2[2] = 0

	checkResult, err := CheckIntegers(t, arr1, arr2, false, true)
	assert.True(t, checkResult)
	assert.Empty(t, err)

}

func TestCheckIntegersFailsWithoutOrder(t *testing.T) {
	arr1 := make([]int, 3)
	arr1[0] = 0
	arr1[2] = 0
	arr1[1] = 3

	arr2 := make([]int, 3)
	arr2[0] = 2
	arr2[1] = 0
	arr2[2] = 0

	checkResult, err := CheckIntegers(t, arr1, arr2, false, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, err)

}

func TestCheckIntegersSuccessfulWithOrder(t *testing.T) {
	arr1 := make([]int, 3)
	arr1[0] = 111
	arr1[2] = 333
	arr1[1] = 222

	arr2 := make([]int, 3)
	arr2[0] = 111
	arr2[1] = 222
	arr2[2] = 333

	checkResult, err := CheckIntegers(t, arr1, arr2, true, true)
	assert.True(t, checkResult)
	assert.Empty(t, err)
}

func TestCheckIntegersFailsWithOrder(t *testing.T) {
	arr1 := make([]int, 3)
	arr1[0] = 111
	arr1[2] = 3334
	arr1[1] = 222

	arr2 := make([]int, 3)
	arr2[0] = 111
	arr2[1] = 222
	arr2[2] = 333

	checkResult, err := CheckIntegers(t, arr1, arr2, true, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, err)
}

func TestCheckFloat32SuccessfulWithOrder(t *testing.T) {
    arr1 := make([]float32, 3)
    arr1[0] = 10.55
    arr1[1] = 99
    arr1[2] = 55.9

    arr2 := make([]float32, 3)
    arr2[0] = 11.01
    arr2[1] = 99
    arr2[2] = 55.6

    checkResult, err := CheckFloats(t, arr1, arr2, 0.5, true, true)
    assert.True(t, checkResult)
    assert.Empty(t, err)
}

func TestCheckFloat32FailsWithOrder(t *testing.T) {
	arr1 := make([]float32, 3)
	arr1[0] = 10.55
	arr1[1] = 99
	arr1[2] = 55.9

	arr2 := make([]float32, 3)
	arr2[0] = 11.51
	arr2[1] = 99
	arr2[2] = 55.6

	checkResult, err := CheckFloats(t, arr1, arr2, 0.5, true, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, err)
}

func TestCheckFloat32SuccessfulWithoutOrder(t *testing.T) {
	arr1 := make([]float32, 3)
	arr1[0] = 129.489
	arr1[1] = 1173.1
	arr1[2] = 199.64

	arr2 := make([]float32, 3)
	arr2[0] = 1172.77
	arr2[1] = 199.99
	arr2[2] = 129.123

	checkResult, err := CheckFloats(t, arr1, arr2, 0.5, false, true)
	assert.True(t, checkResult)
	assert.Empty(t, err)
}

func TestCheckFloat32FailsWithoutOrder(t *testing.T) {
	arr1 := make([]float32, 3)
	arr1[0] = 129.489
	arr1[1] = 1173.1
	arr1[2] = 199.64

	arr2 := make([]float32, 3)
	arr2[0] = 1172.77
	arr2[1] = 299.99
	arr2[2] = 129.123

	checkResult, err := CheckFloats(t, arr1, arr2, 0.5, true, false)
	assert.False(t, checkResult)
	assert.NotEmpty(t, err)
}
