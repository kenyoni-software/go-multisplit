package multisplit_test

import (
	"bytes"
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

// testFix dir is the path relative to testdata pointing to the directory containing the in/out directories.
func testFix(t *testing.T, analyzer *analysis.Analyzer, dir string) {
	t.Helper()

	baseDir := filepath.Join(analysistest.TestData(), dir)
	inDir := filepath.Join(baseDir, "in")
	outDir := filepath.Join(baseDir, "out")
	results := analysistest.Run(&testingDummy{}, inDir, analyzer, "./...")
	expectedFilenames := collectExpected(t, outDir)

	for _, res := range results {
		require.NoError(t, res.Err)
		editsByFiles, err := collectEdits(res)
		require.NoError(t, err)

		for file, edits := range editsByFiles {
			src, err := os.ReadFile(file.Name())
			require.NoError(t, err)
			out := applyEdits(src, file, edits)

			srcRelName, err := filepath.Rel(inDir, file.Name())
			require.NoError(t, err)

			expFilename := filepath.Join(outDir, srcRelName)
			var expect []byte
			expectedFileIdx := slices.Index(expectedFilenames, expFilename)
			// if the expected file does not exist, assume no change
			if expectedFileIdx == -1 {
				expect = src
			} else {
				expect, err = os.ReadFile(expFilename) //nolint:gosec
				require.NoError(t, err)
				expectedFilenames = slices.Delete(expectedFilenames, expectedFileIdx, expectedFileIdx+1)
			}
			diff := cmp.Diff(string(normalizeLineEndings(expect)), string(normalizeLineEndings(out)))

			if diff != "" && expectedFileIdx == -1 {
				t.Errorf("unexpected diff for %s\n\n%s", file.Name(), diff)
			} else if diff != "" {
				t.Errorf("unexpected diff for %s - %s\n\n%s", file.Name(), expFilename, diff)
			}
		}
	}

	for _, filename := range expectedFilenames {
		expect, err := os.ReadFile(filename) //nolint:gosec
		require.NoError(t, err)

		expRelName, err := filepath.Rel(outDir, filename)
		require.NoError(t, err)
		srcFilename := filepath.Join(inDir, expRelName)
		src, err := os.ReadFile(srcFilename) //nolint:gosec
		require.NoError(t, err)

		diff := cmp.Diff(string(normalizeLineEndings(expect)), string(normalizeLineEndings(src)))
		if diff != "" {
			t.Errorf("unexpected diff for %s - %s\n\n%s", srcFilename, filename, diff)
		}
	}
}

func collectExpected(t *testing.T, dir string) []string {
	t.Helper()

	expectedFiles := []string{}
	err := filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		expectedFiles = append(expectedFiles, path)

		return err
	})
	require.NoError(t, err)

	return expectedFiles
}

// collectEdits collects the suggested fixes from the analysis result and groups them by file.
// The returned edits are sorted in order of their position and do not overlap with each other.
//
//nolint:gocognit
func collectEdits(result *analysistest.Result) (map[*token.File][]analysis.TextEdit, error) {
	var err error
	editsByFile := map[*token.File][]analysis.TextEdit{}
	for _, diag := range result.Diagnostics {
		for _, fix := range diag.SuggestedFixes {
			for _, edit := range fix.TextEdits {
				posFile := result.Pass.Fset.File(diag.Pos)
				endFile := result.Pass.Fset.File(edit.End)
				if posFile == nil || endFile == nil || posFile != endFile {
					return nil, fmt.Errorf("unexpected edit with different files: %s and %s", posFile.Name(), endFile.Name()) //nolint:err113
				}

				editsByFile[posFile], err = appendEdit(editsByFile[posFile], edit)
				if err != nil {
					return nil, fmt.Errorf("failed to add edit for file %s: %w", posFile.Name(), err)
				}
			}
		}
	}

	for _, edit := range editsByFile {
		slices.SortFunc(edit, func(lhs analysis.TextEdit, rhs analysis.TextEdit) int {
			return int(lhs.Pos - rhs.Pos)
		})
	}

	return editsByFile, nil
}

func applyEdits(src []byte, file *token.File, edits []analysis.TextEdit) []byte {
	out := []byte{}

	offset := 0
	for _, edit := range edits {
		start := file.Offset(edit.Pos)
		out = append(out, src[offset:start]...)
		out = append(out, edit.NewText...)
		offset = file.Offset(edit.End)
	}
	if offset < len(src) {
		out = append(out, src[offset:]...)
	}

	return out
}

func appendEdit(edits []analysis.TextEdit, newEdit analysis.TextEdit) ([]analysis.TextEdit, error) {
	for _, edit := range edits {
		if (newEdit.Pos <= edit.Pos && newEdit.End > edit.Pos) || (newEdit.Pos < edit.End && newEdit.End >= edit.Pos) {
			return nil, fmt.Errorf("overlapping edit with position %d - %d", edit.Pos, edit.End) //nolint:err113
		}
	}

	return append(edits, newEdit), nil
}

func normalizeLineEndings(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
}

type testingDummy struct{}

func (*testingDummy) Errorf(_format string, _args ...any) {}
