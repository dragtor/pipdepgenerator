package parse

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func GetImportedLibrariesList(lines []string) []string {
	var libstmtList []string
	for _, line := range lines {
		regExpr := "([ ]*from[ ]+[a-zA-Z0-9]+[.a-zA-Z_0-9]*)*[ ]*import[ ]+([a-zA-Z_0-9]+[.[a-zA-Z_0-9]*)[ ]*([ ]*,[ ]*[a-zA-Z_0-9]+)*[ \n]*$"
		validImportStmt := regexp.MustCompile(regExpr)
		if validImportStmt.MatchString(line) {
			libstmtList = append(libstmtList, line)
		}
	}
	return libstmtList
}
