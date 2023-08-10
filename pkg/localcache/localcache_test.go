package localcache

import (
	"fmt"
	"testing"
	"time"
)

type Student struct {
	name string
}

func TestLocalCache(t *testing.T) {
	// 构建测试本地缓存长度为3，有效期为1分钟.
	localCache := New(func(param any) (any, error) {
		name, ok := param.(string)
		if !ok {
			return nil, fmt.Errorf("parma not string")
		}
		return &Student{
			name: name,
		}, nil
	}, 3, 1*time.Minute)
	test := []struct {
		key   string
		param any
		want  string
	}{
		{
			key:   "1",
			param: "1",
			want:  "1",
		},
		{
			key:   "2",
			param: "2",
			want:  "2",
		},
		{
			key:   "3",
			param: "3",
			want:  "3",
		},
		{
			key:   "4",
			param: "4",
			want:  "4",
		},
	}
	for _, item := range test {
		itr, err := localCache.LoadOrCreate(item.key, item.param)
		if err != nil {
			t.Errorf("item:[%+v] err:[%+v]", item, err)
			continue
		}
		val, ok := itr.(*Student)
		if !ok {
			t.Errorf("cache obj not student item:[%+v]", item)
			continue
		}
		if val.name != item.want {
			t.Errorf("cache obj not want,cache:[%+v] want:[%+v]", val, item.want)
		}
	}
	time.Sleep(1 * time.Minute)
	// 缓存更新之后未更新到nodes中.
	_, err := localCache.LoadOrCreate(test[1].key, test[1].param)
	if err != nil {
		t.Errorf("err:[%+v]", err)
	}
	// 数据过期之后有一点点问题.
}