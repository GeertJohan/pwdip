package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get working directory: %v\n", err)
		os.Exit(1)
	}

	pkg, err := build.ImportDir(wd, build.ImportComment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not find import path: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(pkg.ImportPath)
}
