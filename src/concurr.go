package concurr

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*并发
Go的goroutine只是官方实现的"线程池"
Go可以设置跑的CPU核数, 以发挥多核能力实现并行
*/
func Con() {
	n := runtime.GOMAXPROCS(runtime.NumCPU()) //设置CPU核数, 返回之前的CPU核数

	runtime.Gosched() //该线程主动让出时间片, 让其它线程执行(如果有的话)

	runtime.Goexit() //终止该线程

	/*Channel
	Go使用Channel实现消息共享, 而不是内存共享. Channel大多数都是阻塞同步
	通过make创建, close关闭
	Channel是引用类型
	可以设置单向(只读, 只写)或双向
	可以设置缓存大小, 在未被填满时不会阻塞
	*/
	c := make(chan bool, 5) //Channel, 可读可写, 缓存5

	/* 单向
	双向可以强制转为单向的
	*/
	cr := make(chan<- float64) //只写
	var cw <-chan int          //只读

	fmt.Println(len(c), cap(c)) //缓冲区剩余数据; 缓冲区长度

	go func() { //运行了一个goroutine(线程)
		fmt.Println("动作")
		c <- true //输入信号
		close(c)
	}()

	<-c //等待这个Channel的信号, 否则被阻塞

	wg := sync.WaitGroup(10)
	for i := 0; i < 10; i++ {
		wg.Done()
	}
	wg.Wait()

	/*Select
	可以设置一个或多个Channel的发送与接收
	所有case都会被求值. 求值顺序: 自上而下、从左到右.
	1. 如果有多个case内语句可以运行, 随机选一个执行, 其他不会执行
	2. 如果没有case内的语句可运行, 且有default语句, 那么执行default的动作
	3. 如果没有case内的语句可运行, 且没有default语句, 那么select将阻塞, 直到某个case通信可以运行
	可设置超时
	*/

	select {
	case c <- x:
	case value, ok := <-c:
	case <-time.After(3 * time.Second): //超时设置
	}
}

//相关blog: https://mr-dai.github.io/go_channels_on_steroid/
