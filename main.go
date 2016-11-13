package main

import (
	"os"

	"github.com/grantseltzer/fs-snapshot/dirtree"
	"github.com/grantseltzer/fs-snapshot/repopulate"
	"github.com/grantseltzer/fs-snapshot/snapshot"
)

func main() {
	arg := os.Args[0]
	if arg == "snap" {
		imagemake.BuildTree()
	} else if arg == "repopulate" {
		x, _ := repopulate.LoadSnapshot("snapTwo.json")
		repopulate.Repopulate(x.(dirtree.Directory))
	}
}
