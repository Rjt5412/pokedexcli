package pokecache

import (
	"fmt"
	"testing"
	"time"
)

const interval = time.Millisecond * 10

func TestCreateCache(t *testing.T) {
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {

			cache := NewCache(interval)

			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)

			if !ok {
				t.Errorf("Expected to find the key")
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("expected to find the value")
				return
			}
		})
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)

	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}

}

func TestReapFail(t *testing.T) {
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)

	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}

}
