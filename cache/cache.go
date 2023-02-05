package cache

import "errors"

type Cache struct {
	Inventory map[any]any
}

func New() *Cache {
	return &Cache{
		Inventory: make(map[any]any),
	}
}

func (c *Cache) Insert(key any, value any) (any, error) {
	c.Inventory[key] = value

	return c.Inventory[key], nil
}

func (c Cache) Select(key any) (any, error) {
	item := c.Inventory[key]

	if item == nil {
		return nil, errors.New("Item not found.")
	}

	return item, nil
}
