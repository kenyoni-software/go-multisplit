package multisplit_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/kenyoni-software/go-multisplit/multisplit"
)

func TestAnalyzerDiagnstic(t *testing.T) {
	t.Parallel()

	analyzer := multisplit.NewAnalyzer()
	analyzer.Settings = multisplit.Settings{
		Assign:           true,
		ConstDeclFunc:    true,
		ConstDeclPkg:     true,
		FuncParams:       true,
		FuncReturnValues: true,
		ShortVarDecl:     true,
		StructFields:     true,
		VarDeclFunc:      true,
		VarDeclPkg:       true,
		VarDeclInitFunc:  true,
		VarDeclInitPkg:   true,
	}
	analysistest.Run(t, filepath.Join(analysistest.TestData(), "diagnostic"), analyzer.Analyzer, "./...")
}

func TestAnalyzerFix(t *testing.T) {
	t.Parallel()

	baseCfg := multisplit.Settings{
		Assign:           true,
		ConstDeclFunc:    true,
		ConstDeclPkg:     true,
		FuncParams:       true,
		FuncReturnValues: true,
		ShortVarDecl:     true,
		StructFields:     true,
		VarDeclFunc:      true,
		VarDeclPkg:       true,
		VarDeclInitFunc:  true,
		VarDeclInitPkg:   true,
	}

	t.Run("block", func(t *testing.T) {
		t.Parallel()

		cfg := baseCfg
		cfg.ConstDeclFuncToBlock = true
		cfg.ConstDeclPkgToBlock = true
		cfg.VarDeclFuncToBlock = true
		cfg.VarDeclPkgToBlock = true
		cfg.VarDeclInitFuncToBlock = true
		cfg.VarDeclInitPkgToBlock = true

		analyzer := multisplit.NewAnalyzer()
		analyzer.Settings = cfg
		testFix(t, analyzer.Analyzer, "fix/block")
	})

	t.Run("block-short", func(t *testing.T) {
		t.Parallel()

		cfg := baseCfg
		cfg.ConstDeclFuncToBlock = true
		cfg.ConstDeclPkgToBlock = true
		cfg.VarDeclFuncToBlock = true
		cfg.VarDeclPkgToBlock = true
		cfg.VarDeclInitFuncToBlock = true
		cfg.VarDeclInitPkgToBlock = true
		cfg.VarDeclInitFuncToShort = true

		analyzer := multisplit.NewAnalyzer()
		analyzer.Settings = cfg
		testFix(t, analyzer.Analyzer, "fix/block-short")
	})

	t.Run("none", func(t *testing.T) {
		t.Parallel()

		analyzer := multisplit.NewAnalyzer()
		analyzer.Settings = baseCfg
		testFix(t, analyzer.Analyzer, "fix/none")
	})

	t.Run("short", func(t *testing.T) {
		t.Parallel()

		cfg := baseCfg
		cfg.VarDeclInitFuncToShort = true
		analyzer := multisplit.NewAnalyzer()
		analyzer.Settings = cfg
		testFix(t, analyzer.Analyzer, "fix/short")
	})
}
