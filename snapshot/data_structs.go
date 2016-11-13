package imagemake

import (
	"os"
	"strings"
	"time"

	"github.com/grantseltzer/fs-snapshot/dirtree"
)

func basisDate() time.Time {
	return time.Date(1995, time.May, 15, 0, 0, 0, 0, time.Local)
}

func insert(root *dirtree.Directory, FullPath string, f os.FileInfo) error {

	splitPath := strings.Split(FullPath, "/")

	// If we're in the correct dirtree.Directory
	if len(splitPath) == 1 || len(splitPath) == 2 {
		if f.IsDir() {
			var newDir = new(dirtree.Directory)
			newDir.DirectoryPath = f.Name()
			root.SubDirectories = append(root.SubDirectories, newDir)
		} else {
			var newEntry = new(dirtree.FileInfo)
			newEntry.FullPath = FullPath
			newEntry.Content = "$"
			newEntry.PermissionNumber = f.Mode()
			newEntry.LastAccessTime = basisDate()
			root.Entries = append(root.Entries, newEntry)
		}
	} else {
		// Step to the correct dirtree.Directory, recursively call insert
		for _, subDir := range root.SubDirectories {
			if subDir.DirectoryPath == splitPath[1] {
				err := insert(subDir, strings.Join(splitPath[1:], "/"), f)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
