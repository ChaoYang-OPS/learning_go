package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "BOb", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) // 返回指针
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

// 第一种定义方式在实例对应方法被调用时，实例的成员会进行复制
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
// }

// 通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	// e := &Employee{"0", "Bob", 20}
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
