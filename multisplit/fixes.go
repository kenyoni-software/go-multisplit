package multisplit

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// createValueSpecFix builds a fix for a non-parenthesized GenDecl.
func createValueSpecFix(fset *token.FileSet, decl *ast.GenDecl, vspec *ast.ValueSpec, toBlock bool, toShort bool) (analysis.SuggestedFix, error) {
	outer := indentAt(fset, decl.Pos())
	inner := outer + "\t"
	useShort := toShort && vspec.Type == nil

	var typeStr string
	if vspec.Type != nil {
		var err error
		typeStr, err = nodeToString(fset, vspec.Type)
		if err != nil {
			return analysis.SuggestedFix{}, err
		}
	}

	var sBuilder strings.Builder
	if toBlock && !useShort {
		_, _ = sBuilder.WriteString(decl.Tok.String() + " (\n")
	}

	for idx, ident := range vspec.Names {
		switch {
		case useShort:
			if idx > 0 {
				_, _ = sBuilder.WriteString("\n" + outer)
			}

			_, _ = sBuilder.WriteString(ident.Name)
		case toBlock:
			_, _ = sBuilder.WriteString(inner + ident.Name)
		default:
			// var name
			if idx > 0 {
				_, _ = sBuilder.WriteString("\n" + outer)
			}

			_, _ = sBuilder.WriteString(decl.Tok.String() + " " + ident.Name)
		}

		if typeStr != "" {
			_, _ = sBuilder.WriteString(" " + typeStr)
		}

		if idx < len(vspec.Values) {
			val, err := nodeToString(fset, vspec.Values[idx])
			if err != nil {
				return analysis.SuggestedFix{}, err
			}

			if useShort {
				_, _ = sBuilder.WriteString(" := " + val)
			} else {
				_, _ = sBuilder.WriteString(" = " + val)
			}
		}

		if toBlock && !useShort {
			_, _ = sBuilder.WriteString("\n")
		}
	}

	if toBlock && !useShort {
		_, _ = sBuilder.WriteString(outer + ")")
	}

	return analysis.SuggestedFix{
		Message: "split multiple declaration",
		TextEdits: []analysis.TextEdit{{
			Pos:     decl.Pos(),
			End:     decl.End(),
			NewText: []byte(sBuilder.String()),
		}},
	}, nil
}

// createBlockValueSpecFix builds a fix for a multiple ValueSpec that already sits inside a parenthesized GenDecl.
// Only the spec's own range is replaced so that sibling specs are preserved.
func createBlockValueSpecFix(fset *token.FileSet, vspec *ast.ValueSpec) (analysis.SuggestedFix, error) {
	var typeStr string
	if vspec.Type != nil {
		var err error
		typeStr, err = nodeToString(fset, vspec.Type)
		if err != nil {
			return analysis.SuggestedFix{}, err
		}
	}

	parts := make([]string, len(vspec.Names))
	for idx, ident := range vspec.Names {
		var sBuilder strings.Builder
		_, _ = sBuilder.WriteString(ident.Name)
		if typeStr != "" {
			_, _ = sBuilder.WriteString(" " + typeStr)
		}

		if idx < len(vspec.Values) {
			val, err := nodeToString(fset, vspec.Values[idx])
			if err != nil {
				return analysis.SuggestedFix{}, err
			}

			_, _ = sBuilder.WriteString(" = " + val)
		}

		parts[idx] = sBuilder.String()
	}

	return analysis.SuggestedFix{
		Message: "split multiple declaration",
		TextEdits: []analysis.TextEdit{{
			Pos:     vspec.Pos(),
			End:     vspec.End(),
			NewText: []byte(strings.Join(parts, "\n"+indentAt(fset, vspec.Pos()))),
		}},
	}, nil
}

// createFieldFix builds a fix that expands a single grouped *ast.Field into individually typed fields.
// The separator between fields is chosen by fieldSep based on whether the list is inline or multi-line.
func createFieldFix(fset *token.FileSet, flist *ast.FieldList, field *ast.Field, flt fieldListType) (analysis.SuggestedFix, error) {
	typeStr, err := nodeToString(fset, field.Type)
	if err != nil {
		return analysis.SuggestedFix{}, err
	}

	var tagStr string
	if field.Tag != nil {
		tagText, err := nodeToString(fset, field.Tag)
		if err != nil {
			return analysis.SuggestedFix{}, err
		}
		tagStr = " " + tagText
	}

	parts := make([]string, len(field.Names))
	for idx, ident := range field.Names {
		parts[idx] = ident.Name + " " + typeStr + tagStr
	}

	var fixMsg string
	switch flt {
	case fieldListFuncParams:
		fixMsg = "split multiple function parameters"
	case fieldListFuncResults:
		fixMsg = "split multiple function return values"
	case fieldListStructFields:
		fixMsg = "split multiple struct fields"
	default:
		fixMsg = "split multiple fields"
	}

	return analysis.SuggestedFix{
		Message: fixMsg,
		TextEdits: []analysis.TextEdit{{
			Pos:     field.Pos(),
			End:     field.End(),
			NewText: []byte(strings.Join(parts, fieldSep(fset, flist, field, flt))),
		}},
	}, nil
}

// createAssignFix splits a multi-LHS assignment into individual statements.
func createAssignFix(fset *token.FileSet, stmt *ast.AssignStmt) (analysis.SuggestedFix, error) {
	parts := make([]string, len(stmt.Lhs))
	for idx := range stmt.Lhs {
		lhs, err := nodeToString(fset, stmt.Lhs[idx])
		if err != nil {
			return analysis.SuggestedFix{}, err
		}

		rhs, err := nodeToString(fset, stmt.Rhs[idx])
		if err != nil {
			return analysis.SuggestedFix{}, err
		}
		tok := stmt.Tok.String()
		// if statement is assigning to "_" it must be "=" instead of ":="
		if stmt.Tok == token.DEFINE {
			if ident, ok := stmt.Lhs[idx].(*ast.Ident); ok && ident.Name == "_" {
				tok = token.ASSIGN.String()
			}
		}

		parts[idx] = lhs + " " + tok + " " + rhs
	}

	var msg string
	switch stmt.Tok { //nolint:exhaustive
	case token.ASSIGN:
		msg = "split multiple assignment"
	case token.DEFINE:
		msg = "split multiple short variable declaration"
	default:
		msg = "split multiple statement"
	}

	return analysis.SuggestedFix{
		Message: msg,
		TextEdits: []analysis.TextEdit{{
			Pos:     stmt.Pos(),
			End:     stmt.End(),
			NewText: []byte(strings.Join(parts, "\n"+indentAt(fset, stmt.Pos()))),
		}},
	}, nil
}

// nodeToString renders an AST node back to formatted Go source using the FileSet that was used to parse the file.
func nodeToString(fset *token.FileSet, node ast.Node) (string, error) {
	var buf bytes.Buffer
	err := format.Node(&buf, fset, node)
	if err != nil {
		return "", err //nolint:wrapcheck
	}

	return buf.String(), nil
}

// indentAt returns the tab-based whitespace prefix for the source line that contains pos.
// Column is 1-based and each tab counts as one byte, so (col - 1) tabs reproduces the indent level correctly for gofmt-formatted source files.
func indentAt(fset *token.FileSet, pos token.Pos) string {
	col := fset.Position(pos).Column // 1-based
	if col <= 1 {
		return ""
	}

	return strings.Repeat("\t", col-1)
}

// fieldSep returns the separator to use when joining expanded field declarations.
// When the field sits on the same line as the list's opening delimiter (typical for inline func params like "func f(a, b int)"),
// ", " is used so the result stays on one line.
// Otherwise a newline followed by the field's own indentation is returned.
func fieldSep(fset *token.FileSet, flist *ast.FieldList, field *ast.Field, flt fieldListType) string {
	if flist.Opening.IsValid() && fset.Position(flist.Opening).Line == fset.Position(field.Pos()).Line {
		if flt == fieldListStructFields {
			return "; "
		}

		return ", "
	}

	return "\n" + indentAt(fset, field.Pos())
}
