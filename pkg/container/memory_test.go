package container

import (
	"fmt"
	"unsafe"
)

type N struct {
	i interface{}
}

type Node struct {
	N
	Age int
}

type NN struct {
	B byte
}

type Stu struct {
	Age int
}

func main() {
	// 如果用了指针,那么size应该就是直接计算的内存地址.
	size := unsafe.Sizeof(&Node{
		Age: 11,
	})
	// 占用24个字节.
	fmt.Println(size)
	siz := unsafe.Sizeof(Node{
		Age: 11,
		N: N{
			i: "1",
		},
	})
	fmt.Println(siz)
	nnn := NN{
		B: uint8(1), // 一个字节.
	}
	s := unsafe.Sizeof(nnn)
	fmt.Println(s)

	var iii int = 11
	s1 := unsafe.Sizeof(iii)
	fmt.Println(s1)

	s2 := unsafe.Sizeof(&Stu{
		Age: 11,
	})
	fmt.Println(s2)
	// 一个空的结构体8个字节
	s3 := unsafe.Sizeof(&Stu{})
	fmt.Println(s3)
}
