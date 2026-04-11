package multisplit

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

// walker implements ast.Visitor.
type walker struct {
	an *Analyzer

	pass *analysis.Pass
	// inFunc is true when the walker is inside a function body.
	inFunc bool
	// forInit and forPost are set to the AssignStmts in the nearest enclosing for loop's Init and Post clauses
	forInit    *ast.AssignStmt
	forPost    *ast.AssignStmt
	commentMap ast.CommentMap
}

// Visit implements ast.Visitor.
func (w *walker) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch node := n.(type) {
	case *ast.AssignStmt:
		w.checkAssignStmt(node)
	case *ast.ForStmt:
		if !w.an.Settings.Assign && !w.an.Settings.ShortVarDecl {
			return nil
		}

		subW := &walker{
			an:         w.an,
			pass:       w.pass,
			inFunc:     w.inFunc,
			commentMap: w.commentMap,
		}
		if stmt, ok := node.Init.(*ast.AssignStmt); ok {
			subW.forInit = stmt
		}
		if stmt, ok := node.Post.(*ast.AssignStmt); ok {
			subW.forPost = stmt
		}

		return subW
	case *ast.FuncDecl, *ast.FuncLit:
		subW := *w
		subW.inFunc = true

		return &subW
	case *ast.FuncType:
		w.checkFuncType(node)
	case *ast.GenDecl:
		w.checkGenDecl(node)
	case *ast.StructType:
		w.checkStruct(node)
	default:
		return w
	}

	return w
}

// checkAssignStmt checks whether stmt is a multiple assignment that should be split and reports it if so.
// The following cases are silently skipped because they cannot be split:
//   - Single RHS call expression (multi-return function): a, b = f()
//   - Fewer RHS values than LHS targets (map index, type assertion, channel receive): a, ok = m[key]
//   - All LHS targets are blank identifiers: _, _ = a, b
func (w *walker) checkAssignStmt(node *ast.AssignStmt) {
	if (w.an.Settings.Assign || w.an.Settings.ShortVarDecl) && node != w.forInit && node != w.forPost {
		// only proceed if it's a multi-assignment where identifiers match values 1 to 1
		if len(node.Lhs) <= 1 || len(node.Lhs) != len(node.Rhs) || allBlank(node.Lhs) {
			return
		}

		if node.Tok == token.ASSIGN && w.an.Settings.Assign {
			reportAssignStmt(w.pass, node, w.commentMap)
		}

		if node.Tok == token.DEFINE && w.an.Settings.ShortVarDecl {
			reportAssignStmt(w.pass, node, w.commentMap)
		}
	}
}

func (w *walker) checkFuncType(node *ast.FuncType) {
	if w.an.Settings.FuncParams {
		checkFieldList(w.pass, node.Params, fieldListFuncParams)
	}

	if w.an.Settings.FuncReturnValues {
		checkFieldList(w.pass, node.Results, fieldListFuncResults)
	}
}

func (w *walker) checkGenDecl(node *ast.GenDecl) {
	for _, spec := range node.Specs {
		vspec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		isInit := len(vspec.Values) > 0
		// only multiple variable declarations are relevant
		if len(vspec.Names) <= 1 || isInit && len(vspec.Values) != len(vspec.Names) {
			continue
		}

		enabled, toBlock, toShort := w.declConfig(node.Tok, isInit)
		if enabled {
			reportValueSpec(w.pass, node, vspec, toBlock, toShort)
		}
	}
}

func (w *walker) checkStruct(node *ast.StructType) {
	if !w.an.Settings.StructFields {
		return
	}

	checkFieldList(w.pass, node.Fields, fieldListStructFields)
}

//nolint:gocognit,gocyclo
func (w *walker) declConfig(tok token.Token, isInit bool) (enabled bool, toBlock bool, toShort bool) {
	cfg := w.an.Settings

	switch tok { //nolint:exhaustive
	case token.VAR:
		if w.inFunc {
			enabled = (isInit && cfg.VarDeclInitFunc) || (!isInit && cfg.VarDeclFunc)
			toBlock = (isInit && cfg.VarDeclInitFuncToBlock) || (!isInit && cfg.VarDeclFuncToBlock)
			toShort = isInit && cfg.VarDeclInitFuncToShort && cfg.VarDeclInitFunc
		} else {
			enabled = (isInit && cfg.VarDeclInitPkg) || (!isInit && cfg.VarDeclPkg)
			toBlock = (isInit && cfg.VarDeclInitPkgToBlock) || (!isInit && cfg.VarDeclPkgToBlock)
		}

	case token.CONST:
		if w.inFunc {
			enabled = cfg.ConstDeclFunc
			toBlock = cfg.ConstDeclFuncToBlock
		} else {
			enabled = cfg.ConstDeclPkg
			toBlock = cfg.ConstDeclPkgToBlock
		}
	default:
		enabled = false
	}

	return enabled, toBlock, toShort
}
