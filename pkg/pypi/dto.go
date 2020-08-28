package pypi

type ProjectMeta struct {
	Info     Info                 `json:"info"`
	Releases map[string][]Release `json:"releases"`
}
type Info struct {
	Name           string   `json:"name"`
	RequiresDist   []string `json:"requires_dist"`
	RequiresPython string   `json:"requires_python"`
}

type Release struct {
	PythonVersion string `json:"python_version"`
}
