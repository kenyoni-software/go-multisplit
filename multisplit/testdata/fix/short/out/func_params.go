package multisplit

func fp(p1 int, p2 int, p3 int, p4, p5, p6 struct{}, p7 StructT, p8 StructT, p9 StructT, p10 string) {} // comment

type fpt1 func(p11 int, p12 int, p13 int, p14, p15, p16 struct{}, p17 StructT, p18 StructT, p19 StructT, p20 string) // comment

type fpt2 = func(p21 int, p22 int, p23 int, p24, p25, p26 struct{}, p27 StructT, p28 StructT, p29 StructT, p30 string) // comment

type (
	fpt3 func(p31 int, p32 int, p33 int, p34, p35, p36 struct{}, p37 StructT, p38 StructT, p39 StructT, p40 string)   // comment
	fpt4 = func(p41 int, p42 int, p43 int, p44, p45, p46 struct{}, p47 StructT, p48 StructT, p49 StructT, p50 string) // comment
)

func fpif(fn func(p51 int, p52 int, p53 int, p54, p55, p56 struct{}, p57 StructT, p58 StructT, p59 StructT, p60 string)) {} // comment

type fpi interface {
	fn(p61 int, p62 int, p63 int, p64, p65, p66 struct{}, p67 StructT, p68 StructT, p69 StructT, p70 string) // comment
}

type fps struct {
	fn func(p71 int, p72 int, p73 int, p74, p75, p76 struct{}, p77 StructT, p78 StructT, p79 StructT, p80 string) // comment
}
