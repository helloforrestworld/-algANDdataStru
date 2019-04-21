package main

import (
	"fmt"
)

/* panic
	1.停止当前函数执行
	2.一直向上返回，执行每一层的defer
	3.如果没有遇见recover, 程序退出
*/

/* recover
	1.仅在defer函数调用
	2.获取panic的值
	3.如果无法处理,可重新panic
*/

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("Unknow error: %v", r))
		}
	}()
	//panic(errors.New("robust"))
	panic("123")
}

func main() {
	tryRecover()
}
