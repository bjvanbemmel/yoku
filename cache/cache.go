package cache

import (
	"errors"
	"strconv"
)

type Cache struct {
	Id        int
	Inventory map[string]any
}

var CacheList []*Cache

func New() *Cache {
	cache := &Cache{
		Id:        len(CacheList),
		Inventory: make(map[string]any),
	}

	CacheList = append(CacheList, cache)

	return cache
}

func (c *Cache) Insert(key string, value any) (any, error) {
	c.Inventory[key] = value

	return c.Inventory[key], nil
}

func (c *Cache) Clear() error {
	c.Inventory = make(map[string]any)

	if len(c.Inventory) > 0 {
		return errors.New("Could not empty cache inventory.")
	}

	return nil
}

func (c Cache) Select(key string) (any, error) {
	item := c.Inventory[key]

	if item == nil {
		return nil, errors.New("Item not found.")
	}

	return item, nil
}

func (c Cache) SelectById(id int) (any, error) {
	key := strconv.Itoa(id)

	item := c.Inventory[key]

	if item == nil {
		return nil, errors.New("Item not found.")
	}

	return item, nil
}
