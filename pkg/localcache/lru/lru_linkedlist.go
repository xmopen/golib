package lru

import (
	"sync"
	"time"

	"github.com/xmopen/golib/pkg/container/linkedlist"
)

// LRU: 最近最少未使用淘汰算法
// 增加：头节点增加
// 删除：按照制定key删除
// 查询：按照制定Key进行查询，每查询一次将其移动到头节点

// DefaultLocalCacheCapacity 默认LRUCache大小
const DefaultLocalCacheCapacity = 1024

// LRULocalCacheFunc LRU 缓存加载函数类型
type LRULocalCacheFunc func(param any) (any, error)

// LocalCache LRU LocalCache
type LocalCache struct {
	lock     *sync.RWMutex
	ttl      time.Duration
	capacity int
	list     linkedlist.IDoubleLinkedList
	mapping  *sync.Map
	fn       LRULocalCacheFunc
}

// New 初始化LRU LocalCache
func New(ttl time.Duration, capacity int, fn LRULocalCacheFunc) *LocalCache {
	if capacity <= 0 {
		capacity = DefaultLocalCacheCapacity
	}
	return &LocalCache{
		lock:     &sync.RWMutex{},
		ttl:      ttl,
		capacity: capacity,
		list:     linkedlist.NewDoubleLinkedList(),
		mapping:  &sync.Map{},
		fn:       fn,
	}
}

// Load 从缓存中加载数据
func (l *LocalCache) Load(key string, param any) (any, error) {
	val, isExist := l.mapping.Load(key)
	var cacheNode *node
	if isExist {
		if val != nil {
			cacheNode = val.(*node)
		}
	}

	if cacheNode != nil && !cacheNode.isExpire() {
		l.flushNodeCacheToHead(cacheNode)
		return cacheNode.item, nil
	}

	l.lock.RLock()
	defer l.lock.RUnlock()
	// remove cache node
	l.removeNodeCache(cacheNode)
	return l.addNode(key, param)
}

func (l *LocalCache) flushNodeCacheToHead(cacheNode *node) {
	l.list.RemoveWithValue(cacheNode.item)
	l.list.PushWithHead(cacheNode.item)
}

func (l *LocalCache) removeNodeCache(cacheNode *node) {
	if cacheNode == nil {
		return
	}
	l.list.RemoveWithValue(cacheNode.item)
}

func (l *LocalCache) addNode(key string, param any) (any, error) {
	item, err := l.fn(param)
	if err != nil {
		return nil, err
	}
	if l.list.Size() >= l.capacity {
		l.list.RemoveFromTail()
		l.mapping.Delete(key)
	}
	// 缓存为空则不应该缓存
	// 缓存的属性是在保证缓存一致性的前提下加速数据访问(降低DB访问),同时又是为了避免流量突发的时带来的DB压力
	// 所以如果为空也进行缓存
	cacheNode := newNode(item, l.ttl)
	l.mapping.Store(key, cacheNode)
	l.list.PushWithHead(cacheNode.item)
	return item, nil
}

func (l *LocalCache) ClearAll() {
	l.lock.RLock()
	defer l.lock.RUnlock()
	l.list = linkedlist.NewDoubleLinkedList()
	l.mapping = &sync.Map{}
}
