package main

import (
	"awesomeProject/charter4/mock"
	"awesomeProject/charter4/real"
	"fmt"
)

// duck typing https://studygolang.com/articles/214
// 在面向对象的编程语言中，当某个地方（比如某个函数的参数）需要符合某个条件的变量（比如要求这个变量实现了某种方法）时，什么是判断这个变量是否“符合条件”的标准？
// 如果某种语言在这种情况下的标准是： 这个变量的类型是否实现了这个要求的方法(并不要求显式地声明），那么这种语言的类型系统就可以称为 duck typing

// interface {} 表示任何类型

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

const url  = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string {
		"value": "the content poster add",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"value": "the content poster add",
	})
	return s.Get(url)
}


func main()  {
	var r1 Retriever = &mock.Retriever{"this is a fake imooc.com"}
	//fmt.Println(download(r1))
	fmt.Println("r1 is ", r1)
	inspect(r1)

	var r2 Retriever = &real.Retriever{"mozlla/5.0", 500}
	//fmt.Println(download(r2))
	fmt.Println("r2 is ", r2)

	// Type assertion
	if realRetriever, ok := r2.(*real.Retriever); ok {
		fmt.Println("Real Retriever TimeOut: ",realRetriever.TimeOut)
	}
	inspect(r2)

	// interface 组合
	fmt.Println("session run")
	var r3 RetrieverPoster = &mock.Retriever{"this is a fake imooc.com"}
	fmt.Println(session(r3))
}

func inspect(r Retriever)  {
	fmt.Println("inspecting", r)
	fmt.Printf(" > type=%T value=%v\n ", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("> can read Contents, ", v.Contents)
	case *real.Retriever:
		fmt.Println("> can read UserAgent ", v.UserAgent)
	}
}
