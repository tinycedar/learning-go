package io

import (
	"fmt"
	"io"
	"testing"
)

type writerTest struct {
	bw io.ByteWriter
}

func (wt *writerTest) WriteByte(c byte) error {
	return nil
}

func TestRead(t *testing.T) {
	rb := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(rb)

	byteWriter := new(writerTest)
	byteWriter.WriteByte('a')
}
