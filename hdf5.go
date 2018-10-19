// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hdf5 reads and writes files in the HDF5 format
package hdf5

import (
	"github.com/pkg/errors"
)

//go:generate go run ./testdata/gen-create-empty-file.go

var (
	// Signature identifies a file as being an HDF5 file.
	Signature = [8]byte{0x89, 'H', 'D', 'F', '\r', '\n', 0x1a, '\n'}

	ErrNotHDF5File          = errors.New("hdf5: not a HDF5 file")
	ErrBadSuperblockVersion = errors.New("hdf5: bad superblock version number")
)
