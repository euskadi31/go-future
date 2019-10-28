// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package future

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFutureError(t *testing.T) {
	f := New()

	f.Error(errors.New("fail"))

	_, err := f.Get()
	assert.Error(t, err)
}

func TestFutureValue(t *testing.T) {
	f := New()

	f.Value("foo")

	v, err := f.Get()
	assert.NoError(t, err)
	assert.Equal(t, "foo", v)
}

func TestFutureFill(t *testing.T) {
	f := New()

	f.Value("foo")

	var v string

	err := f.Fill(&v)
	assert.NoError(t, err)
	assert.Equal(t, "foo", v)
}

func TestFutureFillWithError(t *testing.T) {
	f := New()

	f.Error(errors.New("fail"))

	var v string

	err := f.Fill(&v)
	assert.Error(t, err)
	assert.Equal(t, "", v)
}

func TestFutureFillWithBadDest(t *testing.T) {
	f := New()

	f.Value("foo")

	var v string

	err := f.Fill(v)
	assert.EqualError(t, err, "the fill destination should be a pointer to a `string`, but you used a `string`")
	assert.Equal(t, "", v)
}
