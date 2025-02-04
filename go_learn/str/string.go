package str

import (
	"fmt"
	"unsafe"
)

func b2s(b []byte) string {
	// return *(*string)(unsafe.Pointer(&b))  // 第一种零拷贝强转方法
	return unsafe.String(unsafe.SliceData(b), len(b)) // 第二种方法
}

func s2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func Test_forceChange() {
	b := []byte("king of the world")
	s := b2s(b)
	fmt.Println(s)
	b1 := s2b(s)
	b1[0] = 'i'
	fmt.Printf("s : %#v, b1 : %s\n", s, b1)
}
