package multisplit

func assign() {
	var (
		v1 int
		v2 int
		v3 struct{}
		v4 struct{}
		v5 StructT
	)

	v1, v2, v3, v5 = 1, value(), struct{}{}, StructT{} // want `assignment with multiple targets \(v1, v2, v3, v5\) should be split into individual assignments`

	v3, v4 = struct{}{}, struct{}{} // want `assignment with multiple targets \(v3, v4\) should be split into individual assignments`

	v1, v2 = 1, 2 // want `assignment with multiple targets \(v1, v2\) should be split into individual assignments`
	v1, _ = 1, 2 // want `assignment with multiple targets \(v1, _\) should be split into individual assignments`

	_ = v1
	_ = v2
	_ = v3
	_ = v4
	_ = v5
}
