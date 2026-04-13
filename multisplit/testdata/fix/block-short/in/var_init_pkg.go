package multisplit

var vpkgi1, vpkgi2 = 1, 2

var vpkgi3, vpkgi4 = 3, 4 // comment

var vpkgi5, vpkgi6 = struct{}{}, struct{}{}

var vpkgi7, vpkgi8 = StructT{}, StructT{}

var vpkgi9, vpkgi10 = 1, value()

var vpkgi11, vpkgi12 = value2()

var (
	vpkgig1, vpkgig2   = 1, 2
	vpkgig3, vpkgig4   = 3, 4 // comment
	vpkgig5, vpkgig6   = struct{}{}, struct{}{}
	vpkgig7, vpkgig8   = StructT{}, StructT{}
	vpkgig9, vpkgig10  = 1, value()
	vpkgig11, vpkgig12 = value2()
)

var vpkgit1, vpkgit2 int = 1, 2

var vpkgit3, vpkgit4 int = 3, 4 // comment

var vpkgit5, vpkgit6 struct{} = struct{}{}, struct{}{}

var vpkgit7, vpkgit8 StructT = StructT{}, StructT{}

var vpkgit9, vpkgit10 int = 1, value()

var vpkgit11, vpkgit12 int = value2()

var (
	vpkgitg1, vpkgitg2   int      = 1, 2
	vpkgitg3, vpkgitg4   int      = 3, 4 // comment
	vpkgitg5, vpkgitg6   struct{} = struct{}{}, struct{}{}
	vpkgitg7, vpkgitg8   StructT  = StructT{}, StructT{}
	vpkgitg9, vpkgitg10  int      = 1, value()
	vpkgitg11, vpkgitg12 int      = value2()
)
