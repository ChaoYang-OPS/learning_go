package main

import (
	"fmt"
	"math"
	"time"
)

// 方法声明和普通函数的声明类似，只是在函数名字前面多了一个参数。这个参数把这个方法绑定到这个参数对应的类型上。

type Point struct{ X, Y float64 }

// 普通的函数
// 声明包级别的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法
// 附加的参数p称为方法的接收者，用来描述主调方法就像向对象发送消息
// 声明一个类型Point的方法，因此它的名字是Point.Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 指针接收者的方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// nil 是一个合法的接收者
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum 返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// 当定义一个类型允许nil作为接收者时，应当在文档注释中显式地标明
func Test(list int) func() int {
	return func() int {
		return 1 + list
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds()) // 86400
	// 调用方法的时候，接收者在方法名的前面
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // 5  函数调用
	// 表达式p.Distance称作选择子,因为它接收者p选择合法的方法Distance方法，选择子也用于选择结构类型中的某些字段值，就像p.X中的字段值
	fmt.Println(p.Distance(q)) // 5  方法调用
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // {2 4}
	tt := Test(5)
	fmt.Println(tt()) // 6
}
