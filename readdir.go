// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"errors"
	"sort"
)

// ReadDirFS is the interface implemented by a file system
// that provides an optimized implementation of ReadDir.
type ReadDirFS interface {
	System

	// ReadDir reads the named directory
	// and returns a list of directory entries sorted by filename.
	ReadDir(name string) ([]FileInfo, error)
}

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
//
// If System implements ReadDirFS, ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func ReadDir(fs System, name string) ([]FileInfo, error) {
	if fs, ok := fs.(ReadDirFS); ok {
		return fs.ReadDir(name)
	}

	file, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// TODO: Do we really need the Stat?
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("TODO")
	}

	dir, ok := file.(ReadDirFile)
	if !ok {
		return nil, errors.New("TODO")
	}

	list, err := dir.ReadDir(-1)
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, err
}
