package osexit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var ExitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for os.Exit",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	expr := func(x *ast.SelectorExpr) {
		// проверяем, что выражение представляет собой вызов функции os.Exit,
		ident, ok := x.X.(*ast.Ident)
		if !ok {
			return
		}
		sel := x.Sel

		if ident.Name == "os" && sel.Name == "Exit" {
			pass.Reportf(x.Pos(), "in file finded os.Exit")
		}
	}

	for _, file := range pass.Files {
		// ast.Print(pass.Fset, file)
		// функцией ast.Inspect проходим по всем узлам AST
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.SelectorExpr: // выражение
				expr(x)
			}
			return true
		})

	}
	return nil, nil
}
