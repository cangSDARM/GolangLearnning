package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	fmt.Println("Hello world")
	fmt.Printf("%d \n", 1)

	sort.Ints() //对int类型排序

	time.Sleep(2 * time.Second) //睡眠2s
}

func timeout() {
	timer := time.NewTimer(3 * time.Second)

	timer.Reset(1 * time.MicroScond)

	timer.Stop()
	<-timer.C
}

func interval() {
	ticker := time.NewTicker(3 * time.Second)
	i := 0
	for {
		<-ticker
		i++
		if i > 5 {
			ticker.Stop()
			break
		}
	}
}

func reflection(o interface{}) {
	/*Reflect*/
	t := reflect.TypeOf(o)
	fmt.Printf("Type: ", t.Name())
	v := reflect.ValueOf(o)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("该对象不能被反射. 若o是个对象的地址, 请传值拷贝")
	}

	fmt.Println("反射的字段: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := t.Field(i).Interface()
		fmt.Printf("%s: %v = %v", f.Name, f.Type, val)
	}

	fmt.Println("反射的方法: ")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v", m.Name, m.Type)
	}

	fmt.Println("反射的匿名字段: ")
	fmt.Printf("%#v", t.Field(0))                  //若t的序号为0的是其匿名字段
	fmt.Printf("%#v", t.FieldByIndex([]int{0, 0})) //取出序号为0的匿名字段的0序号的内容

	fmt.Println("反射对内容的修改: ")
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不能被反射")
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsVaild() {
		fmt.Println("该字段不存在")
	}
	if f.Kind() == reflect.String {
		f.SetString("修改的内容")
	}

	fmt.Println("反射对方法的调用: ")
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("第一个参数"), reflect.ValueOf("第二个参数")}
	res = mv.Call(args)
}
