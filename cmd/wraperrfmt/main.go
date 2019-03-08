package main

import (
	"github.com/gostaticanalysis/wraperrfmt"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(wraperrfmt.Analyzer) }
