package printf

import "fmt"

func print(params ...interface{}) {
	fmt.Printf("%s, %d\n", params...)
}

func Test_params() {
	params := []interface{}{"First", 2}
	print(params...) // 解包切片
}
