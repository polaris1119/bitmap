// Copyright 2014 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	studygolang@gmail.com

package bitmap_test

import (
	"testing"
	"bitmap"
)

func TestBitmapSort(t *testing.T) {
	bmap := bitmap.NewBitmapSize(7)

	original := [5]uint64{4,6,3,1,7}
	expected := [5]uint64{1,3,4,6,7}
	actual := [5]uint64{}

	for _, offset := range original {
		bmap.SetBit(offset, 1)
	}

	var i uint64
	var offset, maxpos uint64 = 0, bmap.Maxpos() + 1
	for ; offset < maxpos; offset++ {
		if bmap.GetBit(offset) == 1 {
			actual[i] = offset
			i++
		}
	}
	
	if expected != actual {
		t.Errorf("expected:%#v, actual:%#v", expected, actual)
	}
}
