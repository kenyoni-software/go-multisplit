package multisplit

func shortVar() {
	s1 := 1
	s2 := 2

	s3, s4 := 3, 4 // comment

	s5, s6 := value2()

	s7 := struct{}{}
	s8 := struct{}{}

	s9 := 5
	_ = 6

	_ = s1
	_ = s2
	_ = s3
	_ = s4
	_ = s5
	_ = s6
	_ = s7
	_ = s8
	_ = s9
}