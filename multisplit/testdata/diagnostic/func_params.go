package multisplit

func fp(p1, p2, p3 int, p4, p5, p6 struct{}, p7, p8, p9 StructT, p10 string) {} // want `function parameters with multiple identifiers \(p1, p2, p3\) should be split into individual parameters` `function parameters with multiple identifiers \(p4, p5, p6\) should be split into individual parameters` `function parameters with multiple identifiers \(p7, p8, p9\) should be split into individual parameters`

type fpt1 func(p11, p12, p13 int, p14, p15, p16 struct{}, p17, p18, p19 StructT, p20 string) // want `function parameters with multiple identifiers \(p11, p12, p13\) should be split into individual parameters` `function parameters with multiple identifiers \(p14, p15, p16\) should be split into individual parameters` `function parameters with multiple identifiers \(p17, p18, p19\) should be split into individual parameters`

type fpt2 = func(p21, p22, p23 int, p24, p25, p26 struct{}, p27, p28, p29 StructT, p30 string) // want `function parameters with multiple identifiers \(p21, p22, p23\) should be split into individual parameters` `function parameters with multiple identifiers \(p24, p25, p26\) should be split into individual parameters` `function parameters with multiple identifiers \(p27, p28, p29\) should be split into individual parameters`

type (
	fpt3 func(p31, p32, p33 int, p34, p35, p36 struct{}, p37, p38, p39 StructT, p40 string) // want `function parameters with multiple identifiers \(p31, p32, p33\) should be split into individual parameters` `function parameters with multiple identifiers \(p34, p35, p36\) should be split into individual parameters` `function parameters with multiple identifiers \(p37, p38, p39\) should be split into individual parameters`
	fpt4 = func(p41, p42, p43 int, p44, p45, p46 struct{}, p47, p48, p49 StructT, p50 string) // want `function parameters with multiple identifiers \(p41, p42, p43\) should be split into individual parameters` `function parameters with multiple identifiers \(p44, p45, p46\) should be split into individual parameters` `function parameters with multiple identifiers \(p47, p48, p49\) should be split into individual parameters`
)

func fpif(fn func(p51, p52, p53 int, p54, p55, p56 struct{}, p57, p58, p59 StructT, p60 string)) {} // want `function parameters with multiple identifiers \(p51, p52, p53\) should be split into individual parameters` `function parameters with multiple identifiers \(p54, p55, p56\) should be split into individual parameters` `function parameters with multiple identifiers \(p57, p58, p59\) should be split into individual parameters`

type fpi interface {
	fn(p61, p62, p63 int, p64, p65, p66 struct{}, p67, p68, p69 StructT, p70 string) // want `function parameters with multiple identifiers \(p61, p62, p63\) should be split into individual parameters` `function parameters with multiple identifiers \(p64, p65, p66\) should be split into individual parameters` `function parameters with multiple identifiers \(p67, p68, p69\) should be split into individual parameters`
}

type fps struct {
	fn func(p71, p72, p73 int, p74, p75, p76 struct{}, p77, p78, p79 StructT, p80 string) // want `function parameters with multiple identifiers \(p71, p72, p73\) should be split into individual parameters` `function parameters with multiple identifiers \(p74, p75, p76\) should be split into individual parameters` `function parameters with multiple identifiers \(p77, p78, p79\) should be split into individual parameters`
}
