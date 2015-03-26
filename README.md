## pwdip
**Print Working Directory Import Path** is a tool that prints the Go import path based on the working directory.

#### Usage
`pwdip` has no flags or arguments (yet).

When the working directory is not located in the GOPATH, `pwdip` returns an error.

`pwdip` does not validate the package, it only finds the import path.

#### Example
[rerun](github.com/skelterjohn/rerun) is a simple tool that watches a package and its dependencies, and rebuilds (+runs) it when .go files are changed.

It's used as: `rerun <import path>`.

Normally it would be called like this:
```
rerun github.com/GeertJohan/gomatrix
```

Even if you're already in the gomatrix directory itself. `pwdip` simplifies this to:
```
rerun `pwdip`
```
