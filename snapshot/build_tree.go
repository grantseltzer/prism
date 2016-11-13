package imagemake

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grantseltzer/fs-snapshot/dirtree"
)

var root dirtree.Directory

// AddToTree is called on each fs object to add to the directory tree
func AddToTree(path string, f os.FileInfo, err error) error {
	if !strings.HasPrefix("/proc", path) && !strings.HasPrefix("/var/lib/oci-storage", path) {
		insert(&root, path, f)
	}
	return nil
}

// BuildTree calls AddToTree on each fs object
func BuildTree() error {
	err := filepath.Walk("/", AddToTree)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(&root, "", "    ")
	fmt.Println(string(b))
	return nil
}
