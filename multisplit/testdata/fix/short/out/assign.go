package multisplit

type ash struct {
	v1 int
	v2 int
}

func (_ *ash) value() int {
	return 0
}

func assign() {
	var (
		v1 int
		v2 int
		v3 struct{}
		v4 struct{}
		v5 StructT
	)

	as := ash{}
	as.v1 = 42
	as.v2 = as.value()

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
