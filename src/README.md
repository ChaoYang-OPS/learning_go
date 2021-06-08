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
# go run hello_world.go Yang
hello world Yang
```