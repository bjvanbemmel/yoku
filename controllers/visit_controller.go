package controllers

import (
	"encoding/json"
	"net"
	"net/http"

	"yoku.dev/repo/cache"
	. "yoku.dev/repo/database"
	"yoku.dev/repo/models"
	"yoku.dev/repo/router"
)

type VisitController struct{}

var Cache *cache.Cache = cache.New()

type VisitBody struct {
	URL string
}

func (v VisitController) Create(c *router.Context) {
	var vb VisitBody
	json.NewDecoder(c.Request.Body).Decode(&vb)

	if vb.URL == "" {
		c.WriteMap(map[string]any{
			"error":   "missing_field",
			"message": "The `url` field must be filled.",
		}, 422)

		return
	}

	ip, err := getIp(c.Request)
	if err != nil {
		c.WriteMap(map[string]any{
			"error": "Could not parse IP address.",
		}, 500)

		return
	}

	agent := c.Request.UserAgent()

    var vpi interface{}
    vpi, err = Cache.Select(vb.URL)

    if err != nil {
        vp := models.VisitPath{
            Path: vb.URL,
        }
        Db.Create(&vp)

        vpi, _ = Cache.Insert(vb.URL, vp)
    }

    vp, ok := vpi.(models.VisitPath)

    if !ok {
        c.WriteMap(map[string]any{
            "error": "cache_malfunction",
            "message": "Something went wrong while retrieving from cache.",
        }, 500)
    }

    Db.Create(&models.Visit{
        UserAgent: agent,
        VisitPath: vp,
        IP:        ip,
    })
}

func getIp(c *http.Request) (string, error) {
	var ip string
	var err error

	if ip = c.Header.Get("X-Real-IP"); ip == "" {
		ip, _, err = net.SplitHostPort(c.RemoteAddr)
	}

	return ip, err
}

func addVisitPathToCache(vp models.VisitPath) (*models.VisitPath, error) {


    return nil, nil
}
