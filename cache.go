package cache

import "time"

type Cache struct {
	data map[string]string
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {

}

func (c Cache) Put(key, value string) {
	c.data = append()
}

func (c Cache) Keys() []string {
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
