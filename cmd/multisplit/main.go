package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/kenyoni-software/go-multisplit/multisplit"
)

func main() {
	singlechecker.Main(multisplit.NewAnalyzer().Analyzer)
}
