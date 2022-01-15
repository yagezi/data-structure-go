/*
*@Time        : 2022/01/12
*@File        : option.go
*@Description :
 */

package graph

type Option struct {
	hasVertexValue bool
}

type OptionFunc func(*Option)

func NewOption(opts []OptionFunc) *Option {
	option := &Option{
		hasVertexValue: false,
	}
	for _, o := range opts {
		o(option)
	}
	return option
}

func WithVertexData() OptionFunc {
	return func(o *Option) {
		o.hasVertexValue = true
	}
}
