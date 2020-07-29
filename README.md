# io/fs naming

This repository is based on the draft code in https://go.googlesource.com/go/+/refs/changes/16/243916/1/src/io
refactored for an alternate naming emphasizing the main components of the package
the File and the System.

In short

- rename type FS to System
- rename suffix FS to Sys
- rename argument fsys to fs

Readability is improved in how we reason/talk about things, eg.

- Before: "Glob uses the fsys argument of type FS to ..."
- After: "Glob uses the fs argument of type System to ..."

- Before: "ReadFileFS interface extends the FS interface with ..."
- After: "ReadFileSys interface extends the System interface with ..."

As the System is one of the main interfaces in this package it
shouldn't be abbreviated. However in the extension interfaces(good
name for it), abbreviation is ok as the emphasis is on the extension
func.


## Before

Filtering out the unaffected parts

```go
func Glob(fsys FS, pattern string) (matches []string, err error)
func ReadFile(fsys FS, name string) ([]byte, error)
func Walk(fsys FS, root string, walkFn WalkFunc) error
type FS interface{ ... }
    func ReadDir(fsys FS, name string) ([]FileInfo, error)
    func Stat(fsys FS, name string) (FileInfo, error)
type GlobFS interface{ ... }
type ReadDirFS interface{ ... }
type ReadFileFS interface{ ... }
type StatFS interface{ ... }
```

## After

```go
func Glob(fs System, pattern string) (matches []string, err error)
func ReadFile(fs System, name string) ([]byte, error)
func Walk(fs System, root string, walkFn WalkFunc) error
    func ReadDir(fs System, name string) ([]FileInfo, error)
    func Stat(fs System, name string) (FileInfo, error)
type GlobSys interface{ ... }
type ReadDirSys interface{ ... }
type ReadFileSys interface{ ... }
type StatSys interface{ ... }
type System interface{ ... }
```

Using the renamed variation

```go
func DoSomthing(sys fs.System) error {
    sys.Open("myfile.txt")
}
```
