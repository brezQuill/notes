package itface

import (
	"bufio"
	"bytes"
	"fmt"
)

// 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
type WordsCount int
type LinesCount int

func (s *WordsCount) Write(p []byte) (n int, err error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		// 由于是值传递, 这里想改变原调用者的值, 必须使用指针
		*s++
	}
	return int(*s), nil
}

func (s *LinesCount) Write(p []byte) (n int, err error) {
	var sc = bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*s++
	}
	return int(*s), nil
}

func Test_WordCount_LineCount() {
	var wc WordsCount
	wc.Write([]byte("hello world"))
	var lc LinesCount
	lc.Write([]byte(`hello
1
2
3
world`))
	fmt.Println(wc, lc)
	fmt.Fprintf(&wc, "Hello, %s", "dugulp")
	fmt.Println(wc, lc)
}
