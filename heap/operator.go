/*
*@Time        : 2022/01/11
*@File        : operator.go
*@Description :
 */

package heap

type Operator func(res int) bool

func GreaterThan(res int) bool {
	return res > 0
}

func LessThan(res int) bool {
	return res < 0
}
