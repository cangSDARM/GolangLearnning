# Golang

## 有关环境
```go
go env  //查看环境变量
go get  //获取远程包
go fix  //更新包
go bulid //编译包
go run  //编译且运行程序
go install  //编译且安装包
go test //运行测试文件
go doc  //查看文档
```
需要先配置`GOPATH`环境变量, 来指定工作目录

## 目录结构
    - bin   //可执行文件
        | 项目名.exe   //只有package名为main的才生成可执行文件, 否则生成包文件
    - pkg   //包文件
        |- 当前OS名_当前架构名
            |- 项目名
                 | 包名.a     //编译后的包文件
    - src   //源码

## 组织架构
1. Go通过`package`来组织的
2. 只有`package`为`main`的包可以包含`main`函数
3. 一个程序有且只有一个`main`包
4. 通过`import( package名, )`来导入其它包
5. 没用的包, 编译会报错

## 基本类型
1. bool: true, false
2. int/unit (根据平台, 选择int32/uint32 | int64/uint64)
3. int8/uint8, int16/uint16, int32/uint32, int64/uint64
4. float32, float64
5. byte (uint8别名)
6. complex64/complex128 (复数)
7. array, struct, string
8. slice, map, chan
9. interface,
10. func

## 注意事项
1. 如果要改数据的值, 用指针传递; 不需要改则用值传递. 最好不要直接修改参数, 而返回修改的值

## 装环境的正确步骤
1. vi /etc/profile
2. export PATH=$PATH:/usr/local/go
3. export GOPROXY=http://goproxy.io
4. 随便找个文件夹，go mod init 项目名
5. cd 到项目文件夹下
6. go get 对应git
7. cd 对应git下
8. 参考目录下的.travils.yml，运行script下的代码跑测试
9. 写代码
