package nakedroutine

import (
	"github.com/pkg/errors"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NakedRoutineCodeAnalyzer is a linter to determine does call naked go routine.
var NakedRoutineCodeAnalyzer = &analysis.Analyzer{
	Name:     "nakedroutine",
	Doc:      "linter for detecting context argument usage in functions",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

const errorMsg = "naked go routine is not allowed"

func run(pass *analysis.Pass) (interface{}, error) {
	i, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("analyzer is not type *inspector.Inspector")
	}

	nodeFilter := []ast.Node{
		(*ast.GoStmt)(nil),
	}

	i.Preorder(nodeFilter, func(node ast.Node) {
		foundRecover := false
		switch n := node.(type) {
		case *ast.GoStmt:
			switch fn := n.Call.Fun.(type) {
			case *ast.FuncLit:
				ast.Inspect(fn, func(node ast.Node) bool {
					switch n := node.(type) {
					case *ast.Ident:
						if n.Name == "recover" {
							foundRecover = true
							return false
						}
					}
					return true
				})
			}
		}
		if !foundRecover {
			pass.Reportf(node.Pos(), errorMsg)
		}
	})
	return nil, nil
}
