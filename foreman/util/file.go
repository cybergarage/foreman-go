// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File represents a file or director.
type File struct {
	Path string
}

const (
	errorFileIsNotDirectory = "%s is not a directory"
)

// NewFileWithPath return a file with the specified path.
func NewFileWithPath(path string) *File {
	file := &File{
		Path: path,
	}
	return file
}

// GetPath returns the path.
func (file *File) GetPath() string {
	return file.Path
}

// Ext returns only the extention.
func (file *File) Ext() string {
	return filepath.Ext(file.Path)
}

// IsDir returns true when the file is a directory, otherwise false
func (file *File) IsDir() bool {
	fi, err := os.Stat(file.Path)
	if err != nil {
		return false
	}
	if !fi.IsDir() {
		return false
	}
	return true
}

// ListFilesWithExtention returns files which has the specified extentions in the directory.
func (file *File) ListFilesWithExtention(targetExt string) ([]*File, error) {
	if !file.IsDir() {
		return nil, fmt.Errorf(errorFileIsNotDirectory, file.Path)
	}

	rootPath := file.Path
	files := []*File{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if 0 < len(targetExt) {
			fileExt := filepath.Ext(path)
			if !strings.HasSuffix(fileExt, targetExt) {
				return nil
			}
		}
		files = append(files, NewFileWithPath(path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// ListFiles returns a file list in the directory.
func (file *File) ListFiles() ([]*File, error) {
	return file.ListFilesWithExtention("")
}
