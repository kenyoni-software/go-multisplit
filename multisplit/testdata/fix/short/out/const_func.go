package multisplit

func cFunc() {
	const cfunc1 = 1
	const cfunc2 = 2

	const cfunc3, cfunc4 = 3, 4 // comment

	const (
		cfuncg1 = 1
		cfuncg2 = 2
		cfuncg3, cfuncg4 = 3, 4 // comment
	)

	const cfunct1 int = 1
	const cfunct2 int = 2

	const cfunct3, cfunct4 int = 3, 4 // comment

	const (
		cfunctg1 int = 1
		cfunctg2 int = 2
		cfunctg3, cfunctg4 int = 3, 4 // comment
	)
}
