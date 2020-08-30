package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dragtor/pipdepgenerator/pkg/fsutils"
	"github.com/dragtor/pipdepgenerator/pkg/parse"
	"github.com/dragtor/pipdepgenerator/pkg/pypi"
)

type Dependencies struct {
	packages map[string]PackageDetails
}
type PackageDetails struct {
	Name          string
	LatestVersion string
	// ReleaseList   map[string][]pypi.Release
}

func FindDependenciesForProject(root string) Dependencies {
	packages := make(map[string]bool)
	var dep Dependencies
	dep.packages = make(map[string]PackageDetails)
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
	for k, _ := range packages {
		if _, got := dep.packages[k]; !got {
			log.Println(fmt.Sprintf("Fetching Info for package %s", k))
			projectInfo, err := pypi.FetchProjectMetaData(k)
			if err != nil {
				log.Println(fmt.Sprintf("%s %s", err.Error(), k))
				continue
			}
			var pkg PackageDetails
			pkg.Name = projectInfo.Info.Name
			pkg.LatestVersion = projectInfo.Info.Version
			dep.packages[k] = pkg
		}
	}
	return dep
}

func WriteToRequirementFile(dep Dependencies, filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file path : %s", filepath))
	}
	defer f.Close()
	for _, pkg := range dep.packages {
		f.WriteString(fmt.Sprintf("%s == %s\n", pkg.Name, pkg.LatestVersion))
	}
}
func GenerateRequirementTxtFile(root, filepath string) {
	dep := FindDependenciesForProject(root)
	WriteToRequirementFile(dep, filepath)
}
