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
9、Go语言仅支持循环关键字for
10、if条件
(1) condition表达式结果必须为布尔值
(2) 支持变量赋值
  if var declaration; condition{

  }
11、 switch条件
(1) 条件表达式不限制为常量或者整数
(2) 单个case中，可以出现多个结果选项，使用逗号分隔
(3) 与C语言等规则相反，GO语言不需要用break来明确退出一个case
(4) 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个
位运算符
&^ 按位置零
1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0
12 Map元素的访问
1、在访问Key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在

if .... else....的逻辑作用等同
一、类型的预定义值
1、 math.MaxInt64
2、 math.MaxFloat64
3、 math.MaxUint32

二、用 == 比较数组
1、相同维数且含有相同个数元素的数组才可以比较
2、每个元素都相同的才相等

# go run hello_world.go Yang
hello world Yang
```