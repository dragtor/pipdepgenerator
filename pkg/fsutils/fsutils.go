package fsutils

import (
	"os"
	"path/filepath"
	"strings"
)

func FilePathWalkDir(root string, filesuffix []string, excludeDirPrefixList []string) ([]string, error) {
	var files []string
	var extensionMap = map[string]bool{}
	for _, ext := range filesuffix {
		extensionMap[ext] = true
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			basepath := filepath.Base(path)
			var err error
			for _, pfx := range excludeDirPrefixList {
				if strings.HasPrefix(basepath, pfx) {
					err = filepath.SkipDir
				}
			}
			if err != nil {
				return err
			}
		}
		if !info.IsDir() {
			basepath := filepath.Base(path)
			splitedbasepath := strings.Split(basepath, ".")
			extension := splitedbasepath[len(splitedbasepath)-1]
			if _, ok := extensionMap[extension]; ok {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}
