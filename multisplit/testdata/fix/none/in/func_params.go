package multisplit

func fp(p1, p2, p3 int, p4, p5, p6 struct{}, p7, p8, p9 StructT, p10 string) {} // comment

type fpt1 func(p11, p12, p13 int, p14, p15, p16 struct{}, p17, p18, p19 StructT, p20 string) // comment

type fpt2 = func(p21, p22, p23 int, p24, p25, p26 struct{}, p27, p28, p29 StructT, p30 string) // comment

type (
	fpt3 func(p31, p32, p33 int, p34, p35, p36 struct{}, p37, p38, p39 StructT, p40 string)   // comment
	fpt4 = func(p41, p42, p43 int, p44, p45, p46 struct{}, p47, p48, p49 StructT, p50 string) // comment
)

func fpif(fn func(p51, p52, p53 int, p54, p55, p56 struct{}, p57, p58, p59 StructT, p60 string)) {} // comment

type fpi interface {
	fn(p61, p62, p63 int, p64, p65, p66 struct{}, p67, p68, p69 StructT, p70 string) // comment
}

type fps struct {
	fn func(p71, p72, p73 int, p74, p75, p76 struct{}, p77, p78, p79 StructT, p80 string) // comment
}
