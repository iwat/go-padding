// go-padding - Simple padding package for Golang

// Copyright (c) 2015 Chaiwat Shuetrakoonpaiboon. All rights reserved.
//
// Use of this source code is governed by a MIT license that can be found in
// the LICENSE file.

// Package padding provides simple functions for block padding.
package padding

import (
	"bytes"
)

var Pkcs7 pkcs7

type pkcs7 struct{}

func (p pkcs7) Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func (p pkcs7) Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
