//nolint:dupl,funlen,goconst,revive
package multisplit_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/kenyoni-software/go-multisplit/multisplit"
)

func TestAssign(t *testing.T) {
	t.Parallel()
	run(t, multisplit.Settings{
		Assign: true,
	}, []testResult{
		{
			message: "assignment with multiple targets (v1, v2, v3, v5) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split into individual assignments",
				newText: []string{
					"v1 = 1\n\tv2 = value()\n\tv3 = struct{}{}\n\tv5 = StructT{}",
				},
			}},
		},
		{
			message: "assignment with multiple targets (v3, v4) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split into individual assignments",
				newText: []string{
					"v3 = struct{}{}\n\tv4 = struct{}{}",
				},
			}},
		},
		{
			message: "assignment with multiple targets (v1, v2) should be split into individual assignments",
		},
		{
			message: "assignment with multiple targets (v1, _) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split into individual assignments",
				newText: []string{
					"v1 = 1\n\t_ = 2",
				},
			}},
		},
	})
}

func TestConstFunc(t *testing.T) {
	t.Parallel()
	t.Run("no block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			ConstDeclFunc:        true,
			ConstDeclFuncToBlock: false,
		}, []testResult{
			{
				message: "const declaration with multiple identifiers (cfunc1, cfunc2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const cfunc1 = 1\n\tconst cfunc2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunc3, cfunc4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfuncg1, cfuncg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cfuncg1 = 1\n\t\tcfuncg2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfuncg3, cfuncg4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfunct1, cfunct2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const cfunct1 int = 1\n\tconst cfunct2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunct3, cfunct4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfunctg1, cfunctg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cfunctg1 int = 1\n\t\tcfunctg2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunctg3, cfunctg4) should be split into individual declarations",
			},
		})
	})

	t.Run("block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			ConstDeclFunc:        true,
			ConstDeclFuncToBlock: true,
		}, []testResult{
			{
				message: "const declaration with multiple identifiers (cfunc1, cfunc2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const (\n\t\tcfunc1 = 1\n\t\tcfunc2 = 2\n\t)",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunc3, cfunc4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfuncg1, cfuncg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cfuncg1 = 1\n\t\tcfuncg2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfuncg3, cfuncg4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfunct1, cfunct2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const (\n\t\tcfunct1 int = 1\n\t\tcfunct2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunct3, cfunct4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cfunctg1, cfunctg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cfunctg1 int = 1\n\t\tcfunctg2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cfunctg3, cfunctg4) should be split into individual declarations",
			},
		})
	})
}

func TestConstPkg(t *testing.T) {
	t.Parallel()
	t.Run("no block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			ConstDeclPkg:        true,
			ConstDeclPkgToBlock: false,
		}, []testResult{
			{
				message: "const declaration with multiple identifiers (cpkgi1, cpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const cpkgi1 = 1\nconst cpkgi2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgi3, cpkgi4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgig1, cpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cpkgig1 = 1\n\tcpkgig2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgig3, cpkgig4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgit1, cpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const cpkgit1 int = 1\nconst cpkgit2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgit3, cpkgit4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgitg1, cpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cpkgitg1 int = 1\n\tcpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgitg3, cpkgitg4) should be split into individual declarations",
			},
		})
	})

	t.Run("block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			ConstDeclPkg:        true,
			ConstDeclPkgToBlock: true,
		}, []testResult{
			{
				message: "const declaration with multiple identifiers (cpkgi1, cpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const (\n\tcpkgi1 = 1\n\tcpkgi2 = 2\n)",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgi3, cpkgi4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgig1, cpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cpkgig1 = 1\n\tcpkgig2 = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgig3, cpkgig4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgit1, cpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"const (\n\tcpkgit1 int = 1\n\tcpkgit2 int = 2\n)",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgit3, cpkgit4) should be split into individual declarations",
			},
			{
				message: "const declaration with multiple identifiers (cpkgitg1, cpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual const declarations",
					newText: []string{
						"cpkgitg1 int = 1\n\tcpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "const declaration with multiple identifiers (cpkgitg3, cpkgitg4) should be split into individual declarations",
			},
		})
	})
}

func TestFuncParams(t *testing.T) {
	t.Parallel()
	run(t, multisplit.Settings{
		FuncParams: true,
	}, []testResult{
		{
			message: "function parameters with multiple identifiers (p1, p2, p3) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p1 int, p2 int, p3 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p4, p5, p6) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p7, p8, p9) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p7 StructT, p8 StructT, p9 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p11, p12, p13) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p11 int, p12 int, p13 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p14, p15, p16) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p17, p18, p19) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p17 StructT, p18 StructT, p19 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p21, p22, p23) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p21 int, p22 int, p23 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p24, p25, p26) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p27, p28, p29) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p27 StructT, p28 StructT, p29 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p31, p32, p33) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p31 int, p32 int, p33 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p34, p35, p36) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p37, p38, p39) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p37 StructT, p38 StructT, p39 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p41, p42, p43) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p41 int, p42 int, p43 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p44, p45, p46) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p47, p48, p49) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p47 StructT, p48 StructT, p49 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p51, p52, p53) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p51 int, p52 int, p53 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p54, p55, p56) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p57, p58, p59) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p57 StructT, p58 StructT, p59 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p61, p62, p63) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p61 int, p62 int, p63 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p64, p65, p66) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p67, p68, p69) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p67 StructT, p68 StructT, p69 StructT",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p71, p72, p73) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p71 int, p72 int, p73 int",
				},
			}},
		},
		{
			message: "function parameters with multiple identifiers (p74, p75, p76) should be split into individual parameters",
		},
		{
			message: "function parameters with multiple identifiers (p77, p78, p79) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split into individual parameters",
				newText: []string{
					"p77 StructT, p78 StructT, p79 StructT",
				},
			}},
		},
	})
}

func TestFuncReturnValues(t *testing.T) {
	t.Parallel()
	run(t, multisplit.Settings{
		FuncReturnValues: true,
	}, []testResult{
		{
			message: "function return values with multiple identifiers (r1, r2, r3) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r1 int, r2 int, r3 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r4, r5, r6) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r7, r8, r9) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r7 StructT, r8 StructT, r9 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r11, r12, r13) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r11 int, r12 int, r13 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r14, r15, r16) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r17, r18, r19) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r17 StructT, r18 StructT, r19 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r21, r22, r23) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r21 int, r22 int, r23 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r24, r25, r26) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r27, r28, r29) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r27 StructT, r28 StructT, r29 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r31, r32, r33) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r31 int, r32 int, r33 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r34, r35, r36) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r37, r38, r39) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r37 StructT, r38 StructT, r39 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r41, r42, r43) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r41 int, r42 int, r43 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r44, r45, r46) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r47, r48, r49) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r47 StructT, r48 StructT, r49 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r51, r52, r53) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r51 int, r52 int, r53 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r54, r55, r56) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r57, r58, r59) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r57 StructT, r58 StructT, r59 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r61, r62, r63) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r61 int, r62 int, r63 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r64, r65, r66) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r67, r68, r69) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r67 StructT, r68 StructT, r69 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r71, r72, r73) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r71 int, r72 int, r73 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r74, r75, r76) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r77, r78, r79) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r77 StructT, r78 StructT, r79 StructT",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r81, r82, r83) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r81 int, r82 int, r83 int",
				},
			}},
		},
		{
			message: "function return values with multiple identifiers (r84, r85, r86) should be split into individual return values",
		},
		{
			message: "function return values with multiple identifiers (r87, r88, r89) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split into individual return values",
				newText: []string{
					"r87 StructT, r88 StructT, r89 StructT",
				},
			}},
		},
	})
}

func TestShortVar(t *testing.T) {
	t.Parallel()
	run(t, multisplit.Settings{
		ShortVarDecl: true,
	}, []testResult{
		{
			message: "short variable declaration with multiple identifiers (s1, s2) should be split into individual declarations",
			suggestedFixes: []analysisFix{{
				message: "split into individual short variable declarations",
				newText: []string{
					"s1 := 1\n\ts2 := 2",
				},
			}},
		},
		{
			message: "short variable declaration with multiple identifiers (s3, s4) should be split into individual declarations",
		},
		{
			message: "short variable declaration with multiple identifiers (s7, s8) should be split into individual declarations",
			suggestedFixes: []analysisFix{{
				message: "split into individual short variable declarations",
				newText: []string{
					"s7 := struct{}{}\n\ts8 := struct{}{}",
				},
			}},
		},
		{
			message: "short variable declaration with multiple identifiers (s9, _) should be split into individual declarations",
			suggestedFixes: []analysisFix{{
				message: "split into individual short variable declarations",
				newText: []string{
					"s9 := 5\n\t_ = 6",
				},
			}},
		},
	})
}

func TestStructFields(t *testing.T) {
	t.Parallel()
	run(t, multisplit.Settings{
		StructFields: true,
	}, []testResult{
		{
			message: "struct field declaration with multiple identifiers (v1, v2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"v1 int\n\tv2 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (v3, v4) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"v3 string `tag:\"value\"`\n\tv4 string `tag:\"value\"`",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (v5, v6) should be split into individual fields",
		},
		{
			message: "struct field declaration with multiple identifiers (v7, v8) should be split into individual fields",
		},
		{
			message: "struct field declaration with multiple identifiers (v9, v10) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"v9 int\n\t\tv10 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (spv1, spv2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"spv1 int; spv2 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (spv2, spv3) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"spv2 int; spv3 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (sv1, sv2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"sv1 int\n\t\tsv2 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (sv3, sv4) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"sv3 string `tag:\"value\"`\n\t\tsv4 string `tag:\"value\"`",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (sv5, sv6) should be split into individual fields",
		},
		{
			message: "struct field declaration with multiple identifiers (sv7, sv8) should be split into individual fields",
		},
		{
			message: "struct field declaration with multiple identifiers (sv9, sv10) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"sv9 int\n\t\t\tsv10 int",
					},
				},
			},
		},
		{
			message: "struct field declaration with multiple identifiers (spv2, spv3) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split into individual struct fields",
					newText: []string{
						"spv2 int\n\t\tspv3 int",
					},
				},
			},
		},
	})
}

func TestVarFunc(t *testing.T) {
	t.Parallel()
	t.Run("no block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclFunc:        true,
			VarDeclFuncToBlock: false,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers (vfunc1, vfunc2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vfunc1 int\n\tvar vfunc2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfunc3, vfunc4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfunc5, vfunc6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfunc7, vfunc8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vfunc7 StructT\n\tvar vfunc8 StructT",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb1, vfuncb2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vfuncb1 int\n\t\tvfuncb2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb3, vfuncb4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb5, vfuncb6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb7, vfuncb8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vfuncb7 StructT\n\t\tvfuncb8 StructT",
					},
				}},
			},
		})
	})

	t.Run("block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclFunc:        true,
			VarDeclFuncToBlock: true,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers (vfunc1, vfunc2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvfunc1 int\n\t\tvfunc2 int\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfunc3, vfunc4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfunc5, vfunc6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfunc7, vfunc8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvfunc7 StructT\n\t\tvfunc8 StructT\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb1, vfuncb2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vfuncb1 int\n\t\tvfuncb2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb3, vfuncb4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb5, vfuncb6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vfuncb7, vfuncb8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vfuncb7 StructT\n\t\tvfuncb8 StructT",
					},
				}},
			},
		})
	})
}

//nolint:maintidx
func TestVarInitFunc(t *testing.T) {
	t.Parallel()
	t.Run("no block-no short", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitFunc:        true,
			VarDeclInitFuncToBlock: false,
			VarDeclInitFuncToShort: false,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi1 = 1\n\tvar vpkgi2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi5 = struct{}{}\n\tvar vpkgi6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi7 = StructT{}\n\tvar vpkgi8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi9 = 1\n\tvar vpkgi10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit1 int = 1\n\tvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\n\tvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit9 int = 1\n\tvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\t\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})

	t.Run("block-no short", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitFunc:        true,
			VarDeclInitFuncToBlock: true,
			VarDeclInitFuncToShort: false,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgi1 = 1\n\t\tvpkgi2 = 2\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgi5 = struct{}{}\n\t\tvpkgi6 = struct{}{}\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgi7 = StructT{}\n\t\tvpkgi8 = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgi9 = 1\n\t\tvpkgi10 = value()\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit1 int = 1\n\t\tvpkgit2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit7 StructT = StructT{}\n\t\tvpkgit8 StructT = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit9 int = 1\n\t\tvpkgit10 int = value()\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\t\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})

	t.Run("no-block-short", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitFunc:        true,
			VarDeclInitFuncToBlock: false,
			VarDeclInitFuncToShort: true,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi1 := 1\n\tvpkgi2 := 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi5 := struct{}{}\n\tvpkgi6 := struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi7 := StructT{}\n\tvpkgi8 := StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi9 := 1\n\tvpkgi10 := value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit1 int = 1\n\tvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\n\tvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit9 int = 1\n\tvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\t\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})

	t.Run("block-short", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitFunc:        true,
			VarDeclInitFuncToBlock: true,
			VarDeclInitFuncToShort: true,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi1 := 1\n\tvpkgi2 := 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi5 := struct{}{}\n\tvpkgi6 := struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi7 := StructT{}\n\tvpkgi8 := StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgi9 := 1\n\tvpkgi10 := value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit1 int = 1\n\t\tvpkgit2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit7 StructT = StructT{}\n\t\tvpkgit8 StructT = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\t\tvpkgit9 int = 1\n\t\tvpkgit10 int = value()\n\t)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\t\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})
}

func TestVarPkg(t *testing.T) {
	t.Parallel()
	t.Run("no block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclPkg:        true,
			VarDeclPkgToBlock: false,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers (vpkg1, vpkg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkg1 int\nvar vpkg2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkg3, vpkg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkg5, vpkg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkg7, vpkg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkg7 StructT\nvar vpkg8 StructT",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb1, vpkgb2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgb1 int\n\tvpkgb2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb3, vpkgb4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb5, vpkgb6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb7, vpkgb8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgb7 StructT\n\tvpkgb8 StructT",
					},
				}},
			},
		})
	})

	t.Run("block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclPkg:        true,
			VarDeclPkgToBlock: true,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers (vpkg1, vpkg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkg1 int\n\tvpkg2 int\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkg3, vpkg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkg5, vpkg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkg7, vpkg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkg7 StructT\n\tvpkg8 StructT\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb1, vpkgb2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgb1 int\n\tvpkgb2 int",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb3, vpkgb4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb5, vpkgb6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers (vpkgb7, vpkgb8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgb7 StructT\n\tvpkgb8 StructT",
					},
				}},
			},
		})
	})
}

func TestVarInitPkg(t *testing.T) {
	t.Parallel()
	t.Run("no block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitPkg:        true,
			VarDeclInitPkgToBlock: false,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi1 = 1\nvar vpkgi2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi5 = struct{}{}\nvar vpkgi6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi7 = StructT{}\nvar vpkgi8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgi9 = 1\nvar vpkgi10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit1 int = 1\nvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\nvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var vpkgit9 int = 1\nvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})

	t.Run("block", func(t *testing.T) {
		t.Parallel()
		run(t, multisplit.Settings{
			VarDeclInitPkg:        true,
			VarDeclInitPkgToBlock: true,
		}, []testResult{
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi1, vpkgi2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgi1 = 1\n\tvpkgi2 = 2\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi3, vpkgi4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi5, vpkgi6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgi5 = struct{}{}\n\tvpkgi6 = struct{}{}\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi7, vpkgi8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgi7 = StructT{}\n\tvpkgi8 = StructT{}\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgi9, vpkgi10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgi9 = 1\n\tvpkgi10 = value()\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig1, vpkgig2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig1 = 1\n\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig3, vpkgig4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig5, vpkgig6) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig5 = struct{}{}\n\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig7, vpkgig8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig7 = StructT{}\n\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgig9, vpkgig10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgig9 = 1\n\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit1, vpkgit2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgit1 int = 1\n\tvpkgit2 int = 2\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit3, vpkgit4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit5, vpkgit6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit7, vpkgit8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgit7 StructT = StructT{}\n\tvpkgit8 StructT = StructT{}\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgit9, vpkgit10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"var (\n\tvpkgit9 int = 1\n\tvpkgit10 int = value()\n)",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg1, vpkgitg2) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg1 int = 1\n\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg3, vpkgitg4) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg5, vpkgitg6) should be split into individual declarations",
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg7, vpkgitg8) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "variable declaration with multiple identifiers and initializers (vpkgitg9, vpkgitg10) should be split into individual declarations",
				suggestedFixes: []analysisFix{{
					message: "split into individual variable declarations",
					newText: []string{
						"vpkgitg9 int = 1\n\tvpkgitg10 int = value()",
					},
				}},
			},
		})
	})
}

type analysisFix struct {
	message string
	newText []string
}

type testResult struct {
	message        string
	suggestedFixes []analysisFix
}

func newTestResult(diag analysis.Diagnostic) testResult {
	fixes := make([]analysisFix, len(diag.SuggestedFixes))
	for idx, sfix := range diag.SuggestedFixes {
		fix := analysisFix{
			message: sfix.Message,
		}
		for _, edit := range sfix.TextEdits {
			fix.newText = append(fix.newText, string(edit.NewText))
		}
		fixes[idx] = fix
	}

	return testResult{
		message:        diag.Message,
		suggestedFixes: fixes,
	}
}

func run(t *testing.T, cfg multisplit.Settings, expected []testResult) {
	t.Helper()

	analyzer := multisplit.NewAnalyzer()
	analyzer.Settings = cfg
	results := analysistest.Run(&testingDummy{}, pkgTestPath(t, "../testdata"), analyzer.Analyzer, "./...")

	var testResults []testResult
	for _, result := range results {
		for _, diag := range result.Diagnostics {
			testResults = append(testResults, newTestResult(diag))
		}
	}

	assert.Len(t, testResults, len(expected), "number of diagnostics should match")
	for idx, testResult := range testResults {
		if idx >= len(expected) {
			break
		}
		assert.Equal(t, expected[idx].message, testResult.message, "[%d] diagnostic message should match", idx)
		assert.ElementsMatch(t, expected[idx].suggestedFixes, testResult.suggestedFixes, "[%d] suggested fixes should match", idx)
	}
}

func pkgTestPath(t *testing.T, path string) string {
	t.Helper()
	testdata, err := filepath.Abs(path)
	require.NoError(t, err)

	return testdata
}

type testingDummy struct{}

func (*testingDummy) Errorf(_format string, _args ...any) {}
