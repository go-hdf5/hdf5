// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
)

// File represents an HDF5 file.
type File struct {
	f *os.File

	super Superblock
}

// Open opens the named HDF5 file for reading.
func Open(name string) (*File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	// FIXME(sbinet): the signature+superblock may be located at
	// file offsets: 0, 512, 1024, 2048, ...
	buf := make([]byte, 8)
	_, err = io.ReadFull(f, buf)
	if err != nil {
		f.Close()
		return nil, errors.Wrapf(err, "could not read HDF5 signature")
	}

	if !bytes.Equal(buf, Signature[:]) {
		f.Close()
		return nil, ErrNotHDF5File
	}

	super, err := decodeSuperblock(f)
	if err != nil {
		f.Close()
		return nil, errors.Wrapf(err, "could not decode superblock")
	}

	return &File{f: f, super: super}, nil
}

func (f *File) Close() error {
	return f.f.Close()
}
