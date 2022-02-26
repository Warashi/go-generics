package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Warashi/go-generics/zero"
)

func Unmarshal[T any](data []byte) (T, error) {
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return zero.New[T](), fmt.Errorf("json.Unmarshal: %w", err)
	}
	return value, nil
}

func NewDecoder[T any](r io.Reader) *Decoder[T] {
	return &Decoder[T]{Decoder: json.NewDecoder(r)}
}

type Decoder[T any] struct {
	*json.Decoder
}

func (d *Decoder[T]) Decode() (T, error) {
	var value T
	if err := d.Decoder.Decode(&value); err != nil {
		return zero.New[T](), fmt.Errorf("d.Decoder.Decode: %w", err)
	}
	return value, nil
}

func NewScanner[T any](r io.Reader) *Scanner[T] {
	return &Scanner[T]{Decoder: json.NewDecoder(r)}
}

type Scanner[T any] struct {
	*json.Decoder

	value T
	err   error
}

func (s *Scanner[T]) Scan() bool {
	s.err = s.Decode(&s.value)
	return s.err == nil
}

func (s *Scanner[T]) Value() T {
	return s.value
}

func (s *Scanner[T]) Err() error {
	return s.err
}
