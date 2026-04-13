package multisplit

func shortVar() {
	s1, s2 := 1, 2 // want `short variable declaration with multiple identifiers \(s1, s2\) should be split into individual declarations`

	s3, s4 := 3, 4 // want `short variable declaration with multiple identifiers \(s3, s4\) should be split into individual declarations`

	s5, s6 := value2()

	s7, s8 := struct{}{}, struct{}{} // want `short variable declaration with multiple identifiers \(s7, s8\) should be split into individual declarations`

	s9, _ := 5, 6 // want `short variable declaration with multiple identifiers \(s9, _\) should be split into individual declarations`

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
