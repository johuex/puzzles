package durtymemory

import (
	"strings"
	"unsafe"
)

func GetMem(size int) []int {
	r := make([]int, size)
	return r
}

// field order depends
type slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func GetDurtyMem(size int) []int {
	var s strings.Builder
	s.Grow(size)
	p := unsafe.StringData(s.String())
	return Transform(unsafe.Slice(p, size))

}

func Transform(data []byte) []int {
	sliceData := unsafe.Pointer(&data[0]) // pointer on first elem of array
	sizeType := int(unsafe.Sizeof(int(0)))
	length := len(data) / sizeType

	var result []int
	resultPtr := (*slice)(unsafe.Pointer(&result)) // force converting
	resultPtr.data = sliceData
	resultPtr.len = length
	resultPtr.cap = length

	return result
}
