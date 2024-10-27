package meta

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/wutipong/mangaweb3-backend/configuration"
	"github.com/wutipong/mangaweb3-backend/container"
)

// ListDir returns a list of content of a directory.
func ListDir(path string) (files []string, err error) {

	c := configuration.Get()
	actualPath := filepath.Join(c.DataPath, path)
	dir, err := os.Open(actualPath)
	if err != nil {
		return
	}
	children, err := dir.Readdir(0)
	if err != nil {
		return
	}

	for _, f := range children {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		name := filepath.Join(path, f.Name())
		if f.IsDir() {
			subFiles, e := ListDir(name)
			if e != nil {
				err = e
				return
			}
			files = append(files, subFiles...)
		}

		_, valid := container.GuessContainerType(context.Background(), name, f)
		if !valid {
			continue
		}

		files = append(files, name)
	}
	return
}
