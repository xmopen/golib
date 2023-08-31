package commonutil

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Println(MD5("AFAFA"))
	fmt.Println(MD5("AFAFAA"))
	fmt.Println(MD5("AFAFAAX"))
	fmt.Println(MD5("AFAFFA"))
}
