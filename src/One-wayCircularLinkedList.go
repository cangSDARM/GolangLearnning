package main

import (
	"fmt"
)

type LinkedList struct {
	No   int //data
	Next *LinkedList
}

func Insert(num int) *LinkedList {
	first := &LinkedList{}
	curList := &LinkedList{}

	if num < 1 {
		fmt.Println("链表个数至少为1")
		return first
	}
	for i := 1; i < num; i++ {
		list := &LinkedList{
			No: i,
		}
		if i == 1 {
			first = list
			curList = list
			curList.Next = first
		} else {
			curList.Next = list
			curList = list
			curList.Next = first //构成循环链表
		}
	}
	return first
}

func Print(first *LinkedList) {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}

	curList := first
	for {
		fmt.Printf("当前NO: %d", curList.No)
		if curList.Next == first {
			break
		}
		curList = curList.Next
	}
}

//Josephu问题
//设编号为 1-n 的n个人围坐一圈. 约定以编号k的人从1开始报数, 数到m的那人出列; 第m+1个又从1开始报数, 数到m的那人又出列...以此类推, 直到所有人出列为止, 打印其出列的编号序列
func Josephu(first *LinkedList, k int, m int) {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}

	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	for i := 0; i < k-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	//开始出列
	for {
		//m步
		for i := 0; i < m-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("出列: %d\n", first.No) //出列的仍然在内存中,没有被clear掉
		//删除
		first = first.Next
		tail.Next = first
		//唯一一个
		if tail == first {
			fmt.Printf("最后, 出列: %d", first.No)
			break
		}
	}
}

func main() {
	head := Insert(5)
	Print(head)
}
