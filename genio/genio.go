// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2023 Shinnosuke Sawada. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is copied from io package of go1.21.0 and modified by Shinnosuke Sawada.

// Package genio provides basic interfaces to I/O primitives.
// Its primary job is to wrap existing implementations of such primitives,
// such as those in package os, into shared public interfaces that
// abstract the functionality, plus some other related primitives.
//
// Because these interfaces and primitives wrap lower-level operations with
// various implementations, unless otherwise informed clients should not
// assume they are safe for parallel execution.
package genio

import (
	"errors"
	"io"
)

// ErrInvalidWrite means that a write returned an impossible count.
var ErrInvalidWrite = errors.New("invalid write result")

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
type Reader[T any] interface {
	Read(p []T) (n int, err error)
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer[T any] interface {
	Write(p []T) (n int, err error)
}

// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter[T any] interface {
	Reader[T]
	Writer[T]
}

// ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser[T any] interface {
	Reader[T]
	io.Closer
}

// WriteCloser is the interface that groups the basic Write and Close methods.
type WriteCloser[T any] interface {
	Writer[T]
	io.Closer
}

// ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.
type ReadWriteCloser[T any] interface {
	Reader[T]
	Writer[T]
	io.Closer
}

// ReadSeeker is the interface that groups the basic Read and Seek methods.
type ReadSeeker[T any] interface {
	Reader[T]
	io.Seeker
}

// ReadSeekCloser is the interface that groups the basic Read, Seek and Close
// methods.
type ReadSeekCloser[T any] interface {
	Reader[T]
	io.Seeker
  io.Closer
}

// WriteSeeker is the interface that groups the basic Write and Seek methods.
type WriteSeeker[T any] interface {
	Writer[T]
	io.Seeker
}

// ReadWriteSeeker is the interface that groups the basic Read, Write and Seek methods.
type ReadWriteSeeker[T any] interface {
	Reader[T]
	Writer[T]
	io.Seeker
}

// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOF or error.
// The return value n is the number of elements read.
// Any error except EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
type ReaderFrom[T any] interface {
	ReadFrom(r Reader[T]) (n int64, err error)
}

// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to w until there's no more data to write or
// when an error occurs. The return value n is the number of elements
// written. Any error encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
type WriterTo[T any] interface {
	WriteTo(w Writer[T]) (n int64, err error)
}

// ReaderAt is the interface that wraps the basic ReadAt method.
//
// ReadAt reads len(p) elements into p starting at offset off in the
// underlying input source. It returns the number of elements
// read (0 <= n <= len(p)) and any error encountered.
//
// When ReadAt returns n < len(p), it returns a non-nil error
// explaining why more elements were not returned. In this respect,
// ReadAt is stricter than Read.
//
// Even if ReadAt returns n < len(p), it may use all of p as scratch
// space during the call. If some data is available but not len(p) elements,
// ReadAt blocks until either all the data is available or an error occurs.
// In this respect ReadAt is different from Read.
//
// If the n = len(p) elements returned by ReadAt are at the end of the
// input source, ReadAt may return either err == EOF or err == nil.
//
// If ReadAt is reading from an input source with a seek offset,
// ReadAt should not affect nor be affected by the underlying
// seek offset.
//
// Clients of ReadAt can execute parallel ReadAt calls on the
// same input source.
//
// Implementations must not retain p.
type ReaderAt[T any] interface {
	ReadAt(p []T, off int64) (n int, err error)
}

// WriterAt is the interface that wraps the basic WriteAt method.
//
// WriteAt writes len(p) elements from p to the underlying data stream
// at offset off. It returns the number of elements written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// WriteAt must return a non-nil error if it returns n < len(p).
//
// If WriteAt is writing to a destination with a seek offset,
// WriteAt should not affect nor be affected by the underlying
// seek offset.
//
// Clients of WriteAt can execute parallel WriteAt calls on the same
// destination if the ranges do not overlap.
//
// Implementations must not retain p.
type WriterAt[T any] interface {
	WriteAt(p []T, off int64) (n int, err error)
}

// ElementReader is the interface that wraps the ReadElement method.
//
// ReadElement reads and returns the next element from the input or
// any error encountered. If ReadElement returns an error, no input
// element was consumed, and the returned element value is undefined.
//
// ReadElement provides an efficient interface for element-at-time
// processing. A Reader that does not implement  ElementReader
// can be wrapped using bufio.NewReader to add this method.
type ElementReader[T any] interface {
	ReadElement() (T, error)
}

// ElementScanner is the interface that adds the UnreadElement method to the
// basic ReadElement method.
//
// UnreadElement causes the next call to ReadElement to return the last element read.
// If the last operation was not a successful call to ReadElement, UnreadElement may
// return an error, unread the last element read (or the element prior to the
// last-unread element), or (in implementations that support the Seeker interface)
// seek to one element before the current offset.
type ElementScanner[T any] interface {
	ElementReader[T]
	UnreadElement() error
}

// ElementWriter is the interface that wraps the WriteElement method.
type ElementWriter[T any] interface {
	WriteElement(c T) error
}

// ReadAtLeast reads from r into buf until it has read at least min elements.
// It returns the number of elements copied and an error if fewer elements were read.
// The error is EOF only if no elements were read.
// If an EOF happens after reading fewer than min elements,
// ReadAtLeast returns ErrUnexpectedEOF.
// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
// On return, n >= min if and only if err == nil.
// If r returns an error having read at least min elements, the error is dropped.
func ReadAtLeast[T any](r Reader[T], buf []T, min int) (n int, err error) {
	if len(buf) < min {
		return 0, io.ErrShortBuffer
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return
}

// ReadFull reads exactly len(buf) elements from r into buf.
// It returns the number of elements copied and an error if fewer elements were read.
// The error is EOF only if no elements were read.
// If an EOF happens after reading some but not all the elements,
// ReadFull returns ErrUnexpectedEOF.
// On return, n == len(buf) if and only if err == nil.
// If r returns an error having read at least len(buf) elements, the error is dropped.
func ReadFull[T any](r Reader[T], buf []T) (n int, err error) {
	return ReadAtLeast(r, buf, len(buf))
}

// CopyN copies n elements (or until an error) from src to dst.
// It returns the number of elements copied and the earliest
// error encountered while copying.
// On return, written == n if and only if err == nil.
//
// If dst implements the ReaderFrom interface,
// the copy is implemented using it.
func CopyN[T any](dst Writer[T], src Reader[T], n int64) (written int64, err error) {
	written, err = Copy(dst, LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		// src stopped early; must have been EOF.
		err = io.EOF
	}
	return
}

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of elements
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
func Copy[T any](dst Writer[T], src Reader[T]) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}

// CopyBuffer is identical to Copy except that it stages through the
// provided buffer (if one is required) rather than allocating a
// temporary one. If buf is nil, one is allocated; otherwise if it has
// zero length, CopyBuffer panics.
//
// If either src implements WriterTo or dst implements ReaderFrom,
// buf will not be used to perform the copy.
func CopyBuffer[T any](dst Writer[T], src Reader[T], buf []T) (written int64, err error) {
	if buf != nil && len(buf) == 0 {
		panic("empty buffer in CopyBuffer")
	}
	return copyBuffer(dst, src, buf)
}

// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer[T any](dst Writer[T], src Reader[T], buf []T) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo[T]); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom[T]); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader[T]); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]T, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = ErrInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

// LimitReader returns a Reader that reads from r
// but stops with EOF after n elements.
// The underlying implementation is a *LimitedReader.
func LimitReader[T any](r Reader[T], n int64) Reader[T] { return &LimitedReader[T]{r, n} }

// A LimitedReader reads from R but limits the amount of
// data returned to just N elements. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
type LimitedReader[T any] struct {
	R Reader[T] // underlying reader
	N int64     // max elements remaining
}

func (l *LimitedReader[T]) Read(p []T) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}


// NewSectionReader returns a SectionReader that reads from r
// starting at offset off and stops with EOF after n bytes.
func NewSectionReader[T any](r ReaderAt[T], off int64, n int64) *SectionReader[T] {
	var remaining int64
	const maxint64 = 1<<63 - 1
	if off <= maxint64-n {
		remaining = n + off
	} else {
		// Overflow, with no way to return error.
		// Assume we can read up to an offset of 1<<63 - 1.
		remaining = maxint64
	}
	return &SectionReader[T]{r, off, off, remaining}
}

// SectionReader implements Read, Seek, and ReadAt on a section
// of an underlying ReaderAt.
type SectionReader[T any] struct {
	r     ReaderAt[T]
	base  int64
	off   int64
	limit int64
}

func (s *SectionReader[T]) Read(p []T) (n int, err error) {
	if s.off >= s.limit {
		return 0, io.EOF
	}
	if max := s.limit - s.off; int64(len(p)) > max {
		p = p[0:max]
	}
	n, err = s.r.ReadAt(p, s.off)
	s.off += int64(n)
	return
}

var errWhence = errors.New("Seek: invalid whence")
var errOffset = errors.New("Seek: invalid offset")

func (s *SectionReader[T]) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errWhence
	case io.SeekStart:
		offset += s.base
	case io.SeekCurrent:
		offset += s.off
	case io.SeekEnd:
		offset += s.limit
	}
	if offset < s.base {
		return 0, errOffset
	}
	s.off = offset
	return offset - s.base, nil
}

func (s *SectionReader[T]) ReadAt(p []T, off int64) (n int, err error) {
	if off < 0 || off >= s.limit-s.base {
		return 0, io.EOF
	}
	off += s.base
	if max := s.limit - off; int64(len(p)) > max {
		p = p[0:max]
		n, err = s.r.ReadAt(p, off)
		if err == nil {
			err = io.EOF
		}
		return n, err
	}
	return s.r.ReadAt(p, off)
}

// Size returns the size of the section in bytes.
func (s *SectionReader[T]) Size() int64 { return s.limit - s.base }

// An OffsetWriter maps writes at offset base to offset base+off in the underlying writer.
type OffsetWriter[T any] struct {
	w    WriterAt[T]
	base int64 // the original offset
	off  int64 // the current offset
}

// NewOffsetWriter returns an OffsetWriter that writes to w
// starting at offset off.
func NewOffsetWriter[T any](w WriterAt[T], off int64) *OffsetWriter[T] {
	return &OffsetWriter[T]{w, off, off}
}

func (o *OffsetWriter[T]) Write(p []T) (n int, err error) {
	n, err = o.w.WriteAt(p, o.off)
	o.off += int64(n)
	return
}

func (o *OffsetWriter[T]) WriteAt(p []T, off int64) (n int, err error) {
	if off < 0 {
		return 0, errOffset
	}

	off += o.base
	return o.w.WriteAt(p, off)
}

func (o *OffsetWriter[T]) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errWhence
	case io.SeekStart:
		offset += o.base
	case io.SeekCurrent:
		offset += o.off
	}
	if offset < o.base {
		return 0, errOffset
	}
	o.off = offset
	return offset - o.base, nil
}

// TeeReader returns a Reader that writes to w what it reads from r.
// All reads from r performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
func TeeReader[T any](r Reader[T], w Writer[T]) Reader[T] {
	return &teeReader[T]{r, w}
}

type teeReader[T any] struct {
	r Reader[T]
	w Writer[T]
}

func (t *teeReader[T]) Read(p []T) (n int, err error) {
	n, err = t.r.Read(p)
	if n > 0 {
		if n, err := t.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return
}

// Discard is a Writer on which all Write calls succeed
// without doing anything.
type Discard[T any] struct{}

func (Discard[T]) Write(p []T) (int, error) {
	return len(p), nil
}

// NopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
// If r implements WriterTo, the returned ReadCloser will implement WriterTo
// by forwarding calls to r.
func NopCloser[T any](r Reader[T]) ReadCloser[T] {
	if _, ok := r.(WriterTo[T]); ok {
		return nopCloserWriterTo[T]{r}
	}
	return nopCloser[T]{r}
}

type nopCloser[T any] struct {
	Reader[T]
}

func (nopCloser[T]) Close() error { return nil }

type nopCloserWriterTo[T any] struct {
	Reader[T]
}

func (nopCloserWriterTo[T]) Close() error { return nil }

func (c nopCloserWriterTo[T]) WriteTo(w Writer[T]) (n int64, err error) {
	return c.Reader.(WriterTo[T]).WriteTo(w)
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll[T any](r Reader[T]) ([]T, error) {
	b := make([]T, 0, 512)
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return b, err
		}

		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			var zero T
			b = append(b, zero)[:len(b)]
		}
	}
}
