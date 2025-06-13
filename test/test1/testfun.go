package test1

import (
	"fmt"
	"os"
)

func funArrayArg(arr [5]int) {
	arr[0] = 11
	arr[1] = 12
	fmt.Println("intestArrayArg=", arr)
}

func TestArrayArg() {
	var envVar string = os.Getenv("ENV_VAR")
	fmt.Println("envVar=", envVar)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	funArrayArg(arr)
	fmt.Println("main=", arr)

	fmt.Printf("变量 envVar 的地址: %p\n", &envVar) // 输出地址（以指针形式）
	fmt.Printf("变量 envVar 的地址: %v\n", &envVar) // 也可以用 %v 输出地址
}
