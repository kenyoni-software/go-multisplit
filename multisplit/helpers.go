package multisplit

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// reportValueSpec reports a diagnostic when a ValueSpec declares more than one name (e.g. 'var a, b int' or 'const x, y = 1, 2').
func reportValueSpec(pass *analysis.Pass, decl *ast.GenDecl, vspec *ast.ValueSpec, toBlock bool, toShort bool) {
	var msgFmt string
	switch decl.Tok { //nolint:exhaustive
	case token.VAR:
		if len(vspec.Values) == 0 {
			msgFmt = "variable declaration with multiple identifiers (%s) should be split into individual declarations"
		} else {
			msgFmt = "variable declaration with multiple identifiers and initializers (%s) should be split into individual declarations"
		}
	case token.CONST:
		msgFmt = "const declaration with multiple identifiers (%s) should be split into individual declarations"
	default:
		return
	}

	diag := analysis.Diagnostic{
		Pos:     vspec.Pos(),
		End:     vspec.End(),
		Message: fmt.Sprintf(msgFmt, joinIdents(vspec.Names)),
	}

	_, isStructType := vspec.Type.(*ast.StructType)
	// only create a fix if we do not have a comment, so we may not change the intent
	// also, do not create a fix for struct type fields as this would lead to duplicate code
	if vspec.Comment == nil && !isStructType {
		fix, err := createValueSpecFix(pass.Fset, decl, vspec, toBlock, toShort)
		if err == nil {
			diag.SuggestedFixes = []analysis.SuggestedFix{fix}
		}
	}

	pass.Report(diag)
}

// fieldListType distinguishes where a *ast.FieldList appears so that diagnostic messages and fix formatting can be set accordingly.
type fieldListType int

const (
	fieldListFuncParams fieldListType = iota
	fieldListFuncResults
	fieldListStructFields
)

// checkFieldList reports a diagnostic for every *ast.Field in the list that carries more than one name (e.g. 'a, b int' in a struct or function signature).
func checkFieldList(pass *analysis.Pass, flist *ast.FieldList, flt fieldListType) {
	if flist == nil {
		return
	}

	var msgFmt string
	switch flt {
	case fieldListFuncParams:
		msgFmt = "function parameters with multiple identifiers (%s) should be split into individual parameters"
	case fieldListFuncResults:
		msgFmt = "function return values with multiple identifiers (%s) should be split into individual return values"
	case fieldListStructFields:
		msgFmt = "struct field declaration with multiple identifiers (%s) should be split into individual fields"
	default:
		return
	}

	for _, field := range flist.List {
		if len(field.Names) <= 1 {
			continue
		}

		diag := analysis.Diagnostic{
			Pos:     field.Pos(),
			End:     field.End(),
			Message: fmt.Sprintf(msgFmt, joinIdents(field.Names)),
		}
		// only create a fix if we do not have a comment, so we may not change the intent
		// also, do not create a fix for struct type fields as this would lead to duplicate code
		_, isStructType := field.Type.(*ast.StructType)
		if field.Comment == nil && !isStructType {
			fix, err := createFieldFix(pass.Fset, flist, field, flt)
			if err == nil {
				diag.SuggestedFixes = []analysis.SuggestedFix{fix}
			}
		}

		pass.Report(diag)
	}
}

// reportAssignStmt reports a diagnostic when an assignment has multiple LHS targets that can meaningfully be split into individual statements.
func reportAssignStmt(pass *analysis.Pass, stmt *ast.AssignStmt, commentMap ast.CommentMap) {
	var msgFmt string
	//nolint:exhaustive
	switch stmt.Tok {
	case token.ASSIGN:
		msgFmt = "assignment with multiple targets (%s) should be split into individual assignments"
	case token.DEFINE:
		msgFmt = "short variable declaration with multiple identifiers (%s) should be split into individual declarations"
	default:
		return
	}

	diag := analysis.Diagnostic{
		Pos:     stmt.Pos(),
		End:     stmt.End(),
		Message: fmt.Sprintf(msgFmt, joinLHSExprs(stmt.Lhs)),
	}

	// only create a fix if we do not have a comment, so we may not change the intent
	if commentMap[stmt] == nil {
		fix, err := createAssignFix(pass.Fset, stmt)
		if err == nil {
			diag.SuggestedFixes = []analysis.SuggestedFix{fix}
		}
	}

	pass.Report(diag)
}

// allBlank reports whether every expression in exprs is the blank identifier.
func allBlank(exprs []ast.Expr) bool {
	for _, expr := range exprs {
		ident, ok := expr.(*ast.Ident)
		if !ok || ident.Name != "_" {
			return false
		}
	}

	return true
}

// joinIdents returns a comma-separated string of identifier names.
func joinIdents(idents []*ast.Ident) string {
	parts := make([]string, len(idents))
	for idx, ident := range idents {
		parts[idx] = ident.Name
	}

	return strings.Join(parts, ", ")
}

// joinLHSExprs returns a comma-separated string of LHS expression source names.
// Non-identifier expressions (e.g. index expressions like 'a[idx]') are represented as "_" since they have no simple name.
func joinLHSExprs(exprs []ast.Expr) string {
	parts := make([]string, len(exprs))
	for idx, expr := range exprs {
		if ident, ok := expr.(*ast.Ident); ok {
			parts[idx] = ident.Name
		} else {
			parts[idx] = "_"
		}
	}

	return strings.Join(parts, ", ")
}
