package cmd

import (
	"fmt"

	"github.com/dragtor/pipdepgenerator/pkg/fsutils"
	"github.com/dragtor/pipdepgenerator/pkg/parse"
)

func GenerateRequirementTxt(root string, reqtxtpath string) {
	packages := make(map[string]bool)
	filePaths, _ := fsutils.FilePathWalkDir(root, []string{"py"}, []string{"."})
	for _, path := range filePaths {
		contents := parse.ReadFileLineByLine(path)
		importStmtList := parse.FilterImportStatements(contents)
		for _, stmt := range importStmtList {
			pkgs := parse.TokenizeImportedPackage(stmt)
			for _, pkg := range pkgs {
				packages[pkg] = true
			}
		}
	}
	fmt.Println(packages)
}
