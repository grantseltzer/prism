package repopulate

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/grantseltzer/prism/dirtree"
)

// LoadSnapshot loads the snapshot json file into memory
func LoadSnapshot(fileName string) (interface{}, error) {
	jsonString, err := ioutil.ReadFile(fileName)
	jsonBytes := []byte(jsonString)
	var root dirtree.Directory
	err = json.Unmarshal(jsonBytes, root)
	return jsonBytes, err
}

// Repopulate is calling CreateFullPath on each entry from root
func Repopulate(root dirtree.Directory) {
	for _, x := range root.Entries {
		CreateFullPath(*x)
	}
	for _, i := range root.SubDirectories {
		Repopulate(*i)
	}
}

// CreateFullPath takes a path and will build the whole thing
func CreateFullPath(fileInfo dirtree.FileInfo) {
	path := strings.Split(fileInfo.FullPath, "/")
	currentPath := path[0]
	for i := range path {

		if !exists(currentPath) {
			os.Mkdir(currentPath, fileInfo.PermissionNumber)
		}
		currentPath = strings.Join(path[0:i], "/")
		if currentPath == fileInfo.FullPath {
			ioutil.WriteFile(currentPath, []byte(fileInfo.Content), fileInfo.PermissionNumber)
		}
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
