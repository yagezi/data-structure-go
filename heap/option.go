/*
*@Time        : 2022/01/11
*@File        : option.go
*@Description :
 */
package heap

type Option struct {
	cap      int
	cmp      CmpFunc
	operator Operator
}

type OptionFunc func(*Option)

func NewOption(opts []OptionFunc) *Option {
	option := &Option{
		cap:      1,
		cmp:      CmpInt,
		operator: LessThan,
	}
	for _, o := range opts {
		o(option)
	}
	return option
}

func WithCapacity(cap int) OptionFunc {
	return func(o *Option) {
		o.cap = cap
	}
}

func WithCmpFunc(f CmpFunc) OptionFunc {
	return func(o *Option) {
		o.cmp = f
	}
}

func MinHeap() OptionFunc {
	return func(o *Option) {
		o.operator = LessThan
	}
}

func MaxHeap() OptionFunc {
	return func(o *Option) {
		o.operator = GreaterThan
	}
}
