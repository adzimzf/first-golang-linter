package main

import (
	nakedroutine "github.com/azdimzf/first-golang-linter/linters/nakedgoroutine"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(nakedroutine.NakedRoutineCodeAnalyzer)
}
