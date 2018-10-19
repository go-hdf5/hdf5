// Copyright 2018 The go-hdf5 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type rbuffer struct {
	r   io.Reader
	err error
	buf []byte

	sz struct {
		offset func() uint64
		length func() uint64
	}
}

func newRBuffer(r io.Reader, offset, length byte) *rbuffer {
	rbuf := &rbuffer{r: r, buf: make([]byte, 8)}
	switch offset {
	case 4:
		rbuf.sz.offset = func() uint64 { return uint64(rbuf.readU32()) }
	case 8:
		rbuf.sz.offset = rbuf.readU64
	default:
		panic(errors.Errorf("hdf5: invalid offset size (%v)", offset))
	}

	switch length {
	case 4:
		rbuf.sz.length = func() uint64 { return uint64(rbuf.readU32()) }
	case 8:
		rbuf.sz.length = rbuf.readU64
	default:
		panic(errors.Errorf("hdf5: invalid length size (%v)", length))
	}

	return rbuf
}

func (r *rbuffer) Read(p []byte) (int, error) {
	if r.err != nil {
		return 0, r.err
	}
	var n int
	n, r.err = r.r.Read(p)
	return n, r.err
}

func (r *rbuffer) readOffset() uint64 {
	return r.sz.offset()
}

func (r *rbuffer) readLen() uint64 {
	return r.sz.length()
}

func (r *rbuffer) readU16() uint16 {
	const n = 2
	r.load(n)
	return binary.LittleEndian.Uint16(r.buf[:n])
}

func (r *rbuffer) readU32() uint32 {
	const n = 4
	r.load(n)
	return binary.LittleEndian.Uint32(r.buf[:n])
}

func (r *rbuffer) readU64() uint64 {
	const n = 8
	r.load(n)
	return binary.LittleEndian.Uint64(r.buf[:n])
}

func (r *rbuffer) load(n int) {
	if r.err != nil {
		return
	}
	_, r.err = io.ReadFull(r, r.buf[:n])
}

var (
	_ io.Reader = (*rbuffer)(nil)
)
