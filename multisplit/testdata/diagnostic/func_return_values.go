package multisplit

func ft1() (r1, r2, r3 int, r4, r5, r6 struct{}, r7, r8, r9 StructT, r10 string) { // want `function return values with multiple identifiers \(r1, r2, r3\) should be split into individual return values` `function return values with multiple identifiers \(r4, r5, r6\) should be split into individual return values` `function return values with multiple identifiers \(r7, r8, r9\) should be split into individual return values`
	return 0, 0, 0, struct{}{}, struct{}{}, struct{}{}, StructT{}, StructT{}, StructT{}, ""
}

type ftt1 func() (r11, r12, r13 int, r14, r15, r16 struct{}, r17, r18, r19 StructT, r20 string) // want `function return values with multiple identifiers \(r11, r12, r13\) should be split into individual return values` `function return values with multiple identifiers \(r14, r15, r16\) should be split into individual return values` `function return values with multiple identifiers \(r17, r18, r19\) should be split into individual return values`

type ftt2 = func() (r21, r22, r23 int, r24, r25, r26 struct{}, r27, r28, r29 StructT, r30 string) // want `function return values with multiple identifiers \(r21, r22, r23\) should be split into individual return values` `function return values with multiple identifiers \(r24, r25, r26\) should be split into individual return values` `function return values with multiple identifiers \(r27, r28, r29\) should be split into individual return values`

type (
	ftt3 func() (r31, r32, r33 int, r34, r35, r36 struct{}, r37, r38, r39 StructT, r40 string) // want `function return values with multiple identifiers \(r31, r32, r33\) should be split into individual return values` `function return values with multiple identifiers \(r34, r35, r36\) should be split into individual return values` `function return values with multiple identifiers \(r37, r38, r39\) should be split into individual return values`
	ftt4 = func() (r41, r42, r43 int, r44, r45, r46 struct{}, r47, r48, r49 StructT, r50 string) // want `function return values with multiple identifiers \(r41, r42, r43\) should be split into individual return values` `function return values with multiple identifiers \(r44, r45, r46\) should be split into individual return values` `function return values with multiple identifiers \(r47, r48, r49\) should be split into individual return values`
)

func ftif() func() (r51, r52, r53 int, r54, r55, r56 struct{}, r57, r58, r59 StructT, r60 string) { // want `function return values with multiple identifiers \(r51, r52, r53\) should be split into individual return values` `function return values with multiple identifiers \(r54, r55, r56\) should be split into individual return values` `function return values with multiple identifiers \(r57, r58, r59\) should be split into individual return values`
	return func() (r61, r62, r63 int, r64, r65, r66 struct{}, r67, r68, r69 StructT, r70 string) { // want `function return values with multiple identifiers \(r61, r62, r63\) should be split into individual return values` `function return values with multiple identifiers \(r64, r65, r66\) should be split into individual return values` `function return values with multiple identifiers \(r67, r68, r69\) should be split into individual return values`
		return 0, 0, 0, struct{}{}, struct{}{}, struct{}{}, StructT{}, StructT{}, StructT{}, ""
	}
}

type fti interface {
	fn() (r71, r72, r73 int, r74, r75, r76 struct{}, r77, r78, r79 StructT, r80 string) // want `function return values with multiple identifiers \(r71, r72, r73\) should be split into individual return values` `function return values with multiple identifiers \(r74, r75, r76\) should be split into individual return values` `function return values with multiple identifiers \(r77, r78, r79\) should be split into individual return values`
}

type fts struct {
	fn func() (r81, r82, r83 int, r84, r85, r86 struct{}, r87, r88, r89 StructT, r90 string) // want `function return values with multiple identifiers \(r81, r82, r83\) should be split into individual return values` `function return values with multiple identifiers \(r84, r85, r86\) should be split into individual return values` `function return values with multiple identifiers \(r87, r88, r89\) should be split into individual return values`
}
