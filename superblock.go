// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

type SuperblockV0 struct {
	Signature [8]byte
	Version   struct {
		Superblock  byte
		FreeSpace   byte
		SymbolTable byte
	}
	_                 byte
	SharedHeader      byte
	Offsets           byte
	Lengths           byte
	_                 byte
	GroupLeafNode     uint16
	GroupInternalNode uint16
	Flags             uint32
	Address           struct {
		Base   uint64
		Free   uint64
		EOF    uint64
		Driver uint64
	}
	SymbolTable uint64
}
