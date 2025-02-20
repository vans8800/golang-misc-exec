package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func main() {
	// 从对象池中获取一个 bytes.Buffer 对象
	buffer := bufferPool.Get().(*bytes.Buffer)

	// 使用获取的对象
	buffer.WriteString("Hello, ")
	buffer.WriteString("World!")

	// 输出结果
	fmt.Println(buffer.String()) // 输出: Hello, World!

	// 清空对象以便重用
	buffer.Reset()

	// 继续使用同一个对象
	buffer.WriteString("Reusing buffer.")
	fmt.Println(buffer.String()) // 输出: Reusing buffer.

	// 将对象放回对象池中
	bufferPool.Put(buffer)

	// 从对象池中获取另一个 bytes.Buffer 对象
	anotherBuffer := bufferPool.Get().(*bytes.Buffer)
	anotherBuffer.WriteString("New buffer from pool.")
	fmt.Println(anotherBuffer.String()) // 输出: New buffer from pool.

	// 清理并放回对象池
	anotherBuffer.Reset()
	bufferPool.Put(anotherBuffer)
}
