// Package localcache 本地缓存组件实现.
package localcache

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/xmopen/golib/pkg/xlogging"
)

// defaultLRUSize 默认LRU大小.
const defaultLRUSize = 1 << 1

// LoadFunc LocalCache node load func
type LoadFunc = func(param any) (any, error)

// LocalCache of LRU algorithm with complexity O(1) realized by map + linked list
type LocalCache struct {
	expire   time.Duration // 过期时间.
	capacity int           // LRU容量.
	length   int           // LRU实际大小.
	loadFunc LoadFunc      // loadFunc 加载函数.
	nodes    *sync.Map     // nodes 节点存储.
	lock     *sync.Mutex   // lock 单进程同步锁.
	head     *node         // head 链表头节点.
	tail     *node         // tail 链表尾部节点.
	xlog     *xlogging.Entry
}

type node struct {
	createTime int64  // createTime node create time.
	key        string // key node key.
	param      any    // param node func param.
	value      any    // value node cache value.
	next       *node  // next node.
	pre        *node  // pre node.
}

// New init LocalCache instance.
func New(loadFunc LoadFunc, cap int, expires time.Duration) *LocalCache {
	if cap <= 0 {
		cap = defaultLRUSize
	}
	localCacheInstance := &LocalCache{
		expire:   expires,
		capacity: cap,
		nodes:    &sync.Map{},
		lock:     &sync.Mutex{},
		loadFunc: loadFunc,
	}
	return localCacheInstance
}

// LoadOrCreate  The cache is created if it does not exist,
// and returned from the cache if it does.
func (l *LocalCache) LoadOrCreate(key string, param any) (any, error) {
	value, isExist := l.nodes.Load(key)
	if isExist {
		curNode, ok := value.(*node)
		if !ok {
			return nil, fmt.Errorf("nodes laod value is not a node")
		}
		if !l.checkNodeIsExpire(curNode) {
			l.moveNodeToHead(curNode)
			return curNode.value, nil
		}
		val, err := l.forceUpdate(curNode)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return l.addNode(key, param)
}

// moveNodeToHead move node to head.
func (l *LocalCache) moveNodeToHead(curNode *node) {
	tempHead := l.head
	tempHead.pre = curNode
	curNode.next = tempHead
	l.head = curNode
	// update currNode.pre and currNode.next
	pre := curNode.pre
	pre.next = curNode.next
}

// forceUpdate force update currNode value.
func (l *LocalCache) forceUpdate(curNode *node) (any, error) {
	cacheResult, err := l.addNode(curNode.key, curNode.param)
	if err != nil {
		return nil, err
	}
	return cacheResult, nil
}

// checkNodeIsExpire check currNode is expired.
func (l *LocalCache) checkNodeIsExpire(curNode *node) bool {
	nowTime := time.Now().Unix()
	return nowTime >= curNode.createTime+int64(l.expire.Seconds())
}

// addNode add currNode to head.
func (l *LocalCache) addNode(key string, param any) (any, error) {
	val, err := l.loadFunc(param)
	if err != nil {
		return nil, err
	}
	curNode := &node{
		createTime: time.Now().Unix(),
		key:        key,
		param:      param,
		value:      val,
	}
	// 1、添加到头节点.
	if l.head == nil {
		l.head = curNode
		l.tail = l.head
		l.addLength()
		return val, nil
	}
	//  头节点不为空,同时更新节点前驱节点和后驱节点.
	temp := l.head
	temp.pre = curNode
	l.head = curNode
	curNode.next = temp
	if l.length+1 > l.capacity && l.tail != nil {
		l.nodes.Delete(l.tail.key)
		tailPre := l.tail.pre
		l.tail = tailPre
		// 删除的节点GC回收.
		runtime.GC()
	} else {
		l.addLength()
	}
	l.nodes.Store(curNode.key, curNode)
	return val, nil
}

func (l *LocalCache) addLength() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.length++
}

// ClearAllCache clear all cache
func (l *LocalCache) ClearAllCache() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.nodes.Range(func(key, value any) bool {
		l.nodes.Delete(key)
		return true
	})
	l.length = 0
	l.head = nil
	l.tail = nil
}
