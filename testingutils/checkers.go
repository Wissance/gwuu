package testingutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"testing"
)

/*
 *  This function allow to check two arrays of Strings with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of strings
 *     - actual - another array of strings
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 */
func CheckStrings(t *testing.T, expected []string, actual []string, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
    assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, len(actual))
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
			    if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object {0} exists in actual array", eItem))
		}
	}
}

/*
 *  This function allow to check two arrays of int's with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of int
 *     - actual - another array of int
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 */
func CheckIntegers(t *testing.T, expected []int, actual []int, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, len(actual))
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object {0} exists in actual array", eItem))
		}
	}
}

/*
 *  This function allow to check two arrays of int64 with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of int64
 *     - actual - another array of int64
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 */
func CheckIntegers64(t *testing.T, expected []int64, actual []int64, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, len(actual))
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object {0} exists in actual array", eItem))
		}
	}
}

func CheckUnsignedIntegers(t *testing.T, expected []uint, actual []uint, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, len(actual))
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object {0} exists in actual array", eItem))
		}
	}
}

func CheckUnsignedIntegers64(t *testing.T, expected []uint64, actual []uint64, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, len(actual))
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object {0} exists in actual array", eItem))
		}
	}
}

func CheckFloats(t *testing.T, expected []float32, actual []float32, precision float32, checkOrder bool) {

}

func CheckFloats64(t *testing.T, expected []float64, actual []float64, precision float64, checkOrder bool) {

}

func CheckComplexes(t *testing.T, expected []complex64, actual []complex64, precision float64, checkOrder bool) {

}

func CheckComplexes128(t *testing.T, expected []complex64, actual []complex64, precision float64, checkOrder bool) {

}

// todo: waiting for go 1.18
/*func checkComparable[T any](t *testing.T, expected []T, actual []T, checkOrder bool) {

}*/

func contains(arr []int, item int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}