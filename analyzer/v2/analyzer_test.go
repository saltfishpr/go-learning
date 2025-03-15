package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, `E:\Programming\go\go-learning\testdata\src\a`, Analyzer, "./...")
}
