# headFirstGo

[official website](https://headfirstgo.com/)

# chapter2
```go
1.字符串
  1.1 字符串字面量由双引号("")包围,但是rune字面量由单引号(')包围。
  1.2 Go语言使用Unicode标准来存储rune,rune被保存为数字代码,而不是字符本身。
  1.3 比如 'A'输出的是unicode字符代码。

2.获取类型
  reflect.TypeOf("str")

3.变量/函数/类型
  3.1 名称必须以字母开头,可以有任意数量的额外的字母和数字。
  3.2 名称大写字母开头:可以导出,小写字母开头:只能在当前包中使用。
  3.3 如果一个名称由多个单词组成,之后的每个单词都应该首字母大写,并且应该连在一起
      比如: topPrice(需要从包中导出时 才需要首字母大写。)
  3.4 正常情况下 同一变量名在同一作用域声明两次 会出现编译错误,
      但是只要短变量声明中 至少有一个变量名是新的,是被允许的。
      新变量名视为声明,老的视为赋值。
      a := 1
      a, b := 2, 3
      fmt.Println(a, b)

4.Go工具
  go build
  go run
  go fmt // 使用go标准格式重新格式化源文件
```

# chapter3
```go
1.基础简介
  %f  // 二进制
  %d  // 十进制
  %t  // 布尔值(ture or false)
  %v  // 根据所提供值的类型 选择适当的格式
  %#v // 按其在Go程序代码中显示的格式进行格式化
  %T  // 所提供值的类型
  %%  // 一个完完全全的百分号

2.格式化值宽度
  如果没有匹配最小宽度 则使用空格进行填充
  fmt.Printf("%12s|%3d\n", "Stamps", 50)

3.格式化小数宽度测试
  fmt.Printf("%%7.3f: %7.3f\n", 3.1415926) // 7.3表示最小宽度,小数点后面的表示四舍五入到几位。
  fmt.Printf("%%.1f: %.1f\n", 3.1415926)   // .1f表示四舍五入到1位 无填充。

4.函数
  当return语句运行时 函数立即退出。
  Go是一种"值传递语言,函数形参从函数调用中接收实参的副本。"

5.指针类型
  *int

  // 获取变量指针
  var a int
  &a 
```

# chapter4
```go
1.包
  - 子目录名与包名应该相同

2.包命名规范
 .使用包的开发人员每次调用包中的函数时都需要输入包名。(想想fmt.Printf/fmt.Println/fmt.Print等)为了尽可能地使之不那么痛苦,包名应该遵循以下几个规则:
   ·包名应全部为小写。
   ·如果含义相当明显，名称应该缩写(如fmt)。
   ·如果可能的话，应该是一个词。如果需要两个词，不应该用下划 线分隔，第二个词也不应该大写。(strconv包就是一个例子。)
   ·导入的包名可能与本地变量名冲突，所以不要使用包用户可能也想使用的名称。(例如，如果fmt包被命名为format，那么导入该包的任何人如果把一个局部变量也命名为format，则将面临冲突的风险)。   

3.command
   go doc strconv // 获取帮助信息
   go doc strconv ParseFloat // 获取strconv包中的ParseFloat函数帮助信息
   godoc -http=:80
```

# chapter(5-6)
```go
1.slice
  stu := make([]int,5) // 创建一个5个整数切片,并创建一个变量保存他。

2.os.Args
  arguments := os.Args[1:] // 收个参数为脚本文件本身

3.可变长函数参数
  func SeveralInts(numbers ...int) { // 是一个切片
	fmt.Println(numbers)
  }

  intSlice := []int{1,2,3}
  autoFunc.SeveralInts(intSlice...) //使用int切片代替可变参数
```

# chapter7
```go
1.map
  var myMap map[string]int
  myMap = make(map[string]int)

  or
  myMap := make(map[string]int) // 创建一个映射并声明一个用于保存它的变量。

2.如果声明了一个映射变量未赋值 它的值是nil
3.判断映射是否存在
  counters := map[string]int{"a": 3, "b": 4}

  // 删除一个元素
  delete(counters,"b")

  valueB, ok := counters["b"]
  if ok {
	fmt.Println(valueB)
  }

4. for...range 打印出的结果是无序的 如果有序 sort.Strings(name) 
```
