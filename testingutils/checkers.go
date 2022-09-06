package testingutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"math"
	"testing"
)

const (
	ExpectedOrActualBothNotNil = "One of arrays (expected, actual) is nil other is not nil"
	ArraysLengthAreNotSame     = "Arrays length are not same"
	ItemNotFound               = "Expected array item: \"{0}\" at index: \"{1}\" was not found in actual array"
)

type Numeric interface {
	int | int32 | uint32 | int64 | uint64
}

// CheckStrings
/*  This function allow us to compare two arrays of Strings with order and without it and with asserts (assertErr is true) and without order
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of strings
 *     - actual - another array of strings
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckStrings(t *testing.T, expected []string, actual []string, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			itemCheck = itemCheck && (expected[i] == actual[i])
			if assertErr {
				assert.Equal(t, expected[i], actual[i])
			}
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
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
}

// CheckNumeric
/*  This function allow us to compare two arrays of fixed point values with order and without it
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of int
 *     - actual - another array of int
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckNumeric[T Numeric](t *testing.T, expected []T, actual []T, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			itemCheck = itemCheck && (expected[i] == actual[i])
			if assertErr {
				assert.Equal(t, expected[i], actual[i])
			}
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
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
}

// CheckFloats
/*  This function allow us to compare two arrays of float32 with order and without it with tolerance
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of float32
 *     - actual - another array of float32
 *     - tolerance - tolerance between array items, used as math.abs(expected - actual) < tolerance
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckFloats(t *testing.T, expected []float32, actual []float32, tolerance float64, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			comparisonResult := math.Abs(float64(expected[i]-actual[i])) < tolerance
			itemCheck = itemCheck && comparisonResult
			if assertErr {
				assert.True(t, comparisonResult)
			}
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				// if j (index) already contains
				if !contains(usedObjects, j) {
					comparisonResult := math.Abs(float64(eItem-aItem)) < tolerance
					if comparisonResult {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
}

// CheckFloats64
/*  This function allow us to compare two arrays of float64 with order and without it with tolerance
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of float64
 *     - actual - another array of float64
 *     - tolerance - tolerance between array items, used as math.abs(expected - actual) < tolerance
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckFloats64(t *testing.T, expected []float64, actual []float64, tolerance float64, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			comparisonResult := math.Abs(expected[i]-actual[i]) < tolerance
			itemCheck = itemCheck && comparisonResult
			if assertErr {
				assert.True(t, comparisonResult)
			}
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				// if j (index) already contains
				if !contains(usedObjects, j) {
					comparisonResult := math.Abs(eItem-aItem) < tolerance
					if comparisonResult {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
}

// CheckComplexes
/*  This function allow us to compare two arrays of complex64 with order and without it with tolerance applying to real and imaginary part independently
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of complex64
 *     - actual - another array of complex64
 *     - tolerance - tolerance between array items, used as math.abs(real(expected) - real(actual)) < tolerance &&
 *                                                          math.abs(imag(expected) - imag(actual)) < tolerance
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckComplexes(t *testing.T, expected []complex64, actual []complex64, tolerance float64, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			comparisonResult := math.Abs(float64(real(expected[i])-real(actual[i]))) < tolerance &&
				math.Abs(float64(imag(expected[i])-imag(actual[i]))) < tolerance
			itemCheck = itemCheck && comparisonResult
			if assertErr {
				assert.True(t, comparisonResult)
			}
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				// if j (index) already contains
				if !contains(usedObjects, j) {
					comparisonResult := math.Abs(float64(real(eItem)-real(aItem))) < tolerance &&
						math.Abs(float64(imag(eItem)-imag(aItem))) < tolerance
					if comparisonResult {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
}

// CheckComplexes128
/*  This function allow us to compare two arrays of complex128 with order and without it with tolerance applying to real and imaginary part independently
*   This function could work without assert (see assertErr param)
 *  Parameters:
 *     - t is Test State, because we provide functions to check equality in tests
 *     - expected - one array of complex128
 *     - actual - another array of complex128
 *     - tolerance - tolerance between array items, used as math.abs(real(expected) - real(actual)) < tolerance &&
 *                                                          math.abs(imag(expected) - imag(actual)) < tolerance
 *     - checkOrder - parameter that is responsible for check data with respect to order of arrays items
 *     - assertErr - assert if True otherwise just use a result to check whether error occurred or not
 *  Functions return nothing and asserts if arrays are not equals
 *  Returns (true, empty str) if there is no assert fail, otherwise - false + reason
*/
func CheckComplexes128(t *testing.T, expected []complex128, actual []complex128, tolerance float64, checkOrder bool, assertErr bool) (bool, string) {
	if expected == nil || actual == nil {
		nilArraysCheck := expected == nil && actual == nil
		if assertErr {
			assert.Nil(t, expected, "Checking that expected is nil")
			assert.Nil(t, actual, "Checking that actual is nil")
		}
		if nilArraysCheck {
			return true, ""
		}
		return false, ExpectedOrActualBothNotNil
	}
	if len(expected) != len(actual) {
		if assertErr {
			assert.Equal(t, len(expected), len(actual), "Checking that arrays length are equals")
		}
		return false, ArraysLengthAreNotSame
	}

	itemCheck := true
	// created empty array that will be populated if we found items
	usedObjects := make([]int, 0)
	for i, eItem := range expected {
		if checkOrder {
			comparisonResult := math.Abs(real(expected[i])-real(actual[i])) < tolerance &&
				math.Abs(imag(expected[i])-imag(actual[i])) < tolerance
			itemCheck = itemCheck && comparisonResult
			if assertErr {
				assert.True(t, comparisonResult)
			}
		} else {
			unOrderedCheck := false
			for j, aItem := range actual {
				// if j (index) already contains
				if !contains(usedObjects, j) {
					comparisonResult := math.Abs(real(eItem)-real(aItem)) < tolerance &&
						math.Abs(imag(eItem)-imag(aItem)) < tolerance
					if comparisonResult {
						unOrderedCheck = true
						usedObjects = append(usedObjects, j)
						break
					}
				}
			}
			itemCheck = itemCheck && unOrderedCheck
			if assertErr {
				assert.True(t, unOrderedCheck, stringFormatter.Format("Checking object \"{0}\" exists in actual array", eItem))
			}
		}
		if !itemCheck {
			return false, stringFormatter.Format(ItemNotFound, eItem, i)
		}
	}
	return true, ""
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
