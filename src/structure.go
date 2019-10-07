//当前程序的包名. 必须写在第一行
package structure

//导入其它包.
//若导入后没用, 编译会报错
import (
	"fmt"
	IO "io"  //等价于: import io as IO
	. "time" //可以使得省略包名调用函数(不建议)
)

/*命名约定:	(强制)
以package为最小的block

1. 若首字母小写为: private
2. 若首字母大写为: public
*/

/*组. 适用于全局变量的声明, 函数内不能使用

const (
	pi = 3.14
	alpha = "a"
)
等价于:
const pi = 3.14
const alpha = "a"
等价于:
const pi, alpha = 3.14, "a"
*/

//全局变量
var name = "go"

//类型别名.
//	底层相同(int方法没保留)
//	算是自定义类型. 因为仍然需要显示转换
//	可以给gopher绑定任何函数
type gopher int

//只有package为main的才能定义main函数
//且package为main的必须定义main函数
func mainS() {
	/*变量的声明*/
	var ii int //声明变量
	var i = 0  //同上. 自动推断类型
	ie := 0    //同上. 极简形式

	/*多个变量同时声明或赋值*/
	var a, b, c int   //多个变量同时声明
	a, b, c = 1, 3, 5 //同时赋值

	/*忽略某值*/
	a, _, c = 1, 4, 5 // _ 为忽略不需要位置的东西(此例的4没了)

	/*类型转换*/
	i = byte(ie) //go不能隐式转换, 必须显示的强制转换
	//string 和 数值转换:
	//
	//  普通互转时, 转换的是ASCII码而不是对应数字
	//  想直接字面量转换时, 使用"strconv"包:
	//    strconv.Itoa(int)
	//    转回去: strconv.Atoi(str)

	/*常量*/
	const cc int = 1  //常量声明和赋值和变量相同(把var换成const就行)
	const cl = len(c) //常量的函数必须是内置函数
	//在常量组中:
	//
	//  省略赋值:
	//      没赋值的继承上一行的内容
	//  const(
	//    a, b = 1, 2
	//	  c, d  //则 c==a, d==b
	//  )
	//
	//  常量计数器:(实现枚举)
	//      iota遇见一个常量组即清零; 组中每定义一个常量iota+=1
	//  const(
	//	  a = "A"
	//	  b = iota	// b=1
	//	  c = 5
	//	  d = iota  // d=4
	//  )

	/*运算符*/
	/*根据优先级 high->low:

	^(按位取反)  !(非)
	*   /   %  <<  >>  &(与)  &^(x&^y: 如果y的bit位上为0, 则取x上对应位置的值; 如果y的bit位上为1, 则取0)
	+   -   |(或)  ^(异或: 不进位加法)
	== !=   <  <=  >=  >
	<-							(用于Chanel)
	&&(逻辑与)
	||(逻辑或)
	*/

	i++ //允许. 但是没有 ++i; 也不允许 i = i++, 只能单独写

	/*控制语句
	所有控制语句都可以初始化一个变量给自己用
	*/
	if a := 1; a == 1 {
		fmt.Printf('a')
	} else {
		fmt.Printf('a')
	}

	for i := 0; i < 1; i++ {
		fmt.Println("普通for循环")
	}

	for {
		i++
		if i < 1 {
			fmt.Println("go中的死循环while")
		}
	}

	for i < 2 {
		fmt.Println("while")
	}

	for index, value := range pab {
		//value 是对应的拷贝, 而不是引用
		fmt.Println("golang的for-range循环. 用于slice和map")
	}

	switch a := 1; a {
	case a > 0:
		//go自己默认break
		fallthrough //穿透下一case
	default:
	}

	/*标签, goto, continue, break*/
LABEL:
	for {
		for i := 1; i < 10; i++ {
			if i > 5 {
				break LABEL    //break 之后不会进刚刚打断的循环
				goto LABEL     //goto 之后会继续下一条指令, 不管是不是刚刚打断的循环
				continue LABEL //continue 也能用标签, 下一条若是刚刚打断的循环, 则同样是继续下次循环但不会死循环
			}
		}
	}

	/*数组
	值类型,
	} 必须在同一行, 换行会报错
	*/
	var ia [2]int
	iaa := [5]int{0: 1, 3: 5} //指定iaa[0] = 1; iaa[3] = 5
	iab := [...]int{2, 3, 5}  //len(iab) = 3
	iac := [...]int{19: 20}   //len(iac) = 20

	ia[1] = 2

	var iad = [2][3]int{ /*多维数组*/
		{1, 2, 3},
		{2, 3, 4}}

	/*切片
	引用类型, 可变数组格式
	但是append或其它操作改变容量时, 会重新分配地址, 造成改变失效
	*/
	var is []int                //默认使用内部的数组来创建
	var isa = iaa[2:5]          //isa截取iaa的2-4
	var isb = iaa[2:]           //isa截取iaa的2-最后
	var isd = iaa[:2]           //isd截取iaa的开始-2
	var isc = make([]int, 3, 5) //初始化3个int的切片, 且容量n次分配以2n*5算

	fmt.Println(is[2]) //索引值
	cap(is)            //现有容量, 索引不可超过这个值

	is = append(is, 2, 3, 5) //往is后面加:2,3,5

	copy(isa, isb) //将isb的东西拷贝到isa对应位置中. 具体位置限制(拷贝多少)以较短的那个为准

	/*Map
	key必须支持 == 和 != 运算(即没有函数/Map/切片)
	*/
	var im map[int]string //c#: var im = new Dictionary<int, string>();
	im = map[int]string{} //等价于: im = make(map[int]string)
	im[1] = "MAP"

	_, i = im[4] // _ 是对应的值, i是bool表示该键值对是否存在

	delete(im, 1) //删除对应key, value

	/*defer
	在定义的函数return后, 按照定义顺序的逆序执行
	严重错误也执行(类似于其它语言的finally)
	*/
	defer fmt.Println("2") //后打印
	defer fmt.Println("1") //先打印

	for i := 0; i < 3; i++ {
		defer fmt.Println(i) //打印: 2, 1, 0    定义时拷贝每个i -> 编译后, 三个defer逆序调用
	}
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) //打印: 3, 3, 3.		定义时是闭包, 获取i的地址 -> 编译后, 三个defer逆序调用
		}()
	}
	for i := 0; i < 3; i++ {
		func() {
			fmt.Println(i) //打印: 3, 3, 3.    定义时是闭包, 获取i的地址 -> 编译后, 三个闭包顺序调用
		}()
	}

	/*错误处理
	recover() 必须在 panic() 之前
	*/
	defer func() {
		err := recover() //recover()捕获Panic. 必须定义在defer里(因为defer在严重错误时也执行)
		if err != nil {
			fmt.Println("Recover")
		}
	}()
	panic("Panic") //类似于py的 raise

	/*结构体
	 */
	type str struct { //声明
		Name string //每个字段声明后面跟着分号, 但是可以忽略
		Age  int
	}
	var ist = str{
		Name: "A",
		Age:  2, //初始化的话必须加 ,
	}
	ist.Age = 1 //赋值

	/*匿名结构体*/
	var istt = struct {
		Name, City string
	}{
		Name: "A",
		City: "1",
	}

	/*嵌套结构体*/
	type stri struct {
		Name, City string
		CON        struct {
			Age int
		}
	}
	var isttt = stri{}
	isttt.CON.Age = 1 //嵌套的结构体只能用这种来赋值, 不能在初始化时赋值

	/*结构体匿名字段*/
	type strin struct {
		string
		int
	}
	var istttt = strin{"A", 1} //赋值顺序必须是定义时顺序, 且只能这么赋值

	/*结构体Tag*/
	type stringg struct {
		Name string `Tag:"AA" TAG:"BB"` //使用 `` 指定tag, 键值对形式, 且使用space区分多个tag
	}
	user := &stringg{"Name"}
	s := reflect.TypeOf(user)              //通过反射获取type定义
	v, err := s.Field(0).Tag.Lookup("Tag") //获取指定Tag

	/*接口和结构体的一部分看fakeOOP.go*/

	/*反射部分看main.go*/

	/*并发部分看concurr.go*/

	/*指针*/
	var p *int = &a
	var pa *[4]int     //数组指针
	pab := new([4]int) //同上

	pac := &str{}  //结构体指针
	pac.Name = "A" //操作结构体指针的数据不需要前面加 *

	var pstt = &struct { //匿名结构体指针
		Name string
		Age  int
	}{
		Name: "A",
		Age:  1,
	}

	fmt.Println("hello")
}

/*函数
不支持: 嵌套, 重载, 默认参数
支持: 无需声明类型, 不定长形参, 多返回值, 命名返回值参数, 匿名函数, 闭包

对参数有值拷贝:{
	int, string, struct		//变成地址拷贝: 传指针
]和地址拷贝:{
	slice, map
}
func 函数名(形参 类型)(返回值类型) {

}
*/
//同类型简写类型: a, b, c都是int
func f(a, b, c int) int {}

//命名返回值参数
func u() (a, b, c int) {
	a, b, c = 1, 2, 3 //直接赋值, 不用声明
	return            //直接返回a, b, c(也可以显示的加上)
	//return 1	没什么卵用. 仍然返回a, b, c
	//return t  没什么卵用. 仍然返回a, b, c
}

//不定长参数. (不定长int类型参数), 必须是参数列表的最后一个
func n(a ...int) {
	fmt.Println("a是个切片")
	//切片是值拷贝. 因此不定长参数改变后, 不影响原参数
	//直接传的切片会被改变
}

//匿名函数
func warp() { //并不是指warp函数
	var c = func() {} //a是匿名函数
	c()
}

//闭包
func t(x int) func(int) int {
	fmt.Printf("%p", &x)
	return func(y int) int {
		fmt.Printf("%p", &x) //两次打印的x地址相同 -> 因此是闭包
		return x + y
	}
	//调用:
	//t(10)(100) == 110
}

//无需声明类型
func i(i) {}
