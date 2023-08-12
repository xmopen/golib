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

type LoadFunc = func(param any) (any, error)

// LocalCache map + 双向链表实现 lru.
type LocalCache struct {
	expire   time.Duration // 过期时间.
	capacity int           // LRU容量.
	length   int           // LRU实际大小.
	loadFunc LoadFunc      // loadFunc 加载函数.
	nodes    *sync.Map     // nodes 节点存储.
	lock     *sync.Mutex   // 单进程同步锁.
	head     *node         // head 链表头节点.
	tail     *node         // tail 链表尾部节点.
	xlog     *xlogging.Entry
}

type node struct {
	createTime int64  // createTime 节点创建时间.
	key        string // key 节点key.
	param      any    // 节点参数.
	value      any    // 节点缓存值.

	next *node // next 下一个节点.
	pre  *node // pre 前一个节点.
}

// New 初始化本地缓存组件.
func New(loadFunc LoadFunc, cap int, expires time.Duration) *LocalCache {
	if cap <= 0 {
		cap = defaultLRUSize
	}
	localcacheInstance := &LocalCache{
		expire:   expires,
		capacity: cap,
		nodes:    &sync.Map{},
		lock:     &sync.Mutex{},
		loadFunc: loadFunc,
	}
	return localcacheInstance
}

// LoadOrCreate  从缓存中加载数据或者创建缓存.
func (l *LocalCache) LoadOrCreate(key string, param any) (any, error) {
	// 1、查看缓存中是否存在,如果缓存存在,则直接返回.
	value, isExist := l.nodes.Load(key)
	if isExist {
		curNode, ok := value.(*node)
		if !ok {
			return nil, fmt.Errorf("nodes laod value is not a node")
		}
		if !l.checkNodeIsExpire(curNode) {
			l.moveNodeToHead(curNode)
			return value, nil
		}
		val, err := l.forceUpdate(curNode)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return l.addNode(key, param)
}

// moveNodeToHead 移动当前节点到头节点.
func (l *LocalCache) moveNodeToHead(curNode *node) {
	tempHead := l.head
	tempHead.pre = curNode
	curNode.next = tempHead
	l.head = curNode
	// 更新当前节点的前驱节点以及后驱节点的后驱和前驱.
	pre := curNode.pre
	pre.next = curNode.next
}

// forceUpdate 更新节点值.
func (l *LocalCache) forceUpdate(curNode *node) (any, error) {
	// 过期重新创建.
	cacheResult, err := l.addNode(curNode.key, curNode.param)
	if err != nil {
		return nil, err
	}
	return cacheResult, nil
}

// checkNodeIsExpire 检查当前才节点存储的数据是否过期.
func (l *LocalCache) checkNodeIsExpire(curNode *node) bool {
	nowTime := time.Now().Unix()
	return nowTime >= curNode.createTime+int64(l.expire.Seconds())
}

// addNode 添加当前节点到队头.
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
	// 1.1 头节点不为空,同时更新节点前驱节点和后驱节点.
	temp := l.head
	temp.pre = curNode
	l.head = curNode
	curNode.next = temp
	// 2、是否删除尾节点.
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
