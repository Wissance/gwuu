package testingutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"testing"
)

// CheckStrings
/*  This function allow to compare two arrays of Strings with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of strings
 *     - actual - another array of strings
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns true if there is no assert fail, otherwise - false
 */
func CheckStrings(t *testing.T, expected []string, actual []string, checkOrder bool) bool {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
    assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				// if j (index) already contains
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
		}
	}
	return !t.Failed()
}

// CheckIntegers
/*  This function allow to compare two arrays of int's with order and without it
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
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
		}
	}
}

// CheckIntegers64
/*  This function allow to check two arrays of int64 with order and without it
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
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
		}
	}
}

// CheckUnsignedIntegers
/*  This function allow to compare two arrays of uint with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of uint
 *     - actual - another array of uint
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 */
func CheckUnsignedIntegers(t *testing.T, expected []uint, actual []uint, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
		}
	}
}

// CheckUnsignedIntegers64
/*  This function allow to compare two arrays of uint64 with order and without it
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of uint64
 *     - actual - another array of uint64
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *  Functions return nothing and asserts if arrays are not equals
 */
func CheckUnsignedIntegers64(t *testing.T, expected []uint64, actual []uint64, checkOrder bool) {
	if expected == nil || actual == nil{
		assert.Nil(t, expected, "Checking that expected is nil")
		assert.Nil(t, actual, "Checking that actual is nil")
	}
	assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			assert.Equal(t, expected[i], actual[i])
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				if !contains(usedObjects, j) {
					if eItem == aItem {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
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

// contains
/* Function check that int item in arr
 * Parameters:
 *     - arr - array where item could be located
 *     - item - what we are searching in arr
 * Returns true if item in array otherwise false
 */
func contains(arr []int, item int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
