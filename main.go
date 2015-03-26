package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pkg, err := build.ImportDir(wd, build.FindOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: can not find import path: %v\n", err)
		os.Exit(1)
	}

	if pkg.ImportPath == "." {
		fmt.Fprint(os.Stderr, "error: package is not located in GOPATH\n")
		os.Exit(1)
	}

	fmt.Print(pkg.ImportPath)
}
