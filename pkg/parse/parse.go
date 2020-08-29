package parse

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Failed to read file")
	}
	return string(data)
}

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read file")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var data []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		data = append(data, string(line))
	}
	return data
}

func FilterImportStatements(lines []string) []string {
	var libstmtList []string
	for _, line := range lines {
		regExpr := `([\s]*^from[\s]+([_\w]+)(.[_\w]+)*[\s]+)?import[\s]+[\w]+(.[\w]+)*[\s]*(as[\s][\w]+)?(,[\s]*[\w]+[\s]*(as[\s][\w]+)?)*$`
		validImportStmt := regexp.MustCompile(regExpr)
		if validImportStmt.MatchString(line) {
			libstmtList = append(libstmtList, line)
		}
	}
	return libstmtList
}

type Package string
type ImportStatmentAST struct {
	PackageList []Package
}

func ListPackagesInImportSyntax(classStmt string) []string {
	tokens := strings.Split(classStmt, ",")
	var pkgList []string
	for _, pkg := range tokens {
		p := strings.TrimSpace(pkg)
		pkgList = append(pkgList, p)
	}
	return pkgList
}

func ListPackagesInFromSytax(fromStmt string) []string {
	baseTokenList := strings.Split(fromStmt, ",")
	var pkgList []string
	for _, baseToken := range baseTokenList {
		basepkgnamewithsubmodule := strings.TrimSpace(baseToken)
		submoduleList := strings.Split(basepkgnamewithsubmodule, ".")
		pkgList = append(pkgList, submoduleList[0])
	}
	return pkgList
}

func TokenizeImportedPackage(importStatement string) []string {
	regExpr := `(from (?P<package>[, .\w]+))*import (?P<classes>[,. \w]+)`
	pathMetadata := regexp.MustCompile(regExpr)
	matches := pathMetadata.FindStringSubmatch(importStatement)
	names := pathMetadata.SubexpNames()
	var tmpMap = map[string]string{}
	for i, match := range matches {
		if i != 0 {
			tmpMap[names[i]] = match
		}
	}
	if packageList, ok := tmpMap["package"]; ok {
		if strings.TrimSpace(packageList) == "" {
			return ListPackagesInImportSyntax(tmpMap["classes"])
		}
		return ListPackagesInFromSytax(tmpMap["package"])
	}
	return []string{}
}
