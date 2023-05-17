// Package gpt
// Create  2023-03-21 00:19:36
package gpt

import (
	"fmt"
	"log"
	"testing"
)

func TestGpt(t *testing.T) {
	content, err := Do("Go实现分布式队列")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(content)
}
