package preprocessor

import (
	"testing"
)

func TestImportPreprocessor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`import ("test",)`, `let a = 5;
`},
	}
	for _, tt := range tests {
		pp := New(tt.input)
		pp.ResolveImports()

		if pp.Output != tt.expected {
			t.Fatalf("import parsing failed")
		}
	}
}
