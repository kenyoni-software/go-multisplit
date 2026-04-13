package multisplit

var (
	vpkgi1 = 1
	vpkgi2 = 2
)

var vpkgi3, vpkgi4 = 3, 4 // comment

var (
	vpkgi5 = struct{}{}
	vpkgi6 = struct{}{}
)

var (
	vpkgi7 = StructT{}
	vpkgi8 = StructT{}
)

var (
	vpkgi9 = 1
	vpkgi10 = value()
)

var vpkgi11, vpkgi12 = value2()

var (
	vpkgig1 = 1
	vpkgig2 = 2
	vpkgig3, vpkgig4   = 3, 4 // comment
	vpkgig5 = struct{}{}
	vpkgig6 = struct{}{}
	vpkgig7 = StructT{}
	vpkgig8 = StructT{}
	vpkgig9 = 1
	vpkgig10 = value()
	vpkgig11, vpkgig12 = value2()
)

var (
	vpkgit1 int = 1
	vpkgit2 int = 2
)

var vpkgit3, vpkgit4 int = 3, 4 // comment

var vpkgit5, vpkgit6 struct{} = struct{}{}, struct{}{}

var (
	vpkgit7 StructT = StructT{}
	vpkgit8 StructT = StructT{}
)

var (
	vpkgit9 int = 1
	vpkgit10 int = value()
)

var vpkgit11, vpkgit12 int = value2()

var (
	vpkgitg1 int = 1
	vpkgitg2 int = 2
	vpkgitg3, vpkgitg4   int      = 3, 4 // comment
	vpkgitg5, vpkgitg6   struct{} = struct{}{}, struct{}{}
	vpkgitg7 StructT = StructT{}
	vpkgitg8 StructT = StructT{}
	vpkgitg9 int = 1
	vpkgitg10 int = value()
	vpkgitg11, vpkgitg12 int      = value2()
)
