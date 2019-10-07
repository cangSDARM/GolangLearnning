package OOP

/*对象和继承*/
type Human struct {
	Name string
	Age  int
}

type Person1 struct {
	Human //匿名字段实现继承
	City  string
}

type Person2 struct {
	Human
	Tone string
	id   *string //指针字段初始化时加个取地址符&就行
}

func tmain() {
	var a Person1 = Person1{Human{"A", 1}, "S"}
	var b = Person2{Human: Human{Name: "A"}, Tone: "S", &"xxx"}
	a.Name = "E"            //这可以
	a.Human.Name = "Q"      //这也可以
	a.Human = Human{"A", 1} //这也可以
}

/*对象方法*/
func (a Person1) Print() {
	//强制指定调用者
	//这样就只能用Person来调用Print
	//但是golang没有方法重载

	a.Name = "A" //修改对象的值依然遵循golang的默认规则(值传递或引用传递). 解决办法: a *Person1

	//调用:
	//	1. var a := Person1()
	//	   a.Print()
	//	2. (Person1).Print(a)
}
func (a *Person1) Printf() {
	a.Print()     //指针变量可以调用普通方法
	(*a).Printf() //普通变量也可以调用指针方法
}

/*接口
不需要显示调用接口
只要对象实现了相关接口的方法, 则该对象默认实现了相关接口
因此所有对象都实现了空接口

USB(obj) 将对象转为接口时, 接口获取的是对象的拷贝.
*/
type USB interface {
	Name() string //声明Name函数, 返回类型为string
	Connect()

	//实现:
	//	type Phone struct{}
	//	func (p Phone)Name() string{ return "string" }
	//	func (p Phone)Connect(){ }
	//
	//	var a USB
	//	a = Phone{}
	//	a.Name()	//print: string
	//
	//实现接口的类型判断:
	//	var usb = USB
	//	pc, ok := usb.(Phone)	//判断usb是否是Phone类型. 返回Phone对象和Bool
}

/*嵌入接口*/
type Connector interface {
	Connect()
}
type USB3 interface {
	Name() string
	Connector //这样USB3就有Connect方法了
}

/*高接口变低接口*/
func Switch() {
	var a USB3
	Connector(a) //这是可以的

	var b Connector
	USB3(b) //这是不行的
}

/*type-switch
大量类型判断
*/
func Disc(obj interface{}) { //obj是个空类型
	switch v := obj.(type) { //系统判断obj类型
	case Phone: //若v是Phone对象
	default:
	}
}
