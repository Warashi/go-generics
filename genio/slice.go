// Copyright 2012 The Go Authors. All rights reserved.
// Copyright 2023 Shinnosuke Sawada. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is copied from bytes package of go1.21.0 and modified by Shinnosuke Sawada.
package genio

import (
	"errors"
	"io"
)

type SliceReader[T any] struct {
	s []T
	i int64
}

// Len returns the number of elements of the unread portion of the slice.
func (r *SliceReader[T]) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

// Size returns the original length of the underlying slice.
// Size is the number of elements available for reading via ReadAt.
// The result is unaffected by any method calls except Reset.
func (r *SliceReader[T]) Size() int64 { return int64(len(r.s)) }

// Read implements the genio.Reader interface.
func (r *SliceReader[T]) Read(b []T) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// ReadAt implements the genio.ReaderAt interface.
func (r *SliceReader[T]) ReadAt(b []T, off int64) (n int, err error) {
	// cannot modify state - see genio.ReaderAt
	if off < 0 {
		return 0, errors.New("genio.SliceReader.ReadAt: negative offset")
	}
	if off >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}

// ReadElement implements the genio.ElementReader interface.
func (r *SliceReader[T]) ReadElement() (T, error) {
	if r.i >= int64(len(r.s)) {
		var zero T
		return zero, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

// UnreadElement complements ReadElement in implementing the genio.ElementScanner interface.
func (r *SliceReader[T]) UnreadElement() error {
	if r.i <= 0 {
		return errors.New("genio.SliceReader.UnreadElement: at beginning of slice")
	}
	r.i--
	return nil
}

// Seek implements the io.Seeker interface.
func (r *SliceReader[T]) Seek(offset int64, whence int) (int64, error) {
	var abs int64
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = r.i + offset
	case io.SeekEnd:
		abs = int64(len(r.s)) + offset
	default:
		return 0, errors.New("genio.SliceReader.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("genio.SliceReader.Seek: negative position")
	}
	r.i = abs
	return abs, nil
}

// WriteTo implements the genio.WriterTo interface.
func (r *SliceReader[T]) WriteTo(w Writer[T]) (n int64, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, nil
	}
	b := r.s[r.i:]
	m, err := w.Write(b)
	if m > len(b) {
		panic("genio.SliceReader.WriteTo: invalid Write count")
	}
	r.i += int64(m)
	n = int64(m)
	if m != len(b) && err == nil {
		err = io.ErrShortWrite
	}
	return
}

// Reset resets the Reader to be reading from b.
func (r *SliceReader[T]) Reset(b []T) { *r = SliceReader[T]{b, 0} }

// NewSliceReader returns a new SliceReader reading from b.
func NewSliceReader[T any](b []T) *SliceReader[T] { return &SliceReader[T]{b, 0} }
