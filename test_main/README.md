
## 4.4变量
Go和许多编程语言不同，它在声明变量时将变量的类型放在变量的名称之后  
### 示例一
对比c语言指针声明：a是指针b不是
```c
int* a, b;
```
go指针声明：a和b都是指针
```go
var a, b *int
```
### 示例二
从左到右顺序阅读，代码更容易理解
```go
var a int
var b bool
var str string
```
等同形式
```go
var (
  a int
  b bool
  str string
)
```
当一个变量被声明之后，系统自动赋予它该类型的零值：int为0,float为0.0,bool为false,string为空字符串,指针为nil<br>
Go编译器可以根据变量的值来自动推断其类型
```go
var a=15
var b=false
var str = "hello world"
```
### 简短形式，使用:=赋值操作符
```go
a := 50
b := false
```
a和b的类型（int和bool）由编译器自动推断
不能重复声明；局部变量声明但是没有在相同代码块中使用，编译会报错
多变量同时赋值
```go
a, b, c = 5, 7, 'abc'
```
声明
```go
a, b, c := 5, 7, 'abc'
```
如果想交换两个变量值：a, b = b, a
### init函数
变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

每一个源文件都可以包含一个或多个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。

一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。

## 4.5基本类型和运算符
### 布尔类型boll
当运算符左边表达式的值已经能够决定整个表达式的值的时候（&& 左边的值为 false，|| 左边的值为 true），运算符右边的表达式将不会被执行
### 数字类型
Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码（详情参见 二的补码 页面）。

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

这些类型的长度都是根据运行程序所在的操作系统类型所决定的：

- int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
- uintptr 的长度被设定为足够存放一个指针即可。

:=这个符号直接取代了var和type,这种形式叫做简短声明。不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量

Go 语言中没有 float 类型
## 4.6字符串
go对单双引号使用

通过函数len()来获取字符串所占的字节长度,例如：len(str)

字符串str的第i个字节：str[i - 1]

获取字符串中某个字节的地址的行为是非法的，如：&str[i]
## 4.7strings和strconv包
```go
strings.HasPrefix(s, prefix string) bool //s以prefix开头
strings.HasSuffix(s, suffix string) bool //s以prefix结尾
strings.Contains(s, substr string) bool //s包含prefix
strings.Index(s, str string) int //s中检索第一个prefix
strings.Replace(str, old, new, n) string //str的前n替换为new
...
```
## 4.9指针
```go
var i1 = 6
var pi = &i1 //取变量地址
var p1 *int  //声明指针 nil
```
一个指针变量可以指向任何一个值的内存地址；指针类型前面加上 * 号（前缀）来获取指针所指向的内容；当一个指针被定义后没有分配到任何变量时，它的值为 nil

Go 语言和 C、C++ 以及 D 语言这些低级（系统）语言一样，都有指针的概念。但是对于经常导致 C 语言内存泄漏继而程序崩溃的指针运算（所谓的指针算法，如：pointer+2，移动指针指向字符串的字节数或数组的某个位置）是不被允许的。Go 语言中的指针保证了内存安全，更像是 Java、C# 和 VB.NET 中的引用。

因此 c = *p++ 在 Go 语言的代码中是不合法的
## 5.1 if-else 结构
```go
if condition1 {
	// do something	
} else if condition2 {
	// do something else	
}else {
	// catch-all or default
}
```
关键字 if 和 else 之后的左大括号 { 必须和关键字在同一行
```go
if x{
}
else {	// 无效的，报错
}
```
## 5.4 for结构
```go
for i := 0; i < 5; i++ {
	fmt.Printf("This is the %d iteration\n", i)
}
```
循环的头部，它们之间使用分号 ; 相隔，但并不需要括号 () 将它们括起来。例如：for (i = 0; i < 10; i++) { }，这是无效的代码;

Break退出循环；

continue忽略剩余的循环体而直接进入下一次循环过程
## 6.0函数

Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）

Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。

这样是不正确的 Go 代码：
```go
func g()
{
}
```
它必须是这样的：
```go
func g() {
}
```
目前Go没有泛型概念，就是说不能声明那种支持多种类型的函数；函数重载也不被允许。
### 函数参数与返回值
函数能返回零个或者多个值，通过return返回一组值；

Go 函数参数传递参数的副本。函数接收参数副本之后，对副本的更改，不会影响到原来的变量；可以传递指针修改作为参数的变量。
```go
var num int = 10
var numx2, numx3 int

func main() {
    numx2, numx3 = getX2AndX3(num)
    PrintValues()
    getX2(&num)
    PrintValues()
}
func PrintValues() {
    fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
func getX2AndX3(input int) (int, int) {
    return 2 * input, 3 * input
}
func getX2(input *int) (x2 int) {
    *input = 2* (*input)
    // return x2, x3
    return
}
```
输出结果
```go
num = 10, 2x num = 20, 3x num = 30    
num = 20, 2x num = 20, 3x num = 30 
```
### 变长参数
```go
func Greeting(prefix string, who ...string){
	//who[0]...
}
Greeting("hello:", "Joe", "Anna", "Eileen")
```
### defer
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数。多个defer语句,当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回

类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源
### 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调
### 匿名函数和闭包

### Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制
## struct
Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。
```go
type person struct {
    name string
    age int
}

var tom person
// 赋值初始化
tom.name, tom.age = "Tom", 18

// 两个字段都写清楚的初始化
bob := person{age: 25, name: "Bob"}

// 按照struct定义顺序初始化值
paul := person{"Paul", 43}
```
### 类的继承
```go
type Human struct {
    name string
    age int
    weight int
}
type Student struct {
    Human  // 匿名字段，那么默认Student就包含了Human的所有字段
    speciality string
}
```
遇到相同属性，最外层的优先访问

