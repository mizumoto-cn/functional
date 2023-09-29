// will change to built-in cmp after go 1.21 release

// Copyright (c) 2009 The Go Authors. All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cmp_test

import (
	"math"
	"sort"
	"testing"

	"github.com/mizumoto-cn/functional/pkg/cmp"
)

var negzero = math.Copysign(0, -1)
var tests = []struct {
	x, y    any
	compare int
}{
	{1, 2, -1},
	{1, 1, 0},
	{2, 1, +1},
	{"a", "aa", -1},
	{"a", "a", 0},
	{"aa", "a", +1},
	{1.0, 1.1, -1},
	{1.1, 1.1, 0},
	{1.1, 1.0, +1},
	{math.Inf(1), math.Inf(1), 0},
	{math.Inf(-1), math.Inf(-1), 0},
	{math.Inf(-1), 1.0, -1},
	{1.0, math.Inf(-1), +1},
	{math.Inf(1), 1.0, +1},
	{1.0, math.Inf(1), -1},
	{math.NaN(), math.NaN(), 0},
	{0.0, math.NaN(), +1},
	{math.NaN(), 0.0, -1},
	{math.NaN(), math.Inf(-1), -1},
	{math.Inf(-1), math.NaN(), +1},
	{0.0, 0.0, 0},
	{negzero, negzero, 0},
	{negzero, 0.0, 0},
	{0.0, negzero, 0},
	{negzero, 1.0, -1},
	{negzero, -1.0, +1},
}

func TestLess(t *testing.T) {
	for _, test := range tests {
		var b bool
		switch test.x.(type) {
		case int:
			b = cmp.Less(test.x.(int), test.y.(int))
		case string:
			b = cmp.Less(test.x.(string), test.y.(string))
		case float64:
			b = cmp.Less(test.x.(float64), test.y.(float64))
		}
		if b != (test.compare < 0) {
			t.Errorf("Less(%v, %v) == %t, want %t", test.x, test.y, b, test.compare < 0)
		}
	}
}
func TestCompare(t *testing.T) {
	for _, test := range tests {
		var c int
		switch test.x.(type) {
		case int:
			c = cmp.Compare(test.x.(int), test.y.(int))
		case string:
			c = cmp.Compare(test.x.(string), test.y.(string))
		case float64:
			c = cmp.Compare(test.x.(float64), test.y.(float64))
		}
		if c != test.compare {
			t.Errorf("Compare(%v, %v) == %d, want %d", test.x, test.y, c, test.compare)
		}
	}
}
func TestSort(t *testing.T) {
	// Test that our comparison function is consistent with
	// sort.Float64s.
	input := []float64{1.0, 0.0, negzero, math.Inf(1), math.Inf(-1), math.NaN()}
	sort.Float64s(input)
	for i := 0; i < len(input)-1; i++ {
		if cmp.Less(input[i+1], input[i]) {
			t.Errorf("Less sort mismatch at %d in %v", i, input)
		}
		if cmp.Compare(input[i], input[i+1]) > 0 {
			t.Errorf("Compare sort mismatch at %d in %v", i, input)
		}
	}
}
