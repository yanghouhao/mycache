package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	hash.Add("6", "4", "2")

	Cases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range Cases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yield %s, but have %s", k, v, hash.Get(k))
		}
	}
	hash.Add("8")
	hash.Add("1")
	Cases["11"] = "1"
	Cases["27"] = "8"

	for k, v := range Cases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yield %s, but have %s", k, v, hash.Get(k))
		}
	}
}
