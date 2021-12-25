package skiplist

import (
	"math/rand"
	"time"
)

type Option struct {
	maxLevel      int
	cmp           CmpFunc
	random        *rand.Rand
	isPrintDetail bool
}

type OptionFunc func(*Option)

func NewOption(opts ...OptionFunc) *Option {
	option := &Option{
		maxLevel:      0,
		cmp:           nil,
		isPrintDetail: false,
		random:        rand.New(rand.NewSource(time.Now().Unix())),
	}
	for _, o := range opts {
		o(option)
	}
	return option
}

func WithCompareFunc(cmp CmpFunc) OptionFunc {
	return func(o *Option) {
		o.cmp = cmp
	}
}

func WithMaxLevel(maxLevel int) OptionFunc {
	return func(o *Option) {
		o.maxLevel = maxLevel
	}
}

func WithPrintDetail() OptionFunc {
	return func(o *Option) {
		o.isPrintDetail = true
	}
}
