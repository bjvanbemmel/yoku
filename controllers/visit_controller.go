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

func (v VisitController) Index(c *router.Context) {
    var visits []models.Visit

    Paginate(c.QueryInt("page"), 5000).Preload("VisitPath").Order("id DESC").Find(&visits)

    c.WriteMap(map[string]any {
        "data": visits,
    }, 200)
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

		Db.FirstOrCreate(&vp, models.VisitPath{Path: vb.URL})

		vpi, _ = Cache.Insert(vb.URL, vp)
	}

	vp, ok := vpi.(models.VisitPath)

	if !ok {
		c.WriteMap(map[string]any{
			"error":   "cache_malfunction",
			"message": "Something went wrong while retrieving from cache.",
		}, 500)

        return
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
