package main

import "fmt"

// map 是无序的
func main() {
	m := map[string]string{
		"name": "forrest",
		"age":  "12",
		"sex":  "male",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println("m is", m)
	fmt.Println("m2 is", m2) // m2 == empty map
	fmt.Println("m3 is", m3) // m3 == nil

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	//delete(m, "age")
	if myAge, ok := m["age"]; ok {
		fmt.Println("myAge is ", myAge)
	} else {
		fmt.Println("Getting values error")
	}

	// map的key
	// map使用哈希表，必须可以比较相等的
	// 除了slice, map, function的内建类型都可以作为key
	// struct类型不包括上述字段，也可作为key
}
