// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package future

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	f := Error(errors.New("fail"))

	_, err := f.Get()
	assert.Error(t, err)

}

func TestValue(t *testing.T) {
	f := Value("foo")

	v, err := f.Get()
	assert.NoError(t, err)
	assert.Equal(t, "foo", v)
}
