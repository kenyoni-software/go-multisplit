//nolint:dupl,funlen
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
			message: "multiple assignment (v1, v2, v3, v5) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split multiple assignment",
				newText: []string{
					"v1 = 1\n\tv2 = value()\n\tv3 = struct{}{}\n\tv5 = StructT{}",
				},
			}},
		},
		{
			message: "multiple assignment (v3, v4) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split multiple assignment",
				newText: []string{
					"v3 = struct{}{}\n\tv4 = struct{}{}",
				},
			}},
		},
		{
			message: "multiple assignment (v1, v2) should be split into individual assignments",
		},
		{
			message: "multiple assignment (v1, _) should be split into individual assignments",
			suggestedFixes: []analysisFix{{
				message: "split multiple assignment",
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
				message: "multiple constant declaration (cfunc1, cfunc2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const cfunc1 = 1\n\tconst cfunc2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunc3, cfunc4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfuncg1, cfuncg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cfuncg1 = 1\n\t\tcfuncg2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfuncg3, cfuncg4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfunct1, cfunct2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const cfunct1 int = 1\n\tconst cfunct2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunct3, cfunct4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfunctg1, cfunctg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cfunctg1 int = 1\n\t\tcfunctg2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunctg3, cfunctg4) should be split into individual constants",
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
				message: "multiple constant declaration (cfunc1, cfunc2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const (\n\t\tcfunc1 = 1\n\t\tcfunc2 = 2\n\t)",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunc3, cfunc4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfuncg1, cfuncg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cfuncg1 = 1\n\t\tcfuncg2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfuncg3, cfuncg4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfunct1, cfunct2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const (\n\t\tcfunct1 int = 1\n\t\tcfunct2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunct3, cfunct4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cfunctg1, cfunctg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cfunctg1 int = 1\n\t\tcfunctg2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cfunctg3, cfunctg4) should be split into individual constants",
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
				message: "multiple constant declaration (cpkgi1, cpkgi2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const cpkgi1 = 1\nconst cpkgi2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgi3, cpkgi4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgig1, cpkgig2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cpkgig1 = 1\n\tcpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgig3, cpkgig4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgit1, cpkgit2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const cpkgit1 int = 1\nconst cpkgit2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgit3, cpkgit4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgitg1, cpkgitg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cpkgitg1 int = 1\n\tcpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgitg3, cpkgitg4) should be split into individual constants",
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
				message: "multiple constant declaration (cpkgi1, cpkgi2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const (\n\tcpkgi1 = 1\n\tcpkgi2 = 2\n)",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgi3, cpkgi4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgig1, cpkgig2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cpkgig1 = 1\n\tcpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgig3, cpkgig4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgit1, cpkgit2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"const (\n\tcpkgit1 int = 1\n\tcpkgit2 int = 2\n)",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgit3, cpkgit4) should be split into individual constants",
			},
			{
				message: "multiple constant declaration (cpkgitg1, cpkgitg2) should be split into individual constants",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"cpkgitg1 int = 1\n\tcpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple constant declaration (cpkgitg3, cpkgitg4) should be split into individual constants",
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
			message: "multiple function parameters (p1, p2, p3) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p1 int, p2 int, p3 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p4, p5, p6) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p7, p8, p9) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p7 StructT, p8 StructT, p9 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p11, p12, p13) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p11 int, p12 int, p13 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p14, p15, p16) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p17, p18, p19) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p17 StructT, p18 StructT, p19 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p21, p22, p23) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p21 int, p22 int, p23 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p24, p25, p26) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p27, p28, p29) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p27 StructT, p28 StructT, p29 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p31, p32, p33) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p31 int, p32 int, p33 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p34, p35, p36) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p37, p38, p39) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p37 StructT, p38 StructT, p39 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p41, p42, p43) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p41 int, p42 int, p43 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p44, p45, p46) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p47, p48, p49) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p47 StructT, p48 StructT, p49 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p51, p52, p53) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p51 int, p52 int, p53 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p54, p55, p56) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p57, p58, p59) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p57 StructT, p58 StructT, p59 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p61, p62, p63) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p61 int, p62 int, p63 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p64, p65, p66) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p67, p68, p69) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p67 StructT, p68 StructT, p69 StructT",
				},
			}},
		},
		{
			message: "multiple function parameters (p71, p72, p73) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
				newText: []string{
					"p71 int, p72 int, p73 int",
				},
			}},
		},
		{
			message: "multiple function parameters (p74, p75, p76) should be split into individual parameters",
		},
		{
			message: "multiple function parameters (p77, p78, p79) should be split into individual parameters",
			suggestedFixes: []analysisFix{{
				message: "split multiple function parameters",
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
			message: "multiple function return values (r1, r2, r3) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r1 int, r2 int, r3 int",
				},
			}},
		},
		{
			message: "multiple function return values (r4, r5, r6) should be split into individual return values",
		},
		{
			message: "multiple function return values (r7, r8, r9) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r7 StructT, r8 StructT, r9 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r11, r12, r13) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r11 int, r12 int, r13 int",
				},
			}},
		},
		{
			message: "multiple function return values (r14, r15, r16) should be split into individual return values",
		},
		{
			message: "multiple function return values (r17, r18, r19) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r17 StructT, r18 StructT, r19 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r21, r22, r23) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r21 int, r22 int, r23 int",
				},
			}},
		},
		{
			message: "multiple function return values (r24, r25, r26) should be split into individual return values",
		},
		{
			message: "multiple function return values (r27, r28, r29) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r27 StructT, r28 StructT, r29 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r31, r32, r33) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r31 int, r32 int, r33 int",
				},
			}},
		},
		{
			message: "multiple function return values (r34, r35, r36) should be split into individual return values",
		},
		{
			message: "multiple function return values (r37, r38, r39) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r37 StructT, r38 StructT, r39 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r41, r42, r43) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r41 int, r42 int, r43 int",
				},
			}},
		},
		{
			message: "multiple function return values (r44, r45, r46) should be split into individual return values",
		},
		{
			message: "multiple function return values (r47, r48, r49) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r47 StructT, r48 StructT, r49 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r51, r52, r53) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r51 int, r52 int, r53 int",
				},
			}},
		},
		{
			message: "multiple function return values (r54, r55, r56) should be split into individual return values",
		},
		{
			message: "multiple function return values (r57, r58, r59) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r57 StructT, r58 StructT, r59 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r61, r62, r63) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r61 int, r62 int, r63 int",
				},
			}},
		},
		{
			message: "multiple function return values (r64, r65, r66) should be split into individual return values",
		},
		{
			message: "multiple function return values (r67, r68, r69) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r67 StructT, r68 StructT, r69 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r71, r72, r73) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r71 int, r72 int, r73 int",
				},
			}},
		},
		{
			message: "multiple function return values (r74, r75, r76) should be split into individual return values",
		},
		{
			message: "multiple function return values (r77, r78, r79) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r77 StructT, r78 StructT, r79 StructT",
				},
			}},
		},
		{
			message: "multiple function return values (r81, r82, r83) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
				newText: []string{
					"r81 int, r82 int, r83 int",
				},
			}},
		},
		{
			message: "multiple function return values (r84, r85, r86) should be split into individual return values",
		},
		{
			message: "multiple function return values (r87, r88, r89) should be split into individual return values",
			suggestedFixes: []analysisFix{{
				message: "split multiple function return values",
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
			message: "multiple short variable declaration (s1, s2) should be split into individual short declarations",
			suggestedFixes: []analysisFix{{
				message: "split multiple short variable declaration",
				newText: []string{
					"s1 := 1\n\ts2 := 2",
				},
			}},
		},
		{
			message: "multiple short variable declaration (s3, s4) should be split into individual short declarations",
		},
		{
			message: "multiple short variable declaration (s7, s8) should be split into individual short declarations",
			suggestedFixes: []analysisFix{{
				message: "split multiple short variable declaration",
				newText: []string{
					"s7 := struct{}{}\n\ts8 := struct{}{}",
				},
			}},
		},
		{
			message: "multiple short variable declaration (s9, _) should be split into individual short declarations",
			suggestedFixes: []analysisFix{{
				message: "split multiple short variable declaration",
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
			message: "multiple struct fields (v1, v2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"v1 int\n\tv2 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (v3, v4) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"v3 string `tag:\"value\"`\n\tv4 string `tag:\"value\"`",
					},
				},
			},
		},
		{
			message: "multiple struct fields (v5, v6) should be split into individual fields",
		},
		{
			message: "multiple struct fields (v7, v8) should be split into individual fields",
		},
		{
			message: "multiple struct fields (v9, v10) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"v9 int\n\t\tv10 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (spv1, spv2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"spv1 int; spv2 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (spv2, spv3) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"spv2 int; spv3 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (sv1, sv2) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"sv1 int\n\t\tsv2 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (sv3, sv4) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"sv3 string `tag:\"value\"`\n\t\tsv4 string `tag:\"value\"`",
					},
				},
			},
		},
		{
			message: "multiple struct fields (sv5, sv6) should be split into individual fields",
		},
		{
			message: "multiple struct fields (sv7, sv8) should be split into individual fields",
		},
		{
			message: "multiple struct fields (sv9, sv10) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
					newText: []string{
						"sv9 int\n\t\t\tsv10 int",
					},
				},
			},
		},
		{
			message: "multiple struct fields (spv2, spv3) should be split into individual fields",
			suggestedFixes: []analysisFix{
				{
					message: "split multiple struct fields",
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
				message: "multiple variable declaration (vfunc1, vfunc2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vfunc1 int\n\tvar vfunc2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfunc3, vfunc4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfunc5, vfunc6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfunc7, vfunc8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vfunc7 StructT\n\tvar vfunc8 StructT",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfuncb1, vfuncb2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vfuncb1 int\n\t\tvfuncb2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfuncb3, vfuncb4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfuncb5, vfuncb6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfuncb7, vfuncb8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration (vfunc1, vfunc2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvfunc1 int\n\t\tvfunc2 int\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfunc3, vfunc4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfunc5, vfunc6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfunc7, vfunc8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvfunc7 StructT\n\t\tvfunc8 StructT\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfuncb1, vfuncb2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vfuncb1 int\n\t\tvfuncb2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vfuncb3, vfuncb4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfuncb5, vfuncb6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vfuncb7, vfuncb8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi1 = 1\n\tvar vpkgi2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi5 = struct{}{}\n\tvar vpkgi6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi7 = StructT{}\n\tvar vpkgi8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi9 = 1\n\tvar vpkgi10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit1 int = 1\n\tvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\n\tvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit9 int = 1\n\tvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgi1 = 1\n\t\tvpkgi2 = 2\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgi5 = struct{}{}\n\t\tvpkgi6 = struct{}{}\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgi7 = StructT{}\n\t\tvpkgi8 = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgi9 = 1\n\t\tvpkgi10 = value()\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit1 int = 1\n\t\tvpkgit2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit7 StructT = StructT{}\n\t\tvpkgit8 StructT = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit9 int = 1\n\t\tvpkgit10 int = value()\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi1 := 1\n\tvpkgi2 := 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi5 := struct{}{}\n\tvpkgi6 := struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi7 := StructT{}\n\tvpkgi8 := StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi9 := 1\n\tvpkgi10 := value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit1 int = 1\n\tvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\n\tvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit9 int = 1\n\tvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi1 := 1\n\tvpkgi2 := 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi5 := struct{}{}\n\tvpkgi6 := struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi7 := StructT{}\n\tvpkgi8 := StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgi9 := 1\n\tvpkgi10 := value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\t\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\t\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\t\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\t\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit1 int = 1\n\t\tvpkgit2 int = 2\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit7 StructT = StructT{}\n\t\tvpkgit8 StructT = StructT{}\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\t\tvpkgit9 int = 1\n\t\tvpkgit10 int = value()\n\t)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\t\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\t\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration (vpkg1, vpkg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkg1 int\nvar vpkg2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkg3, vpkg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkg5, vpkg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkg7, vpkg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkg7 StructT\nvar vpkg8 StructT",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkgb1, vpkgb2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgb1 int\n\tvpkgb2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkgb3, vpkgb4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkgb5, vpkgb6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkgb7, vpkgb8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration (vpkg1, vpkg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkg1 int\n\tvpkg2 int\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkg3, vpkg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkg5, vpkg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkg7, vpkg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkg7 StructT\n\tvpkg8 StructT\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkgb1, vpkgb2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgb1 int\n\tvpkgb2 int",
					},
				}},
			},
			{
				message: "multiple variable declaration (vpkgb3, vpkgb4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkgb5, vpkgb6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration (vpkgb7, vpkgb8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi1 = 1\nvar vpkgi2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi5 = struct{}{}\nvar vpkgi6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi7 = StructT{}\nvar vpkgi8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgi9 = 1\nvar vpkgi10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit1 int = 1\nvar vpkgit2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit7 StructT = StructT{}\nvar vpkgit8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var vpkgit9 int = 1\nvar vpkgit10 int = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
				message: "multiple variable declaration with initializer (vpkgi1, vpkgi2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgi1 = 1\n\tvpkgi2 = 2\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi3, vpkgi4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgi5, vpkgi6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgi5 = struct{}{}\n\tvpkgi6 = struct{}{}\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi7, vpkgi8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgi7 = StructT{}\n\tvpkgi8 = StructT{}\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgi9, vpkgi10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgi9 = 1\n\tvpkgi10 = value()\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig1, vpkgig2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig1 = 1\n\tvpkgig2 = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig3, vpkgig4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgig5, vpkgig6) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig5 = struct{}{}\n\tvpkgig6 = struct{}{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig7, vpkgig8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig7 = StructT{}\n\tvpkgig8 = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgig9, vpkgig10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgig9 = 1\n\tvpkgig10 = value()",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit1, vpkgit2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgit1 int = 1\n\tvpkgit2 int = 2\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit3, vpkgit4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit5, vpkgit6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgit7, vpkgit8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgit7 StructT = StructT{}\n\tvpkgit8 StructT = StructT{}\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgit9, vpkgit10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"var (\n\tvpkgit9 int = 1\n\tvpkgit10 int = value()\n)",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg1, vpkgitg2) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg1 int = 1\n\tvpkgitg2 int = 2",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg3, vpkgitg4) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg5, vpkgitg6) should be split into individual variables",
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg7, vpkgitg8) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
					newText: []string{
						"vpkgitg7 StructT = StructT{}\n\tvpkgitg8 StructT = StructT{}",
					},
				}},
			},
			{
				message: "multiple variable declaration with initializer (vpkgitg9, vpkgitg10) should be split into individual variables",
				suggestedFixes: []analysisFix{{
					message: "split multiple declaration",
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
