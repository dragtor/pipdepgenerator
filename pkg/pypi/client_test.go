package pypi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFetchProjectMetaData(t *testing.T) {
	for _, tc := range TestCases {
		actualPackageMeta, err := FetchProjectMetaData(tc.packageName)
		if err != nil {
			t.Fatalf(fmt.Sprintf(err.Error()))
		}
		if !reflect.DeepEqual(*actualPackageMeta, tc.ExpectedPackageMeta) {
			t.Fatal("package meta mismatch")
		}
		t.Logf("project API is working")
	}
}
