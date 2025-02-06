package main

import (
	"fmt"
	"flag"
)

func main() {
	//定义名为intVal的int型的变量占位, 解析后, 利用CLI参数intVal 进行填充
	var intVal = flag.Int("intVal", 0, "int 类型参数")
	var boolVal = flag.Bool("boolVal", false, "bool 类型参数")
	var stringVal = flag.String("stringVal", "", "string 类型参数")
	
	flag.Parse()
	fmt.Println("-intVal", *intVal)
	fmt.Println("-boolVal",*boolVal)
	fmt.Println("-stringVal", *stringVal)
}
