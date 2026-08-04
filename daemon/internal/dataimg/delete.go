package dataimg

import (
	"os"
	"path/filepath"
)

func Delete(name string) error {
	realPath := filepath.Join(baseDir, name+".img")
	return os.RemoveAll(realPath)
}
