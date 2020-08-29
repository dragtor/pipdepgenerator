package fsutils

import (
	"fmt"
	"testing"
)

type TestCase struct {
	rootpath         string
	includesuffix    []string
	excludedirprefix []string
}

var TestCases = []TestCase{
	TestCase{
		rootpath:         "../../test/",
		includesuffix:    []string{"txt"},
		excludedirprefix: []string{"."},
	},
}

func TestFilePathWalkDir(t *testing.T) {
	for _, tc := range TestCases {
		actualResult, _ := FilePathWalkDir(tc.rootpath, tc.includesuffix, tc.excludedirprefix)
		fmt.Println(actualResult)
	}
}
