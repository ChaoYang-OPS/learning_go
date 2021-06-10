```shell
# 运行 go run hello_world.go
hello world
# 构建 go build hello_world.go
./hello_world
hello world
知识点
应用程序入口
必须是main包: package main
必须是main方法： fun main()
文件名不一定是main.go
退出返回值
与其它主要编程语言的差异
1、Go中main函数不支持任何返回值
2、通过os.Exit来返回状态
3、main函数不支持传入参数
4、在程序中直接通过os.Args获取命令行参数
5、GO语言不允许隐式类型转换
6、别名和原有类型也不能进行隐式类型转换
7、不支持指针运算
8、string是值类型，其默认的初始化值为空字符串，而不是nil
一、类型的预定义值
1、 math.MaxInt64
2、 math.MaxFloat64
3、 math.MaxUint32

# go run hello_world.go Yang
hello world Yang
```