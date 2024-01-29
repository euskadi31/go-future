// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package future

// Error return future error.
func Error(err error) *Future {
	return New().Error(err)
}

// Value return future value.
func Value(value interface{}) *Future {
	return New().Value(value)
}
