package main

import "fmt"

// go面向对象
// 仅支持封装， 不支持继承和多态
// 没有class 只有struct
// 面向接口编程

// 关于值接收者和指针接收者
// 1.要改边内容必须使用指针接收者
// 2.结构过大使用指针接收者
// 3.一致性：如果指针接收者，最好都是指针接收者

// 关于封装
// 名字一般使用CamelCase
// 首字母大写: public
// 首字母小写: private

// 包
// 同一个目录下只能有一个包 package
// 为结构定义的方法必须放在同一个包内
// 可以为不同文件

// 如何扩充系统类型或者别人的类型
// 1.定义别名
// 2.使用组合

type treeNode struct {
	value       int
	left, right *treeNode
}

func createTreeNodes(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.left.left = createTreeNodes(7)
	root.left.right = createTreeNodes(10)
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.right.right = createTreeNodes(4)

	//root.left.right.print()
	//root.left.right.setValue(15)
	//root.left.right.print()

	fmt.Println()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	var pRoot *treeNode // nil 指针
	pRoot.setValue(200) // nil setValue不会报错

	root.traverse()
}
