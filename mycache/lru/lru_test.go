package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1 failed\n")
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("there is no key2 here\n")
	}
}

func TestAdd(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k"
	v1, v2, v3 := "1234", "234556", "3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get(k1); ok || lru.ll.Len() != 2 {
		t.Fatalf("Remove the LRU failed\n")
	}
}

func TestOnevicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback)
	lru.Add("key1", String("value1"))
	lru.Add("e1", String("va"))
	lru.Add("y1", String("e1"))
	lru.Add("11", String("a1"))

	expect := []string{"key1", "e1"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("call onevictd function failed, expected key equal to %s", expect)
	}
}
