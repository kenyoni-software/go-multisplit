package plugin

import (
	"errors"
	"fmt"

	"github.com/golangci/plugin-module-register/register"
	"github.com/kenyoni-software/go-multisplit/multisplit"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("multisplit", NewMultiSplitPlugin)
}

var ErrUnknownRule = errors.New("unknown rule")

type Settings struct {
	Rules                  []string `json:"rules"`
	VarDeclPkgToBlock      *bool    `json:"varDeclPkgToBlock"`
	VarDeclFuncToBlock     *bool    `json:"varDeclFuncToBlock"`
	VarDeclInitPkgToBlock  *bool    `json:"varDeclInitPkgToBlock"`
	VarDeclInitFuncToBlock *bool    `json:"varDeclInitFuncToBlock"`
	VarDeclInitFuncToShort *bool    `json:"varDeclInitFuncToShort"`
	ConstDeclPkgToBlock    *bool    `json:"constDeclPkgToBlock"`
	ConstDeclFuncToBlock   *bool    `json:"constDeclFuncToBlock"`
}

func (s *Settings) toMultiSplitSettings() (multisplit.Settings, error) {
	var cfg multisplit.Settings
	if len(s.Rules) == 0 {
		cfg = multisplit.DefaultSettings()
	} else {
		for _, rule := range s.Rules {
			switch rule {
			case "assign":
				cfg.Assign = true
			case "const-decl-pkg":
				cfg.ConstDeclPkg = true
			case "const-decl-func":
				cfg.ConstDeclFunc = true
			case "func-params":
				cfg.FuncParams = true
			case "func-return-values":
				cfg.FuncReturnValues = true
			case "var-decl-func":
				cfg.VarDeclFunc = true
			case "var-decl-pkg":
				cfg.VarDeclPkg = true
			case "var-decl-init-pkg":
				cfg.VarDeclInitPkg = true
			case "var-decl-init-func":
				cfg.VarDeclInitFunc = true
			case "short-var-decl":
				cfg.ShortVarDecl = true
			case "struct-fields":
				cfg.StructFields = true
			default:
				return multisplit.Settings{}, fmt.Errorf("%w: %s", ErrUnknownRule, rule)
			}
		}
	}
	if s.VarDeclPkgToBlock != nil {
		cfg.VarDeclPkgToBlock = *s.VarDeclPkgToBlock
	}
	if s.VarDeclFuncToBlock != nil {
		cfg.VarDeclFuncToBlock = *s.VarDeclFuncToBlock
	}
	if s.VarDeclInitPkgToBlock != nil {
		cfg.VarDeclInitPkgToBlock = *s.VarDeclInitPkgToBlock
	}
	if s.VarDeclInitFuncToBlock != nil {
		cfg.VarDeclInitFuncToBlock = *s.VarDeclInitFuncToBlock
	}
	if s.VarDeclInitFuncToShort != nil {
		cfg.VarDeclInitFuncToShort = *s.VarDeclInitFuncToShort
	}
	if s.ConstDeclPkgToBlock != nil {
		cfg.ConstDeclPkgToBlock = *s.ConstDeclPkgToBlock
	}
	if s.ConstDeclFuncToBlock != nil {
		cfg.ConstDeclFuncToBlock = *s.ConstDeclFuncToBlock
	}

	return cfg, nil
}

type Plugin struct {
	settings multisplit.Settings
}

// NewMultiSplitPlugin constructs a new multisplit plugin.
func NewMultiSplitPlugin(settings any) (register.LinterPlugin, error) {
	ds, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	cfg, err := ds.toMultiSplitSettings()
	if err != nil {
		return nil, err
	}

	return &Plugin{
		settings: cfg,
	}, nil
}

// BuildAnalyzers returns the analyzer instances that this plugin provides.
func (plugin *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer := multisplit.NewAnalyzer()
	analyzer.Settings = plugin.settings
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (plugin *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
