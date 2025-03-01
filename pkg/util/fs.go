package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/otiai10/copy"
	"github.com/weaveworks/ignite/pkg/constants"
)

// Creates the /var/lib/firecracker/{vm,image,kernel} directories
func CreateDirectories() error {
	for _, dir := range []string{constants.VM_DIR, constants.IMAGE_DIR, constants.KERNEL_DIR, constants.MANIFEST_DIR} {
		if err := os.MkdirAll(dir, constants.DATA_DIR_PERM); err != nil {
			return fmt.Errorf("failed to create directory %q: %v", dir, err)
		}
	}

	return nil
}

func PathExists(path string) (bool, os.FileInfo) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, info
}

func FileExists(filename string) bool {
	exists, info := PathExists(filename)
	if !exists {
		return false
	}

	return !info.IsDir()
}

func DirExists(dirname string) bool {
	exists, info := PathExists(dirname)
	if !exists {
		return false
	}

	return info.IsDir()
}

// CopyFile copies both files and directories
func CopyFile(src string, dst string) error {
	return copy.Copy(src, dst)
}

type MountPoint struct {
	Path string
}

func Mount(volume string) (*MountPoint, error) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}

	if _, err := ExecuteCommand("mount", volume, tempDir); err != nil {
		return nil, fmt.Errorf("failed to mount volume %q: %v", volume, err)
	}

	return &MountPoint{
		Path: tempDir,
	}, nil
}

func (mp *MountPoint) Umount() error {
	if _, err := ExecuteCommand("umount", mp.Path); err != nil {
		return fmt.Errorf("failed to unmount volume %q: %v", mp.Path, err)
	}

	if err := os.RemoveAll(mp.Path); err != nil {
		return err
	}

	return nil
}

// FileIsEmpty returns true if the file is empty
func FileIsEmpty(file string) (bool, error) {
	fileInfo, err := os.Stat(file)
	// Check if there was an unexpected error
	if err != nil && !os.IsNotExist(err) {
		return false, err
	}

	// The file exists, and has content. Proceed as usual
	if err == nil && fileInfo.Size() > 0 {
		return false, nil
	}

	// The file exists, but has no content. Remove the file to allow the symlink
	if err == nil && fileInfo.Size() == 0 {
		if err := os.Remove(file); err != nil {
			return false, err
		}
	}

	return true, nil
}
