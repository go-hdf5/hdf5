// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"

	"gonum.org/v1/hdf5"
)

func main() {
	createEmptyFile()
	createDataset_2x5()
}

func createEmptyFile() {
	const fname = "testdata/empty.h5"
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func createDataset_2x5() {
	const fname = "testdata/dataset.h5"

	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		log.Fatalf("CreateFile failed: %s", err)
	}
	defer f.Close()

	dims := []uint{2, 5}
	dspace, err := hdf5.CreateSimpleDataspace(dims, dims)
	if err != nil {
		log.Fatal(err)
	}
	defer dspace.Close()

	dset, err := f.CreateDataset("dset", hdf5.T_NATIVE_USHORT, dspace)
	if err != nil {
		log.Fatal(err)
	}
	defer dset.Close()

	data := [10]uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	err = dset.Write(&data[0])
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
