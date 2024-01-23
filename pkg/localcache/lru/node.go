package lru

import "time"

type node struct {
	item           any
	ttl            time.Duration
	lastActiveTime time.Time
}

// newNode 初始化node
func newNode(item any, ttl time.Duration) *node {
	return &node{
		item:           item,
		ttl:            ttl,
		lastActiveTime: time.Now(),
	}
}

// isExpire 判断node是否过期
func (n *node) isExpire() bool {
	return time.Since(n.lastActiveTime).Seconds() > n.ttl.Seconds()
}

func (n *node) flushLastActiveTime() {
	n.lastActiveTime = time.Now()
}
