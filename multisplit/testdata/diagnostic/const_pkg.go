package multisplit

const cpkgi1, cpkgi2 = 1, 2 // want `const declaration with multiple identifiers \(cpkgi1, cpkgi2\) should be split into individual declarations`

const cpkgi3, cpkgi4 = 3, 4 // want `const declaration with multiple identifiers \(cpkgi3, cpkgi4\) should be split into individual declarations`

const (
	cpkgig1, cpkgig2 = 1, 2 // want `const declaration with multiple identifiers \(cpkgig1, cpkgig2\) should be split into individual declarations`
	cpkgig3, cpkgig4 = 3, 4 // want `const declaration with multiple identifiers \(cpkgig3, cpkgig4\) should be split into individual declarations`
)

const cpkgit1, cpkgit2 int = 1, 2 // want `const declaration with multiple identifiers \(cpkgit1, cpkgit2\) should be split into individual declarations`

const cpkgit3, cpkgit4 int = 3, 4 // want `const declaration with multiple identifiers \(cpkgit3, cpkgit4\) should be split into individual declarations`

const (
	cpkgitg1, cpkgitg2 int = 1, 2 // want `const declaration with multiple identifiers \(cpkgitg1, cpkgitg2\) should be split into individual declarations`
	cpkgitg3, cpkgitg4 int = 3, 4 // want `const declaration with multiple identifiers \(cpkgitg3, cpkgitg4\) should be split into individual declarations`
)
