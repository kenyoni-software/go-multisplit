package multisplit

func cFunc() {
	const cfunc1, cfunc2 = 1, 2 // want `const declaration with multiple identifiers \(cfunc1, cfunc2\) should be split into individual declarations`

	const cfunc3, cfunc4 = 3, 4 // want `const declaration with multiple identifiers \(cfunc3, cfunc4\) should be split into individual declarations`

	const (
		cfuncg1, cfuncg2 = 1, 2 // want `const declaration with multiple identifiers \(cfuncg1, cfuncg2\) should be split into individual declarations`
		cfuncg3, cfuncg4 = 3, 4 // want `const declaration with multiple identifiers \(cfuncg3, cfuncg4\) should be split into individual declarations`
	)

	const cfunct1, cfunct2 int = 1, 2 // want `const declaration with multiple identifiers \(cfunct1, cfunct2\) should be split into individual declarations`

	const cfunct3, cfunct4 int = 3, 4 // want `const declaration with multiple identifiers \(cfunct3, cfunct4\) should be split into individual declarations`

	const (
		cfunctg1, cfunctg2 int = 1, 2 // want `const declaration with multiple identifiers \(cfunctg1, cfunctg2\) should be split into individual declarations`
		cfunctg3, cfunctg4 int = 3, 4 // want `const declaration with multiple identifiers \(cfunctg3, cfunctg4\) should be split into individual declarations`
	)
}
