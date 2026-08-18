package multisplit

func genericFn[T, U any](arg T) T { // want `generic type parameters with multiple identifiers \(T, U\) should be split into individual parameters`
	return arg
}

type S[T, U any] struct { // want `generic type parameters with multiple identifiers \(T, U\) should be split into individual parameters`
	field  T
	field2 U
}

type I[T, U any] interface { // want `generic type parameters with multiple identifiers \(T, U\) should be split into individual parameters`
	Method(arg T) U
}

type IA[T, U any] = I[T, U] // want `generic type parameters with multiple identifiers \(T, U\) should be split into individual parameters`

type SIA[T, U any] S[T, U] // want `generic type parameters with multiple identifiers \(T, U\) should be split into individual parameters`
