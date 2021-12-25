package skiplist

import (
	"strings"
)

type CmpFunc func(a, b interface{}) int

func CmpString(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}
