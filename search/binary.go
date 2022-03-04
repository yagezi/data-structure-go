/*
*@Time        : 2022/01/17
*@File        : binary.go
*@Description : binary search
 */

package search

// BinarySearchRecur return the position of target in sorted list
// if not found, return the inserting position
func BinarySearchRecur(list []int, target int) int {
	return biSearch(list, target, 0, len(list))
}

// biSearch implements binary search by recursion, [l, r)
// return the position or inserting position of target in sorted list
func biSearch(list []int, target, l, r int) int {
	if l >= r {
		return l
	}
	mid := l + (r-l)/2
	if target == list[mid] {
		return mid
	} else if target < list[mid] {
		return biSearch(list, target, l, mid)
	} else {
		return biSearch(list, target, mid+1, r)
	}
}

// BinarySearchIter implements binary search by iteration, [l, r)
// return the position or inserting position of target in sorted list
func BinarySearchIter(list []int, target int) int {
	l := 0
	r := len(list)
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if target == list[mid] {
			return mid
		} else if target < list[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// Sqrt implements integer sqrt() math function
// by binary search method
func Sqrt(x int) int {
	ans := x
	l := 0
	r := x
	var mid int
	for l <= r {
		// avoid overflow
		mid = l + (r-l)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	return ans
}
