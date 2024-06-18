package tree

// Tree: 暂时先不实现吧，不确定具体的实现场景。

// ITree 提供树的接口.
type ITree interface {
	Size() int
	Put(item any)
	Get(item any)
	Remove(item any)
	Range(func())
}
