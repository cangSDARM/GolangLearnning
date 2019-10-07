package pak

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3" //不使用这个包, 但需要这个包
	"xorm"
)

var x = xorm.Engine

/*xorm
针对go的ORM库

支持链式操作
支持结合原生SQL语句操作
基于 LRU 规则的缓存器
使用乐观锁
*/
func orm() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./data.db") //使用sqlite数据库

	err = x.Sync(new(Account)) //自动同步表

	_, err = x.Insert(&Account{Name: "N"}) //增
	_, err = x.Delete(&Account{Id: 1})     //删, 传入的删除条件
	a := &Account{}
	has, err := x.Id(0).Get(a)
	fmt.Println(a)      //查
	a = &Account{Id: 0} //查
	if has || err == nil {
		a.Name = "N" //改
	}
	_, err = x.Update(a) //改了要update

	as = make([]*Account)
	err = x.Asc("name").Find(&as) //按照name正序排列获取

	//事务
	sess := x.NewSession() //创建事务对象
	defer sess.Close()

	sess.Begin()    //开始事务
	sess.Update()   //操作数据库(方法和x的一样，都可以Insert、Delete等)
	sess.RollBack() //回滚
	sess.Commit()   //提交

	//迭代查询
	x.Iterate(new(Account), func(idx int, bean interface{}) {
		//idx表示当前迭代的id, 和数据库没关系
		fmt.Printf("%d, %#v\n", idx, bean.(*Account))
	})

	//同上, 迭代查询
	b := new(Account)
	rows, err := x.Rows(new(Account))
	for rows.Next() {
		if err = rows.Scan(b); err == nil {
			fmt.Printf("%#v\n", b)
		}
	}
}

//表的字段要是全部大写的
type Account struct {
	Id      int64  //默认主键, 就这么写
	Name    string `xorm:"unique"`  //设置unique字段
	Version int    `xorm:"version"` //乐观锁, 每次修改递增1
}

//事件钩子
func (a *Account) BeforeInsert() {
	fmt.Printf("before insert hook", a.Name)
}
