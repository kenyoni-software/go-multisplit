// Package multisplit provides a static analyzer which reports and fixes multiple declarations and assignments.
package multisplit

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Settings controls which multi-declaration checks are enabled.
type Settings struct {
	// VarDeclPkg checks multiple var declarations at package scope,
	VarDeclPkg bool
	// VarDeclPkgToBlock rewrites typed package-scope var declarations as a block.
	VarDeclPkgToBlock bool

	// VarDeclFunc checks multiple var declarations at function scope,
	VarDeclFunc bool
	// VarDeclFuncToBlock rewrites typed function-scope var declarations as a block.
	VarDeclFuncToBlock bool

	// VarDeclInitPkg checks multiple var declarations with an initializer at package scope.
	VarDeclInitPkg bool
	// VarDeclInitPkgToBlock rewrites typed package-scope var declarations with initializer as a block.
	VarDeclInitPkgToBlock bool

	// VarDeclInitFunc checks multiple var declarations with an initializer at function scope.
	VarDeclInitFunc bool
	// VarDeclInitFuncToBlock rewrites typed function-scope var declarations with initializer as a block.
	VarDeclInitFuncToBlock bool
	// VarDeclInitFuncToShort rewrites untyped function-scope var declarations with initializer as a short variable declaration.
	VarDeclInitFuncToShort bool

	// ConstDeclPkg checks multiple const declarations at package scope,
	ConstDeclPkg bool
	// ConstDeclPkgToBlock rewrites typed const declarations as a block.
	ConstDeclPkgToBlock bool

	// ConstDeclFunc checks multiple const declarations at function scope,
	ConstDeclFunc bool
	// ConstDeclFuncToBlock rewrites typed const declarations as a block.
	ConstDeclFuncToBlock bool

	// FuncParams checks multiple function parameters,
	FuncParams bool

	// FuncReturnValues checks multiple named return values,
	FuncReturnValues bool

	// Assign checks multiple plain assignments.
	Assign bool

	// ShortVarDecl checks multiple short variable declarations,
	ShortVarDecl bool

	// StructFields checks multiple struct fields,
	StructFields bool
}

// DefaultSettings returns a viable out-of-the-box configuration.
func DefaultSettings() Settings {
	return Settings{
		VarDeclPkg:             true,
		VarDeclPkgToBlock:      true,
		VarDeclFunc:            false,
		VarDeclFuncToBlock:     false,
		VarDeclInitPkg:         true,
		VarDeclInitPkgToBlock:  true,
		VarDeclInitFunc:        false,
		VarDeclInitFuncToBlock: false,
		VarDeclInitFuncToShort: true,
		ConstDeclPkg:           true,
		ConstDeclPkgToBlock:    true,
		ConstDeclFunc:          false,
		ConstDeclFuncToBlock:   false,
		FuncParams:             true,
		FuncReturnValues:       true,
		Assign:                 false,
		ShortVarDecl:           false,
		StructFields:           true,
	}
}

// allDisabled returns true when every check in s is turned off.
func (s Settings) allDisabled() bool {
	return !s.VarDeclPkg && !s.VarDeclFunc &&
		!s.VarDeclInitPkg && !s.VarDeclInitFunc &&
		!s.ConstDeclPkg && !s.ConstDeclFunc &&
		!s.FuncParams && !s.FuncReturnValues &&
		!s.Assign && !s.ShortVarDecl &&
		!s.StructFields
}

// Analyzer wraps the standard analysis.Analyzer with configurable Settings.
type Analyzer struct {
	*analysis.Analyzer

	Settings Settings
}

// NewAnalyzer constructs an Analyzer with DefaultSettings and registers all command-line flags.
func NewAnalyzer() *Analyzer {
	analyzer := &Analyzer{
		Analyzer: &analysis.Analyzer{
			Name: "multisplit",
			Doc:  "Split multiple declarations, assignments, function parameters/return values and struct fields into individual lines",
		},
		Settings: DefaultSettings(),
	}
	analyzer.Run = analyzer.run

	fs := flag.NewFlagSet("multisplit", flag.ExitOnError)
	registerFlags(fs, &analyzer.Settings)
	analyzer.Flags = *fs

	return analyzer
}

// registerFlags binds every Settings field to a named flag in fs.
func registerFlags(fSet *flag.FlagSet, cfg *Settings) {
	fSet.BoolFunc("split-all", "enable all split flags", func(string) error {
		cfg.VarDeclPkg = true
		cfg.VarDeclFunc = true
		cfg.VarDeclInitPkg = true
		cfg.VarDeclInitFunc = true
		cfg.ConstDeclPkg = true
		cfg.ConstDeclFunc = true
		cfg.FuncParams = true
		cfg.FuncReturnValues = true
		cfg.Assign = true
		cfg.ShortVarDecl = true
		cfg.StructFields = true

		return nil
	})
	const blockVarDesc = "when splitting, use a parenthesized block instead of separate 'var' lines"
	fSet.BoolVar(&cfg.VarDeclPkg, "var-decl-pkg", cfg.VarDeclPkg,
		"split multiple var declarations in package scope (e.g. 'var a, b int')")
	fSet.BoolVar(&cfg.VarDeclPkgToBlock, "var-decl-pkg-to-block", cfg.VarDeclPkgToBlock, blockVarDesc)
	fSet.BoolVar(&cfg.VarDeclFunc, "var-decl-func", cfg.VarDeclFunc,
		"split multiple var declarations in function scope (e.g. 'var a, b int')")
	fSet.BoolVar(&cfg.VarDeclFuncToBlock, "var-decl-func-to-block", cfg.VarDeclFuncToBlock, blockVarDesc)
	fSet.BoolVar(&cfg.VarDeclInitPkg, "var-decl-init-pkg", cfg.VarDeclInitPkg,
		"split multiple var declarations with initializer in package scope (e.g. 'var a, b = 1, 2')")
	fSet.BoolVar(&cfg.VarDeclInitPkgToBlock, "var-decl-init-pkg-to-block", cfg.VarDeclInitPkgToBlock, blockVarDesc)
	fSet.BoolVar(&cfg.VarDeclInitFunc, "var-decl-init-func", cfg.VarDeclInitFunc,
		"split multiple var declarations with initializer in function scope")
	fSet.BoolVar(&cfg.VarDeclInitFuncToBlock, "var-decl-init-func-to-block", cfg.VarDeclInitFuncToBlock, blockVarDesc)
	fSet.BoolVar(&cfg.VarDeclInitFuncToShort, "var-decl-init-func-to-short", cfg.VarDeclInitFuncToShort,
		"when splitting untyped function-scope var declarations with initializer, use short variable declaration syntax instead of 'var' lines")
	fSet.BoolVar(&cfg.ConstDeclPkg, "const-decl-pkg", cfg.ConstDeclPkg,
		"split multiple const declarations in package scope (e.g. 'const a, b = 1, 2')")
	fSet.BoolVar(&cfg.ConstDeclPkgToBlock, "const-decl-pkg-to-block", cfg.ConstDeclPkgToBlock,
		"when splitting, use a grouped block instead of separate 'const' lines")
	fSet.BoolVar(&cfg.ConstDeclFunc, "const-decl-func", cfg.ConstDeclFunc,
		"split multiple const declarations in function scope")
	fSet.BoolVar(&cfg.ConstDeclFuncToBlock, "const-decl-func-to-block", cfg.ConstDeclFuncToBlock,
		"when splitting, use a grouped block instead of separate 'const' lines")
	fSet.BoolVar(&cfg.FuncParams, "func-params", cfg.FuncParams,
		"split multiple function parameters (e.g. 'func f(a, b int)' → 'func f(a int, b int)')")
	fSet.BoolVar(&cfg.FuncReturnValues, "func-return-values", cfg.FuncReturnValues,
		"split multiple named function return values (e.g. 'func f() (a, b int)' → 'func f() (a int, b int)')")
	fSet.BoolVar(&cfg.Assign, "assign", cfg.Assign,
		"split multiple assignments (e.g. 'a, b = 1, 2')")
	fSet.BoolVar(&cfg.ShortVarDecl, "short-var-decl", cfg.ShortVarDecl,
		"split multiple short variable declarations (e.g. 'a, b := 1, 2')")
	fSet.BoolVar(&cfg.StructFields, "struct-fields", cfg.StructFields,
		"split multiple struct fields (e.g. 'type S struct { a, b int }')")
}

// run is the analysis entry point.
func (an *Analyzer) run(pass *analysis.Pass) (any, error) {
	if an.Settings.allDisabled() {
		return nil, nil //nolint:nilnil
	}

	for _, file := range pass.Files {
		ast.Walk(&walker{
			an:         an,
			pass:       pass,
			commentMap: ast.NewCommentMap(pass.Fset, file, file.Comments),
		}, file)
	}

	return nil, nil //nolint:nilnil
}
