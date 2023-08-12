package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1680239779 - time.Now().Unix()
	fmt.Println(i / (24 * 60 * 60))
	fmt.Println(24 * 60)
}
