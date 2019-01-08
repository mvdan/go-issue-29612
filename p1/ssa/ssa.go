package ssa

type T struct{}

func (T) foo() {}

func Unused(v interface{}) {
	v.(interface{ foo() }).foo()
}
