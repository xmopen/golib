package localcache

// ILocalCache LocalCache interface
type ILocalCache interface {
	Load(key string, param any) (any, error)
	ClearAll()
}
