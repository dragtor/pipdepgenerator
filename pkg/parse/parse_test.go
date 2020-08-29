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
		InputStr:       []string{"from some.Resource import something,somethinganother ,ok", "for i in range:\n"},
		ExpectedOutput: []string{"from some.Resource import something,somethinganother ,ok"},
	},
}

func TestFilterImportStatements(t *testing.T) {
	for _, tc := range TestCases3 {
		actualResult := FilterImportStatements(tc.InputStr)
		if !reflect.DeepEqual(tc.ExpectedOutput, actualResult) {
			t.Errorf("Failed")
		}
	}
}

type TestCase4 struct {
	Importstatement string
	ExpectedPkgName []string
}

var TC4 = []TestCase4{
	TestCase4{
		Importstatement: "from dotenv import load_dotenv",
		ExpectedPkgName: []string{"dotenv"},
	},
	TestCase4{
		Importstatement: "from os.path import dirname, abspath",
		ExpectedPkgName: []string{"os"},
	},
	TestCase4{
		Importstatement: "from django.contrib.auth.models import Group",
		ExpectedPkgName: []string{"django"},
	},
	TestCase4{
		Importstatement: "import os,random",
		ExpectedPkgName: []string{"os", "random"},
	},
}

func TestTokenizeImportedPackage(t *testing.T) {
	for _, tc := range TC4 {
		actualResult := TokenizeImportedPackage(tc.Importstatement)
		if !reflect.DeepEqual(actualResult, tc.ExpectedPkgName) {
			t.Errorf(fmt.Sprintf("Failed expected %+v actual %+v", tc.ExpectedPkgName, actualResult))
		}
	}
}
