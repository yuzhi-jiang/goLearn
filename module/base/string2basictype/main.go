package main

import (
	"fmt"
	"strconv"
)

// string to 基本数据类型
func main() {
	//如果str不是数字，将会转为0
	var str = "hello"
	//var  intValue int64

	intValue, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("type %T, intValue = %d\n", intValue, intValue)
}
