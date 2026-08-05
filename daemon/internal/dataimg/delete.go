package dataimg

import (
	"os"
	"path/filepath"
)

func Delete(id string) error {
	realPath := filepath.Join(BaseDir, id+".img")
	return os.RemoveAll(realPath)
}
