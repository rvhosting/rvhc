package dataimg

import (
	"path/filepath"

	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
)

func Create(name string, size uint16) error {
	realPath := filepath.Join(baseDir, name+".img")
	realSize := int64(size) * 1024 * 1024
	img, err := diskfs.Create(realPath, realSize, diskfs.SectorSizeDefault)
	if err != nil {
		return err
	}

	spec := disk.FilesystemSpec{
		Partition: 0,
		FSType:    filesystem.TypeExt4,
	}

	if _, err := img.CreateFilesystem(spec); err != nil {
		return err
	}

	return nil
}
