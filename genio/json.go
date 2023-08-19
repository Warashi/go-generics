package genio

import (
	"encoding/json"
	"io"
)

// NewJSONDecoder returns a new JSONDecoder that reads from r.
func NewJSONDecoder[T any](r io.Reader) *JSONDecoder[T] {
	return &JSONDecoder[T]{decoder: json.NewDecoder(r)}
}

// JSONDecoder is a wrapper around json.Decoder that implements the Reader interface.
type JSONDecoder[T any] struct {
	decoder *json.Decoder
}

// Read implements the Reader interface.
func (r *JSONDecoder[T]) Read(p []T) (n int, err error) {
	for i := range p {
		if err := r.decoder.Decode(&p[i]); err != nil {
			return i, err
		}
	}
	return len(p), nil
}

// NewJSONEncoder returns a new JSONEncoder that reads from r.
func NewJSONEncoder[T any](r Reader[T]) *JSONEncoder[T] {
	return &JSONEncoder[T]{
		r:   r,
		buf: make([]byte, 0, 8192),
	}
}

// JSONEncoder is a wrapper around json.Encoder that implements the io.Reader interface.
type JSONEncoder[T any] struct {
	r   Reader[T]
	buf []byte
}

// Read implements the io.Reader interface.
func (r *JSONEncoder[T]) Read(p []byte) (n int, err error) {
  var b [1]T
  nn, err := r.r.Read(b[:])
	if err != nil {
		return 0, err
	}
	if nn == 0 {
		return 0, nil
	}

	for i := range b {
		b, err := json.Marshal(b[i])
		if err != nil {
			return 0, err
		}
		r.buf = append(r.buf, b...)
	}

  n = copy(p, r.buf)
  r.buf = r.buf[n:]

  return n, nil
}
