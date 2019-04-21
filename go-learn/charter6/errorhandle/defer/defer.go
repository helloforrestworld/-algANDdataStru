package main

import (
	"awesomeProject/charter5/functional/fib"
	"bufio"
	"fmt"
	"os"
)

// defer 调用
// 确保调用在函数结束时发生
// 参数在defer语句时计算
// defer列表为先进后出

//	何时使用defer调用
//	Open/Close
//	Lock/Unlock
//	PrintHeader/PrintFooter

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
}

func tryDeferInFor()  {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			break
		}
	}
}

func writeFile(filename string)  {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("this is a custom error!")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//tryDeferInFor()
	writeFile("fib.txt")
}
