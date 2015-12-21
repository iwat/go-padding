// go-padding - Simple padding package for Golang

// Copyright (c) 2015 Chaiwat Shuetrakoonpaiboon. All rights reserved.
//
// Use of this source code is governed by a MIT license that can be found in
// the LICENSE file.

// Package padding provides simple functions for block padding.
package padding

type BlockPadding interface {
	Pad(data []byte, blockSize int) []byte
	Unpad(data []byte) []byte
}
