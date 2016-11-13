package dirtree

import (
	"os"
	"time"
)

// Directory is information for each directory
type Directory struct {
	DirectoryPath  string
	Entries        []*FileInfo
	SubDirectories []*Directory
}

// FileInfo is information for each file
type FileInfo struct {
	FullPath         string
	PermissionNumber os.FileMode
	LastAccessTime   time.Time
	Content          string
}
