package multisplit

func genericFn[T any, U any](arg T) T { // comment
	return arg
}

type S[T any, U any] struct { // comment
	field  T
	field2 U
}

func (s S[T, U]) Method[X int, Y int](lhs X, rhs Y) U { // comment
	return s.field2
}

type I[T any, U any] interface { // comment
	Method(arg T) U
}

type IA[T any, U any] = I[T, U] // comment

type SIA[T any, U any] S[T, U] // comment
