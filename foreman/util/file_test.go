// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FIXME : Disable TestGraphiteFeedAPI*() for Linux because of these tests timeout On CentOS
//go:build !linux
// +build !linux

package util

import (
	"os"
	"strings"
	"testing"
)

const (
	errorFileListNotFound     = "File (%s) not found"
	errorFileListBadExtention = "Invalid Extention (%s) != *.%s"
)

func TestFileListFiles(t *testing.T) {
	file := NewFileWithPath("./")

	// All files

	files, err := file.ListFiles()
	if err != nil {
		t.Error(err)
	}

	if 0 < len(files) {
		for _, file := range files {
			_, err := os.Stat(file.GetPath())
			if os.IsNotExist(err) {
				t.Error(err)
			}
		}
	} else {
		t.Errorf(errorFileListNotFound, "")
	}

	// *.go files

	ext := "go"
	files, err = file.ListFilesWithExtention(ext)
	if err != nil {
		t.Error(err)
	}

	if 0 < len(files) {
		for _, file := range files {
			path := file.Path
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				t.Error(err)
			}
			if !strings.HasSuffix(path, ext) {
				t.Errorf(errorFileListBadExtention, file.Path, ext)
			}
		}
	} else {
		t.Errorf(errorFileListNotFound, ext)
	}

	// *.go files

	ext = ".go"
	files, err = file.ListFilesWithExtention(ext)
	if err != nil {
		t.Error(err)
	}

	if 0 < len(files) {
		for _, file := range files {
			path := file.Path
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				t.Error(err)
			}
			if !strings.HasSuffix(path, ext) {
				t.Errorf(errorFileListBadExtention, file.Path, ext)
			}
		}
	} else {
		t.Errorf(errorFileListNotFound, ext)
	}

}
