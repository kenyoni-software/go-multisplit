package multisplit

func genericFn[T, U any](arg T) T { // comment
	return arg
}

type S[T, U any] struct { // comment
	field  T
	field2 U
}

func (s S[T, U]) Method[X, Y int](lhs X, rhs Y) U { // comment
	return s.field2
}

type I[T, U any] interface { // comment
	Method(arg T) U
}

type IA[T, U any] = I[T, U] // comment

type SIA[T, U any] S[T, U] // comment
