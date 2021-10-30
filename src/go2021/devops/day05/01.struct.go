package main

import "fmt"

type User struct {
	id    int64
	add   string
	name  string
	tel   string
	order Order
}

type Order struct {
	size int
	age  int64
	info string
}

func main() {
	var u User
	u2 := User{
		id:   1,
		add:  "test",
		name: "hello",
		tel:  "123",
	}
	fmt.Printf("%T\n", u2)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u2)
	u2.name = "Lsunstack"
	fmt.Printf("%#v\n", u2)

	var pu *User
	fmt.Printf("%T\n", pu)  // *main.User
	fmt.Printf("%#v\n", pu) // (*main.User)(nil)
	// 匿名结构体 --->直接初始化给一个变量
	var test struct {
		id   int
		name string
		age  int
	}
	fmt.Printf("%#v\n", test) // struct { id int; name string; age int }{id:0, name:"", age:0}
	fmt.Printf("%T\n", test)  // struct { id int; name string; age int }

	user := struct {
		id   int
		name string
		age  int32
	}{1, "Yang", 30}
	// struct { id int; name string; age int32 } --- struct { id int; name string; age int32 }{id:1, name:"Yang", age:30}
	fmt.Printf("%T --- %#v\n", user, user)
}
