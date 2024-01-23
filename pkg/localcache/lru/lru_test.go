package lru

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestLru(t *testing.T) {
	lrucache := New(10*time.Second, 5, func(param any) (any, error) {
		number := param.(int)
		fmt.Printf("fn:[%+v]\n", number)
		return number + 100, nil
	})
	// 4 3 2 1 0
	// 放入5
	// 5 4 3 2 1 // 这个时候0就是最近最少使用，则应该剔除掉
	for i := 0; i < 6; i++ {
		res, err := lrucache.Load(strconv.Itoa(i), i)
		if err != nil {
			t.Errorf("err:[%+v]\n", err)
		}
		t.Logf("res:[%+v]", res)
	}
	// list: 5 4 3 2 1
	time.Sleep(12 * time.Second)
}

func TestLRUCacheTTL(t *testing.T) {
	lrucache := New(10*time.Second, 5, func(param any) (any, error) {
		number := param.(int)
		fmt.Printf("fn:[%+v]\n", number)
		return number, nil
	})
	res, err := lrucache.Load("1", 1)
	if err != nil {
		t.Errorf("err:[%+v]\n", err)
	}
	t.Logf("res:[%+v]", res)
	res, err = lrucache.Load("1", 1)
	if err != nil {
		t.Errorf("err:[%+v]\n", err)
	}
	t.Logf("res:[%+v]", res)
	time.Sleep(12 * time.Second)
	res, err = lrucache.Load("1", 1)
	if err != nil {
		t.Errorf("err:[%+v]\n", err)
	}
	t.Logf("res:[%+v]", res)
}
