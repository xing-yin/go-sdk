// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"syscall"
)

// KernelVersion returns major and minor kernel version numbers, parsed from
// the syscall.Uname's Release field, or 0, 0 if the version can't be obtained
// or parsed.
//
// Currently only implemented for Linux.
func KernelVersion() (major int, minor int) {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return
	}

	rl := uname.Release
	var values [2]int
	vi := 0
	value := 0
	for _, c := range rl {
		if '0' <= c && c <= '9' {
			value = (value * 10) + int(c-'0')
		} else {
			// Note that we're assuming N.N.N here.
			// If we see anything else, we are likely to mis-parse it.
			values[vi] = value
			vi++
			if vi >= len(values) {
				break
			}
			value = 0
		}
	}
	switch vi {
	case 0:
		return 0, 0
	case 1:
		return values[0], 0
	case 2:
		return values[0], values[1]
	}
	return
}
