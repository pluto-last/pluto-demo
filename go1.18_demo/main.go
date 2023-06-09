package main

import (
	"fmt"
)

func main() {

}

/**
1. [T] 是定义泛型类型的语法， T是可以随意替换的标志符
2. any 就是 interface{} 的别称
3. [T1, T2 any]， 代表定义了T1，T2这两个形参，他的实参是any类型的
4.
*/
type User[T1, T2 any] struct {
	ID   T1
	Name T2
}

// 泛型函数， 函数名字后边要加[T any]
// 返回泛型结构体， 返回值要加[T]
func NewUser[T1, T2 any](ID T1, name T2) *User[T1, T2] {
	return &User[T1, T2]{ID: ID, Name: name}
}

func Test1() {
	user := NewUser[int, int](1, 2)
	fmt.Println(user.Name)
	fmt.Println(user.ID)
}

// 定义一个泛型结构体，List可以是任意类型的
type Queue[T interface{}] struct {
	List []T
}

// 这个泛型结构体可以作为接收者，也就是相当任意类型都实现 Push Pop方法
func (s *Queue[T]) Push(elem T) {
	s.List = append(s.List, elem)
}

func (s *Queue[T]) Pop() (T, bool) {
	var elem T
	if len(s.List) == 0 {
		return elem, false
	}
	elem = s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return elem, true
}

func Test2() {
	s := Queue[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)

	s2 := Queue[float64]{}
	s2.Push(10.1)
	s2.Push(20.1)
	s2.Push(30.1)
	fmt.Println(s2)
	fmt.Println(s2.Pop())
	fmt.Println(s2)
}
