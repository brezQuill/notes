package slice

import "fmt"

type Ss struct {
	Len int
	Cap int
}

func Test_checkSlice() {
	var s1 []int
	s2 := make([]int, 0)
	s3 := make([]int, 1)
	s4 := make([]int, 1, 2)
	fmt.Printf("len(s1)=%d, len(s2)=%d, len(s3)=%d\n", len(s1), len(s2), len(s3))
	fmt.Printf("len(s4)=%d, cap(s4)=%d\n", len(s4), cap(s4))
}

func TestAppendNilSlice() {
	var s1 []Ss
	if s1 == nil {
		fmt.Println("s1 is not initialize")
	}
	s1 = append(s1, Ss{0, 0})
	if len(s1) == 1 {
		fmt.Println("s1 is not initialize, but append successful!")
	}
}

func Test_copy() {
	s1 := []byte("king of the world")
	a1 := []int{1, 2, 3}
	ca1 := a1
	cs1 := s1
	ca1[0] = 9
	cs1[0] = '!'
	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Printf("%s\n", s1)
}
