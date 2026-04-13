package multisplit

func vd() {
	var vfunc1, vfunc2 int // want `variable declaration with multiple identifiers \(vfunc1, vfunc2\) should be split into individual declarations`

	var vfunc3, vfunc4 int // want `variable declaration with multiple identifiers \(vfunc3, vfunc4\) should be split into individual declarations`

	var vfunc5, vfunc6 struct{} // want `variable declaration with multiple identifiers \(vfunc5, vfunc6\) should be split into individual declarations`

	var vfunc7, vfunc8 StructT // want `variable declaration with multiple identifiers \(vfunc7, vfunc8\) should be split into individual declarations`

	var (
		vfuncb1, vfuncb2 int // want `variable declaration with multiple identifiers \(vfuncb1, vfuncb2\) should be split into individual declarations`
		vfuncb3, vfuncb4 int // want `variable declaration with multiple identifiers \(vfuncb3, vfuncb4\) should be split into individual declarations`
		vfuncb5, vfuncb6 struct{} // want `variable declaration with multiple identifiers \(vfuncb5, vfuncb6\) should be split into individual declarations`
		vfuncb7, vfuncb8 StructT // want `variable declaration with multiple identifiers \(vfuncb7, vfuncb8\) should be split into individual declarations`
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
