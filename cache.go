package cache

import "time"

type Cache struct {
	values map[string]kv
}

type kv struct {
	value    string
	expiry   bool
	deadline time.Time
}

func NewCache() Cache {
	var cache Cache
	cache.values = make(map[string]kv)
	return cache
}

func (c Cache) Get(key string) (string, bool) {
	return c.values[key].value, c.values[key].expiry
}

func (c Cache) Put(key, value string) {
	c.values[key] = kv{value: value, expiry: false}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.values))
	for k, v := range c.values {
		if v.expiry == false || (v.expiry == true && v.deadline.After(time.Now())) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c Cache) PutTill(key string, value string, deadline time.Time) {
	c.values[key] = kv{value: value, expiry: true, deadline: deadline}
}
