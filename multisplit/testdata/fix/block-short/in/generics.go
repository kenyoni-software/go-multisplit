package multisplit

func genericFn[T, U any](arg T) T { // comment
	return arg
}

type S[T, U any] struct { // comment
	field  T
	field2 U
}

type I[T, U any] interface { // comment
	Method(arg T) U
}

type IA[T, U any] = I[T, U] // comment

type SIA[T, U any] S[T, U] // comment
