package lru

import (
	"container/list"
)

//Cache 定义了一个缓存的结构体
type Cache struct {
	maxBytes  int64
	usedBytes int64
	ll        *list.List
	cache     map[string]*list.Element
	Onevicted func(key string, value Value)
}

//entry 定义了一个缓存实体
type entry struct {
	key   string
	value Value
}

//Value 定义了一个接口，里面的函数返回该内容所占的字节大小
type Value interface {
	Len() int
}

//New 函数生成一个新的Cache对象
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		Onevicted: onEvicted,
	}
}

//Get 方法用于查找对应key值的value并将其置于链表开头
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, err := c.cache[key]; err {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, err
	}
	return
}

//RemoveOldest 移出最久没有被访问过的记录
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.usedBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.Onevicted != nil {
			c.Onevicted(kv.key, kv.value)
		}
	}
}

//Add 新增缓存/修改value
func (c *Cache) Add(key string, value Value) {
	if ele, err := c.cache[key]; err {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.usedBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.usedBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.usedBytes {
		c.RemoveOldest()
	}
}

//Len 返回链表长度
func (c *Cache) Len() int {
	return c.ll.Len()
}
