package mycache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGet(t *testing.T) {
	loadcount := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[slowDB] search key: ", key)
		if v, ok := db[key]; ok {
			if _, ok := loadcount[key]; ok {
				loadcount[key] = 0
			}
			loadcount[key]++
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exists ", key)
	}))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get value")
		}

		if _, err := gee.Get(k); err != nil || loadcount[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := gee.Get("yhh"); err == nil {
		t.Fatalf("the value of yhh should be empty, but %s got", view)
	}
}
