package multisplit

func assign() {
	var (
		v1 int
		v2 int
		v3 struct{}
		v4 struct{}
		v5 StructT
	)

	v1, v2, v3, v5 = 1, value(), struct{}{}, StructT{}

	v3, v4 = struct{}{}, struct{}{}

	v1, v2 = 1, 2 // comment
	v1, _ = 1, 2

	_ = v1
	_ = v2
	_ = v3
	_ = v4
	_ = v5
}
