package pypi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchProjectMetaData(projectName string) (*ProjectMeta, error) {
	path := fmt.Sprintf("https://pypi.org/pypi/%s/json", projectName)
	resp, err := http.Get(path)
	if err != nil {
		return nil, errors.New("Failed to fetch")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to fetch ")
	}
	body, err := ioutil.ReadAll(resp.Body)
	var projectMeta ProjectMeta
	err = json.Unmarshal(body, &projectMeta)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to unmarshal data : %s", err.Error()))
	}
	return &projectMeta, nil
}
