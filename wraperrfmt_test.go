package wraperrfmt_test

import (
	"testing"

	"github.com/gostaticanalysis/wraperrfmt"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, wraperrfmt.Analyzer, "a")
}
