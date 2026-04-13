package multisplit

var vpkg1, vpkg2 int // want `variable declaration with multiple identifiers \(vpkg1, vpkg2\) should be split into individual declarations`

var vpkg3, vpkg4 int // want `variable declaration with multiple identifiers \(vpkg3, vpkg4\) should be split into individual declarations`

var vpkg5, vpkg6 struct{} // want `variable declaration with multiple identifiers \(vpkg5, vpkg6\) should be split into individual declarations`

var vpkg7, vpkg8 StructT // want `variable declaration with multiple identifiers \(vpkg7, vpkg8\) should be split into individual declarations`

var (
	vpkgb1, vpkgb2 int // want `variable declaration with multiple identifiers \(vpkgb1, vpkgb2\) should be split into individual declarations`
	vpkgb3, vpkgb4 int // want `variable declaration with multiple identifiers \(vpkgb3, vpkgb4\) should be split into individual declarations`
	vpkgb5, vpkgb6 struct{} // want `variable declaration with multiple identifiers \(vpkgb5, vpkgb6\) should be split into individual declarations`
	vpkgb7, vpkgb8 StructT // want `variable declaration with multiple identifiers \(vpkgb7, vpkgb8\) should be split into individual declarations`
)
