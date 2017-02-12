package main

import (
	"os"

	"github.com/grantseltzer/prism/snapshot"
)

func main() {
	arg := os.Args[0]
	if arg == "snap" {
		snapshot.BuildTree()
	} else if arg == "repopulate" {
		x, _ := snapshot.LoadSnapshot("snapTwo.json")
		snapshot.Repopulate(x.(snapshot.Directory))
	}
}
