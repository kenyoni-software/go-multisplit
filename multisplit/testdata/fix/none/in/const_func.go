package multisplit

func cFunc() {
	const cfunc1, cfunc2 = 1, 2

	const cfunc3, cfunc4 = 3, 4 // comment

	const (
		cfuncg1, cfuncg2 = 1, 2
		cfuncg3, cfuncg4 = 3, 4 // comment
	)

	const cfunct1, cfunct2 int = 1, 2

	const cfunct3, cfunct4 int = 3, 4 // comment

	const (
		cfunctg1, cfunctg2 int = 1, 2
		cfunctg3, cfunctg4 int = 3, 4 // comment
	)
}
