package test1

import "fmt"

func Test_goClosure() {
	// 准备一个字符串
	str := "hello world"
	fmt.Printf("out-func()-1,str=%s..addr=%p\n", str, &str)

	// 创建一个匿名函数
	foo := func() {

		// 匿名函数中访问str
		str = "hello dude"
		fmt.Printf("in func(),str=%s..addr=%p\n", str, &str)
	}

	// 调用匿名函数
	foo()
	fmt.Printf("out-func()-2,str=%s..addr=%p\n", str, &str)
}
