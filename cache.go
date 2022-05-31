package cache

import "time"

type Cache struct {
	data     map[string]string
	datatime map[string]time.Time
}

func NewCache() Cache {
	var data = make(map[string]string, 1)
	var datatime = make(map[string]time.Time, 1)
	return Cache{data, datatime}
}

func (c Cache) Get(key string) (string, bool) {
	if time.Now().After(c.datatime[key]) && !c.datatime[key].IsZero() {
		delete(c.data, key)
		delete(c.datatime, key)
	}
	v, ok := c.data[key]
	return v, ok
}

func (c Cache) Put(key, value string) {
	c.data[key] = value
	c.datatime[key] = time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
}

func (c Cache) Keys() []string {
	var keys = make([]string, 1)
	for k, _ := range c.data {
		if time.Now().Before(c.datatime[k]) || c.datatime[k].IsZero() {
			keys = append(keys, k)
		} else if time.Now().After(c.datatime[k]) && !c.datatime[k].IsZero() {
			delete(c.data, k)
			delete(c.datatime, k)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.datatime[key] = deadline
}
