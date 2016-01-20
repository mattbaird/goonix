package goonix

import (
	"os"
	"syscall"
)

type Disk struct {
}

func (d *Disk) Space(path string) (uint64, error) {

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

// current user
func (d *Disk) HasWritePermission(path string) (bool, error) {
	_, err := d.Exists(path)
	if err != nil {
		return false, err
	}
	return syscall.Access(path, 2) == nil, nil
}

func (d *Disk) NamedUserHasWritePermission(user, path string) (bool, error) {
	// execute namei -l /usr/local
	// get results that should look like this:
	//f: /home/matthew
	//drwxr-xr-x root    root    /
	//drwxr-xr-x root    root    home
	//drwxr-xr-x matthew matthew matthew
	// read last line
	//
	return false, nil
}

// works for paths and files
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
