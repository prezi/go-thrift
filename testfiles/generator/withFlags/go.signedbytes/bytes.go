// This file is automatically generated. Do not modify.

package gentest

import (
	"fmt"
)

var _ = fmt.Sprintf

type ByteAndListByte struct {
	AByte     int8   `thrift:"1,required" json:"a_byte"`
	AListByte []int8 `thrift:"2,required" json:"a_list_byte"`
}

func (b *ByteAndListByte) GetAByte() (val int8) {
	if b != nil {
		return b.AByte
	}

	return
}

func (b *ByteAndListByte) GetAListByte() (val []int8) {
	if b != nil {
		return b.AListByte
	}

	return
}
