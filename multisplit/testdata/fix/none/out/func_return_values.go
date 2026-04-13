package multisplit

func ft1() (r1 int, r2 int, r3 int, r4, r5, r6 struct{}, r7 StructT, r8 StructT, r9 StructT, r10 string) { // comment
	return 0, 0, 0, struct{}{}, struct{}{}, struct{}{}, StructT{}, StructT{}, StructT{}, ""
}

type ftt1 func() (r11 int, r12 int, r13 int, r14, r15, r16 struct{}, r17 StructT, r18 StructT, r19 StructT, r20 string) // comment

type ftt2 = func() (r21 int, r22 int, r23 int, r24, r25, r26 struct{}, r27 StructT, r28 StructT, r29 StructT, r30 string) // comment

type (
	ftt3 func() (r31 int, r32 int, r33 int, r34, r35, r36 struct{}, r37 StructT, r38 StructT, r39 StructT, r40 string) // comment
	ftt4 = func() (r41 int, r42 int, r43 int, r44, r45, r46 struct{}, r47 StructT, r48 StructT, r49 StructT, r50 string) // comment
)

func ftif() func() (r51 int, r52 int, r53 int, r54, r55, r56 struct{}, r57 StructT, r58 StructT, r59 StructT, r60 string) { // comment
	return func() (r61 int, r62 int, r63 int, r64, r65, r66 struct{}, r67 StructT, r68 StructT, r69 StructT, r70 string) { // comment
		return 0, 0, 0, struct{}{}, struct{}{}, struct{}{}, StructT{}, StructT{}, StructT{}, ""
	}
}

type fti interface {
	fn () (r71 int, r72 int, r73 int, r74, r75, r76 struct{}, r77 StructT, r78 StructT, r79 StructT, r80 string) // comment
}

type fts struct {
	fn func() (r81 int, r82 int, r83 int, r84, r85, r86 struct{}, r87 StructT, r88 StructT, r89 StructT, r90 string) // comment
}
