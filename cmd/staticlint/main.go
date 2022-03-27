package main

import (
	"staticlint/internal/osexit"

	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
)

var checks = map[string]bool{
	"S1000": true,
	"S1001": true,
}

func addAllSAToChecks() {
	for _, v := range staticcheck.Analyzers {
		// добавляем в массив нужные проверки
		if v.Analyzer.Name[:2] == "SA" {
			checks[v.Analyzer.Name] = true
		}
	}
}

func main() {
	addAllSAToChecks()
	var mychecks []*analysis.Analyzer
	for _, v := range staticcheck.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	for _, v := range simple.Analyzers {
		if checks[v.Analyzer.Name] {
			mychecks = append(mychecks, v.Analyzer)
		}
	}
	mychecks = append(mychecks, osexit.ExitCheckAnalyzer)
	mychecks = append(mychecks, bodyclose.Analyzer)
	mychecks = append(mychecks, printf.Analyzer)
	mychecks = append(mychecks, shadow.Analyzer)
	mychecks = append(mychecks, structtag.Analyzer)
	multichecker.Main(
		mychecks...,
	)
}
