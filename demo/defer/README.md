defer 特性

1. 延迟调用
2. 传递参数如果是值传递，则传递的是值拷贝，如果是引用传递，则传递的是引用拷贝
   - 值传递：在传入后固定，不受后面代码影响
   - 引用传递：在传入后，虽然会延迟执行，但依然会受后面代码影响


### 示例

```go

package main

import "fmt"

type student struct {
	age int
}

func main() {
	stud1 := student{1}
	stud1.age = 2
	defer printVar(&stud1)
	stud1.age = 3
	fmt.Println(stud1.age)

}

func printVar(a *student) {
	//print address
	println("before:", a.age)
	a.age = 4
	println("after:", a.age)
}


```
### output

```
3
before: 3
after: 4
```
