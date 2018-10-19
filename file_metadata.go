// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type Superblock interface {
	Version() byte
	Offsets() byte
	Lengths() byte

	decode(r io.Reader) error
}

type SuperblockV0 struct {
	version struct {
		FreeSpace    byte
		SymTable     byte
		_            byte
		SharedHeader byte
	}
	offsets           byte
	lengths           byte
	_                 byte
	GroupLeafNode     uint16
	GroupInternalNode uint16
	Flags             uint32
	addr              struct {
		base uint64
		free uint64
		eof  uint64
		drv  uint64
	}
	SymTable uint32
}

func (*SuperblockV0) Version() byte      { return 0 }
func (sblk *SuperblockV0) Offsets() byte { return sblk.offsets }
func (sblk *SuperblockV0) Lengths() byte { return sblk.lengths }

func (s *SuperblockV0) decode(r io.Reader) error {
	buf := make([]byte, 15)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return err
	}

	s.version.FreeSpace = buf[0]
	s.version.SymTable = buf[1]
	s.version.SharedHeader = buf[3]
	s.offsets = buf[4]
	s.lengths = buf[5]
	_ = buf[6]
	s.GroupLeafNode = binary.LittleEndian.Uint16(buf[7 : 7+2])
	s.GroupInternalNode = binary.LittleEndian.Uint16(buf[9 : 9+2])
	s.Flags = binary.LittleEndian.Uint32(buf[11:15])

	rr := newRBuffer(r, s.offsets, s.lengths)
	s.addr.base = rr.readOffset()
	s.addr.free = rr.readOffset()
	s.addr.eof = rr.readOffset()
	s.addr.drv = rr.readOffset()
	s.SymTable = rr.readU32()

	return rr.err
}

type SuperblockV1 struct {
	version struct {
		FreeSpace byte
		SymTable  byte
	}
	_                 byte
	SharedHeader      byte
	offsets           byte
	lengths           byte
	_                 byte
	GroupLeafNode     uint16
	GroupInternalNode uint16
	flags             uint32
	IndexedStorage    uint16
	_                 uint16
	addr              struct {
		base uint64
		free uint64
		eof  uint64
		drv  uint64
	}
	SymTable uint32
}

func (*SuperblockV1) Version() byte      { return 1 }
func (sblk *SuperblockV1) Offsets() byte { return sblk.offsets }
func (sblk *SuperblockV1) Lengths() byte { return sblk.lengths }

func (s *SuperblockV1) decode(r io.Reader) error {
	panic("not implemented")
}

type SuperblockV2 struct {
	offsets byte
	lengths byte
	flags   uint32
	addr    struct {
		base      uint64
		super     uint64
		eof       uint64
		rootGroup uint64
	}
	chksum uint32
}

func (*SuperblockV2) Version() byte      { return 2 }
func (sblk *SuperblockV2) Offsets() byte { return sblk.offsets }
func (sblk *SuperblockV2) Lengths() byte { return sblk.lengths }

func (s *SuperblockV2) decode(r io.Reader) error {
	panic("not implemented")
}

type SuperblockV3 struct {
	offsets byte
	lengths byte
	flags   uint32
	addr    struct {
		base      uint64
		super     uint64
		eof       uint64
		rootGroup uint64
	}
	chksum uint32
}

func (*SuperblockV3) Version() byte      { return 3 }
func (sblk *SuperblockV3) Offsets() byte { return sblk.offsets }
func (sblk *SuperblockV3) Lengths() byte { return sblk.lengths }

func (s *SuperblockV3) decode(r io.Reader) error {
	panic("not implemented")
}

func decodeSuperblock(r io.Reader) (Superblock, error) {
	var buf [8]byte
	_, err := io.ReadFull(r, buf[:1])
	if err != nil {
		return nil, errors.Wrapf(err, "could not read superblock version")
	}

	var super Superblock
	switch buf[0] {
	case 0:
		super = &SuperblockV0{}
	case 1:
		super = &SuperblockV1{}
	case 2:
		super = &SuperblockV2{}
	case 3:
		super = &SuperblockV3{}
	default:
		return nil, ErrBadSuperblockVersion
	}

	err = super.decode(r)
	if err != nil {
		return nil, err
	}

	return super, nil
}

var (
	_ Superblock = (*SuperblockV0)(nil)
	_ Superblock = (*SuperblockV1)(nil)
	_ Superblock = (*SuperblockV2)(nil)
	_ Superblock = (*SuperblockV3)(nil)
)
