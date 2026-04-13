package multisplit

func assign() {
	var (
		v1 int
		v2 int
		v3 struct{}
		v4 struct{}
		v5 StructT
	)

	v1 = 1
	v2 = value()
	v3 = struct{}{}
	v5 = StructT{}

	v3 = struct{}{}
	v4 = struct{}{}

	v1, v2 = 1, 2 // comment
	v1 = 1
	_ = 2

	_ = v1
	_ = v2
	_ = v3
	_ = v4
	_ = v5
}
