package LRUCache

type ICache interface {
	Len() int
	Cap() int
	Add(key, value interface{})
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
	Clear()
	GetAll() (map[interface{}]interface{}, List)
}
