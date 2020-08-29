package parse

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	InputFilePath   string
	ExpectedContent string
}

var TestCases = []TestCase{
	TestCase{
		InputFilePath:   "../../test/data/samplefile1.txt",
		ExpectedContent: "hello shubham\nimport flask",
	},
}

func TestReadFile(t *testing.T) {
	for _, tc := range TestCases {
		actualData := ReadFile(tc.InputFilePath)
		if actualData != tc.ExpectedContent {
			t.Errorf(fmt.Sprintf("mismatch contents"))
		}
	}
}

type TestCase2 struct {
	Inputcontent   string
	Expectedoutput []string
	IsValid        bool
}

func TestReadFileLineByLine(t *testing.T) {
	actualResult := ReadFileLineByLine("../../test/data/samplefile1.txt")
	fmt.Println(len(actualResult))
}

type TestCase3 struct {
	InputStr       []string
	ExpectedOutput []string
}

var TestCases3 = []TestCase3{
	TestCase3{
		InputStr:       []string{"from some.Resource import something,somethinganother ,ok"},
		ExpectedOutput: []string{"from some.Resource import something,somethinganother ,ok"},
	},
}

func TestGetImportedLibrariesList(t *testing.T) {
	for _, tc := range TestCases3 {
		actualResult := GetImportedLibrariesList(tc.InputStr)
		if !reflect.DeepEqual(tc.ExpectedOutput, actualResult) {
			t.Errorf("Failed")
		}
	}
}
