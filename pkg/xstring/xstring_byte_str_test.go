package xstring

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	str := "Hello Word"
	data := []byte(str)
	result := ToString(data)
	fmt.Printf("result type:%T res:[%+v]\n", result, result)
	d := ToBytes(result)
	fmt.Printf("result type:%T res:[%+v]\n", d, string(d))
}
