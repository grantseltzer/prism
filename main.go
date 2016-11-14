package main

import (
	"os"

	"github.com/grantseltzer/prism/dirtree"
	"github.com/grantseltzer/prism/repopulate"
	"github.com/grantseltzer/prism/snapshot"
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
