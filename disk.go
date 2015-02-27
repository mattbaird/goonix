package goonix

import (
	"os"
	"syscall"
)

type Disk struct {
}

func (d *Disk) DiskSpace(path string) (uint64, error) {

	exists, err := d.Exists(path)
	if err != nil {
		return 0, err
	}
	isDir, err := d.IsDirectory(path)
	if err != nil {
		return 0, err
	}
	if exists && isDir {
		var stat syscall.Statfs_t
		syscall.Statfs(path, &stat)
		// Available blocks * size per block = available space in bytes
		return stat.Bavail * uint64(stat.Bsize), nil
	}
	return uint64(0), nil
}

func (d *Disk) IsDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	return mode.IsDir(), err
}

func (d *Disk) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
