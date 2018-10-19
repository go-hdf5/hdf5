// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"math"
	"reflect"
	"testing"
)

func TestOpen(t *testing.T) {
	f, err := Open("testdata/empty-v0.h5")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var want = new(SuperblockV0)
	want.version.FreeSpace = 0x0
	want.version.SymTable = 0x0
	want.version.SharedHeader = 0x0
	want.offsets = 0x8
	want.lengths = 0x8
	want.GroupLeafNode = 0x04
	want.GroupInternalNode = 0x10
	want.Flags = 0x0
	want.addr.base = 0x0
	want.addr.free = math.MaxUint64
	want.addr.eof = 0x320
	want.addr.drv = math.MaxUint64
	want.SymTable = 0x0

	if !reflect.DeepEqual(want, f.super) {
		t.Fatalf("invalid file:\ngot= %#v\nwant=%#v", f.super, want)
	}
}
