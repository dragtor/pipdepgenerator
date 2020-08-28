package pypi

type TestCase struct {
	packageName         string
	ExpectedPackageMeta ProjectMeta
}

var TestCases = []TestCase{
	TestCase{
		packageName:         "sampleproject",
		ExpectedPackageMeta: projectMetaTC1,
	},
}

var projectMetaTC1 = ProjectMeta{Info: Info{Name: "sampleproject", RequiresDist: []string{"peppercorn", "check-manifest ; extra == 'dev'", "coverage ; extra == 'test'"}, RequiresPython: ">=3.5, <4"}, Releases: map[string][]Release{"1.0": []Release{},
	"1.2.0": []Release{Release{PythonVersion: "2.7"}, Release{PythonVersion: "source"}},
	"1.3.0": []Release{Release{PythonVersion: "py2.py3"}, Release{PythonVersion: "source"}},
	"1.3.1": []Release{Release{PythonVersion: "py2.py3"}, Release{PythonVersion: "py3"}, Release{PythonVersion: "source"}},
	"2.0.0": []Release{Release{PythonVersion: "py3"}, Release{PythonVersion: "source"}}}}
