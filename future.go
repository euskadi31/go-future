// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package future

import (
	"fmt"
	"reflect"
)

// Future struct
type Future struct {
	err    error
	value  chan interface{}
	closed bool
}

// New Future
func New() *Future {
	return &Future{
		value: make(chan interface{}, 1),
	}
}

// Error set error to Future
func (f *Future) Error(err error) *Future {
	f.err = err

	f.Close()

	return f
}

// Value set value to Future
func (f *Future) Value(value interface{}) *Future {
	if !f.closed {
		f.value <- value
	}

	return f
}

// Get result or error
func (f *Future) Get() (interface{}, error) {
	value := <-f.value

	f.Close()

	if f.err != nil {
		return nil, f.err
	}

	return value, nil
}

// Close future
func (f *Future) Close() {
	if !f.closed {
		f.closed = true

		close(f.value)
	}
}

// Fill dest var
func (f *Future) Fill(dest interface{}) error {
	value := <-f.value

	f.Close()

	if f.err != nil {
		return f.err
	}

	return fill(value, dest)
}

func fill(src, dest interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			d := reflect.TypeOf(dest)
			s := reflect.TypeOf(src)
			err = fmt.Errorf("the fill destination should be a pointer to a `%s`, but you used a `%s`", s, d)
		}
	}()

	reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(src))

	return err
}
