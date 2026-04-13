package multisplit

func vd() {
	var (
		vfunc1 int
		vfunc2 int
	)

	var vfunc3, vfunc4 int // comment

	var vfunc5, vfunc6 struct{}

	var (
		vfunc7 StructT
		vfunc8 StructT
	)

	var (
		vfuncb1 int
		vfuncb2 int
		vfuncb3, vfuncb4 int // comment
		vfuncb5, vfuncb6 struct{}
		vfuncb7 StructT
		vfuncb8 StructT
	)

	_ = vfunc1
	_ = vfunc2
	_ = vfunc3
	_ = vfunc4
	_ = vfunc5
	_ = vfunc6
	_ = vfunc7
	_ = vfunc8
	_ = vfuncb1
	_ = vfuncb2
	_ = vfuncb3
	_ = vfuncb4
	_ = vfuncb5
	_ = vfuncb6
	_ = vfuncb7
	_ = vfuncb8
}
