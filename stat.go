// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// A StatSys is a file system with a Stat method.
type StatSys interface {
	System

	// Stat returns a FileInfo describing the file.
	// If there is an error, it should be of type *PathError.
	Stat(name string) (FileInfo, error)
}

// Stat returns a FileInfo describing the named file from the file system.
//
// If System implements StatFS, Stat calls fs.Stat.
// Otherwise, Stat opens the file to stat it.
func Stat(fs System, name string) (FileInfo, error) {
	if fs, ok := fs.(StatSys); ok {
		return fs.Stat(name)
	}

	file, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file.Stat()
}
