package multisplit

func var_func_init() {
	var vpkgi1, vpkgi2 = 1, 2 // want `variable declaration with multiple identifiers and initializers \(vpkgi1, vpkgi2\) should be split into individual declarations`

	var vpkgi3, vpkgi4 = 3, 4 // want `variable declaration with multiple identifiers and initializers \(vpkgi3, vpkgi4\) should be split into individual declarations`

	var vpkgi5, vpkgi6 = struct{}{}, struct{}{} // want `variable declaration with multiple identifiers and initializers \(vpkgi5, vpkgi6\) should be split into individual declarations`

	var vpkgi7, vpkgi8 = StructT{}, StructT{} // want `variable declaration with multiple identifiers and initializers \(vpkgi7, vpkgi8\) should be split into individual declarations`

	var vpkgi9, vpkgi10 = 1, value() // want `variable declaration with multiple identifiers and initializers \(vpkgi9, vpkgi10\) should be split into individual declarations`

	var vpkgi11, vpkgi12 = value2()

	var (
		vpkgig1, vpkgig2   = 1, 2 // want `variable declaration with multiple identifiers and initializers \(vpkgig1, vpkgig2\) should be split into individual declarations`
		vpkgig3, vpkgig4   = 3, 4 // want `variable declaration with multiple identifiers and initializers \(vpkgig3, vpkgig4\) should be split into individual declarations`
		vpkgig5, vpkgig6   = struct{}{}, struct{}{} // want `variable declaration with multiple identifiers and initializers \(vpkgig5, vpkgig6\) should be split into individual declarations`
		vpkgig7, vpkgig8   = StructT{}, StructT{} // want `variable declaration with multiple identifiers and initializers \(vpkgig7, vpkgig8\) should be split into individual declarations`
		vpkgig9, vpkgig10  = 1, value() // want `variable declaration with multiple identifiers and initializers \(vpkgig9, vpkgig10\) should be split into individual declarations`
		vpkgig11, vpkgig12 = value2()
	)

	var vpkgit1, vpkgit2 int = 1, 2 // want `variable declaration with multiple identifiers and initializers \(vpkgit1, vpkgit2\) should be split into individual declarations`

	var vpkgit3, vpkgit4 int = 3, 4 // want `variable declaration with multiple identifiers and initializers \(vpkgit3, vpkgit4\) should be split into individual declarations`

	var vpkgit5, vpkgit6 struct{} = struct{}{}, struct{}{} // want `variable declaration with multiple identifiers and initializers \(vpkgit5, vpkgit6\) should be split into individual declarations`

	var vpkgit7, vpkgit8 StructT = StructT{}, StructT{} // want `variable declaration with multiple identifiers and initializers \(vpkgit7, vpkgit8\) should be split into individual declarations`

	var vpkgit9, vpkgit10 int = 1, value() // want `variable declaration with multiple identifiers and initializers \(vpkgit9, vpkgit10\) should be split into individual declarations`

	var vpkgit11, vpkgit12 int = value2()

	var (
		vpkgitg1, vpkgitg2   int      = 1, 2 // want `variable declaration with multiple identifiers and initializers \(vpkgitg1, vpkgitg2\) should be split into individual declarations`
		vpkgitg3, vpkgitg4   int      = 3, 4 // want `variable declaration with multiple identifiers and initializers \(vpkgitg3, vpkgitg4\) should be split into individual declarations`
		vpkgitg5, vpkgitg6   struct{} = struct{}{}, struct{}{} // want `variable declaration with multiple identifiers and initializers \(vpkgitg5, vpkgitg6\) should be split into individual declarations`
		vpkgitg7, vpkgitg8   StructT  = StructT{}, StructT{} // want `variable declaration with multiple identifiers and initializers \(vpkgitg7, vpkgitg8\) should be split into individual declarations`
		vpkgitg9, vpkgitg10  int      = 1, value() // want `variable declaration with multiple identifiers and initializers \(vpkgitg9, vpkgitg10\) should be split into individual declarations`
		vpkgitg11, vpkgitg12 int      = value2()
	)

	_ = vpkgi1
	_ = vpkgi2
	_ = vpkgi3
	_ = vpkgi4
	_ = vpkgi5
	_ = vpkgi6
	_ = vpkgi7
	_ = vpkgi8
	_ = vpkgi9
	_ = vpkgi10
	_ = vpkgi11
	_ = vpkgi12
	_ = vpkgig1
	_ = vpkgig2
	_ = vpkgig3
	_ = vpkgig4
	_ = vpkgig5
	_ = vpkgig6
	_ = vpkgig7
	_ = vpkgig8
	_ = vpkgig9
	_ = vpkgig10
	_ = vpkgig11
	_ = vpkgig12
	_ = vpkgit1
	_ = vpkgit2
	_ = vpkgit3
	_ = vpkgit4
	_ = vpkgit5
	_ = vpkgit6
	_ = vpkgit7
	_ = vpkgit8
	_ = vpkgit9
	_ = vpkgit10
	_ = vpkgit11
	_ = vpkgit12
	_ = vpkgitg1
	_ = vpkgitg2
	_ = vpkgitg3
	_ = vpkgitg4
	_ = vpkgitg5
	_ = vpkgitg6
	_ = vpkgitg7
	_ = vpkgitg8
	_ = vpkgitg9
	_ = vpkgitg10
	_ = vpkgitg11
	_ = vpkgitg12
}
