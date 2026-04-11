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
	var typeStr string
	var err error
	if vspec.Type != nil {
		typeStr, err = nodeToString(fset, vspec.Type)
		if err != nil {
			return analysis.SuggestedFix{}, err
		}
	}

	var edit string
	switch {
	// already inside an block
	case decl.Lparen.IsValid():
		edit, err = valueSpecEditInBlock(fset, decl, vspec, typeStr)
	case toShort && vspec.Type == nil:
		edit, err = valueSpecEditAsShort(fset, decl, vspec, typeStr)
	case toBlock:
		edit, err = valueSpecEditAsBlock(fset, decl, vspec, typeStr)
	default:
		edit, err = valueSpecEditIndividual(fset, decl, vspec, typeStr)
	}
	if err != nil {
		return analysis.SuggestedFix{}, err
	}

	return analysis.SuggestedFix{
		Message: valueFixMsg(decl.Tok),
		TextEdits: []analysis.TextEdit{{
			Pos:     decl.Pos(),
			End:     decl.End(),
			NewText: []byte(edit),
		}},
	}, nil
}

// createFieldFix builds a fix that expands a single block *ast.Field into individually typed fields.
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

	return analysis.SuggestedFix{
		Message: fieldFixMsg(flt),
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

	return analysis.SuggestedFix{
		Message: assignFixMsg(stmt.Tok),
		TextEdits: []analysis.TextEdit{{
			Pos:     stmt.Pos(),
			End:     stmt.End(),
			NewText: []byte(strings.Join(parts, "\n"+indentAt(fset, stmt.Pos()))),
		}},
	}, nil
}

func valueSpecEdit(fset *token.FileSet, vspec *ast.ValueSpec, typeStr string, indent string, declOp string, assignOp string) (string, error) {
	var sBuilder strings.Builder
	for idx, ident := range vspec.Names {
		if idx > 0 {
			_, _ = sBuilder.WriteString(indent)
		}
		_, _ = sBuilder.WriteString(declOp + ident.Name)
		if typeStr != "" {
			_, _ = sBuilder.WriteString(" " + typeStr)
		}
		if idx < len(vspec.Values) {
			val, err := nodeToString(fset, vspec.Values[idx])
			if err != nil {
				return "", err
			}
			_, _ = sBuilder.WriteString(" " + assignOp + " " + val)
		}
	}

	return sBuilder.String(), nil
}

func valueSpecEditIndividual(fset *token.FileSet, decl *ast.GenDecl, vspec *ast.ValueSpec, typeStr string) (string, error) {
	indent := indentAt(fset, decl.Pos())

	return valueSpecEdit(fset, vspec, typeStr, "\n"+indent, decl.Tok.String()+" ", "=")
}

// valueSpecEditInBlock creates inner part of an grouped block.
func valueSpecEditInBlock(fset *token.FileSet, decl *ast.GenDecl, vspec *ast.ValueSpec, typeStr string) (string, error) {
	indent := indentAt(fset, decl.Pos())

	return valueSpecEdit(fset, vspec, typeStr, "\n"+indent+"\t", "", "=")
}

// valueSpecEditAsBlock creates the whole grouped block for declarations.
func valueSpecEditAsBlock(fset *token.FileSet, decl *ast.GenDecl, vspec *ast.ValueSpec, typeStr string) (string, error) {
	indent := indentAt(fset, decl.Pos())
	inner, err := valueSpecEditInBlock(fset, decl, vspec, typeStr)
	if err != nil {
		return "", err
	}

	return decl.Tok.String() + " (\n" + indent + "\t" + inner + "\n" + indent + ")", nil
}

func valueSpecEditAsShort(fset *token.FileSet, decl *ast.GenDecl, vspec *ast.ValueSpec, typeStr string) (string, error) {
	indent := indentAt(fset, decl.Pos())

	return valueSpecEdit(fset, vspec, typeStr, "\n"+indent, "", ":=")
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
	col := fset.Position(pos).Column
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

func valueFixMsg(tok token.Token) string {
	switch tok { //nolint:exhaustive
	case token.VAR:
		return "split into individual variable declarations"
	case token.CONST:
		return "split into individual const declarations"
	default:
		return "split into individual declarations"
	}
}

func fieldFixMsg(flt fieldListType) string {
	switch flt {
	case fieldListFuncParams:
		return "split into individual parameters"
	case fieldListFuncResults:
		return "split into individual return values"
	case fieldListStructFields:
		return "split into individual struct fields"
	default:
		return "split into individual fields"
	}
}

func assignFixMsg(tok token.Token) string {
	switch tok { //nolint:exhaustive
	case token.ASSIGN:
		return "split into individual assignments"
	case token.DEFINE:
		return "split into individual short variable declarations"
	default:
		return "split into individual statements"
	}
}
