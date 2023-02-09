package controllers

import (
	"encoding/json"
	"strconv"

	"yoku.dev/repo/cache"
	"yoku.dev/repo/router"
)

type CacheController struct{}

func (ca CacheController) Index(c *router.Context) {
	cacheList, err := json.Marshal(cache.CacheList)
	if err != nil {
		c.WriteMap(map[string]any{
			"error":   "json_malfunction",
			"message": "Could not encode JSON from CacheList.",
			"details": err.Error(),
		}, 500)

		return
	}

	c.Write(cacheList, 200)
}

func (ca CacheController) Delete(c *router.Context) {
	id, ok := c.Value("cache").(string)

	if !ok {
		c.WriteMap(map[string]any{
			"error": "resource_not_found",
		}, 404)

		return
	}

	for _, c := range cache.CacheList {
		if id, _ := strconv.Atoi(id); c.Id == id {
			c.Clear()
		}
	}

	c.WriteBool(true, 200)
}
