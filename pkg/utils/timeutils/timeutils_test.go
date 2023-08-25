package timeutils

import (
	"fmt"
	"testing"
)

func TestStringTimeToCNTime(t *testing.T) {
	str := "2021-01-11 01:14:53.377"
	res, err := StringTimeToCNTime(str)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
