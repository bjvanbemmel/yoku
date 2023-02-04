package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	HttpGet    = "GET"
	HttpPost   = "POST"
	HttpPut    = "PUT"
	HttpPatch  = "PATCH"
	HttpDelete = "DELETE"
)

// Wrapper for the context.Context type.
// Also contains a *http.Request.
type Context struct {
	Context    context.Context
	Request    *http.Request
}

// A struct containing the information necessary
// to create a functioning route.
// Contains a user-generated URL, automatically
// generated regular expression based on the given URL
// and a callback that will be executed upon receiving requests.
type Route struct {
	URL         string
	Regex       string
	Callback    func(c *Context)
	Method      string
	Middlewares []func(c *Context) error
	params      map[int]string
}

// A struct that holds a bunch of Routes together.
type RouteGroup struct {
    Routes []*Route
}

// Slice containing pointers to every route.
// We store the references because we want these
// to be mutable (e.g. adding a middleware after generating the route).
var Routes []*Route = []*Route{}

// Wrapper for the context.Value method.
func (c *Context) Value(key string) interface{} {
	return c.Context.Value(key)
}

// Wrapper for the context.WithValue method.
func (c *Context) WithValue(ctx context.Context, key any, value any) *Context {
	c.Context = context.WithValue(ctx, key, value)

	return c
}

// Creates a new Route{} with the GET method.
// Takes a relative path and a callback.
// This callback will be executed after all middlewares (if any)
// have been executed.
func Get(path string, callback func(c *Context)) *Route {
	r := Route{
		URL:      path,
		Regex:    pathToRegex(path),
		Callback: callback,
		Method:   HttpGet,
		params:   make(map[int]string),
	}

	createParamsFromRoute(&r)
	Routes = append(Routes, &r)

	return &r
}

// Creates a new Route{} with the POST method.
// Takes a relative path and a callback.
// This callback will be executed after all middlewares (if any)
// have been executed.
func Post(path string, callback func(c *Context)) *Route {
	r := Route{
		URL:      path,
		Regex:    pathToRegex(path),
		Callback: callback,
		Method:   HttpPost,
		params:   make(map[int]string),
	}

	createParamsFromRoute(&r)
	Routes = append(Routes, &r)

	return &r
}

// Creates a new Route{} with the PUT method.
// Takes a relative path and a callback.
// This callback will be executed after all middlewares (if any)
// have been executed.
func Put(path string, callback func(c *Context)) *Route {
	r := Route{
		URL:      path,
		Regex:    pathToRegex(path),
		Callback: callback,
		Method:   HttpPut,
		params:   make(map[int]string),
	}

	createParamsFromRoute(&r)
	Routes = append(Routes, &r)

	return &r
}

// Creates a new Route{} with the DELETE method.
// Takes a relative path and a callback.
// This callback will be executed after all middlewares (if any)
// have been executed.
func Delete(path string, callback func(c *Context)) *Route {
	r := Route{
		URL:      path,
		Regex:    pathToRegex(path),
		Callback: callback,
		Method:   HttpDelete,
		params:   make(map[int]string),
	}

	createParamsFromRoute(&r)
	Routes = append(Routes, &r)

	return &r
}

// Add a separate callback function that will be called before the route's callback function.
func (r *Route) Middleware(m func (c *Context) error) *Route {
    r.Middlewares = append(r.Middlewares, m)

    return r
}

// Initialize multiple groups at once.
// Useful when adding a single Middleware to multiple routes simultaneously.
func Group(routes ...*Route) *RouteGroup {
    return &RouteGroup{
        Routes: routes,
    }
}

// Add a Middleware to a group of routes.
func (g *RouteGroup) Middleware(m func (c *Context) error) *RouteGroup {
    for _, r := range g.Routes {
        r.Middlewares = append(r.Middlewares, m)
    }

    return g
}

// Creates a map containing every dynamic
// parameter assigned to a route.
// Uses the position within a path as it's index,
// and the given {name} as the value.
func createParamsFromRoute(route *Route) {
	path := strings.Trim(route.URL, "/")
	split := strings.Split(path, "/")
	regex := regexp.MustCompile("{[A-z]+}")

	for i, param := range split {
		if regex.MatchString(param) == false {
			continue
		}

		route.params[i] = param
	}
}

// Turns a human-readable path into a regex.
//
// Example: '/user/{user}' becomes
// '^\/user\/[A-z-_0-9]+$'
func pathToRegex(path string) string {
	r := regexp.MustCompile("{[A-z]+}").ReplaceAllString(path, "[A-z-_0-9]+")
	r = strings.ReplaceAll(r, "/", "\\/")
	r = "^" + r + "$"

	return r
}

// Handles an HTTP request.
// Performs a list of actions step by step.
// Find matching route >> run middlewares >> run callback
func serve(w http.ResponseWriter, r *http.Request) {
	route, err := findRouteByRequest(r)
	if err != nil {
		panic(err)
	}

	ctx := &Context{
		Context: r.Context(),
		Request: r,
	}

	trimUrl := strings.Trim(r.URL.String(), "/")
	splitUrl := strings.Split(trimUrl, "/")

	for i, param := range route.params {
		param = strings.Trim(param, "{}")
		ctx.Context = context.WithValue(ctx.Context, param, splitUrl[i])
	}

	for _, m := range route.Middlewares {
		m(ctx)
	}

	route.Callback(ctx)
}

// Finds a route matching to an HTTP request.
func findRouteByRequest(r *http.Request) (*Route, error) {
	for _, route := range Routes {
		if r.Method != route.Method {
			continue
		}

		if match, _ := regexp.MatchString(route.Regex, r.URL.String()); match == false {
			continue
		}

		return route, nil
	}

	//TODO: Replace placeholder error with global error variable (e.g. ErrNotFound)
	return nil, errors.New("Resource not found")
}

// Creates the http server.
// Will use port 80 when no arguments have been given.
// Pass along a port within a string (e.g. "8081") to overwrite.
func Listen(args ...string) {
	var port string = "80"
	if len(args) > 0 {
		//TODO: Add proper validation to this. Right now it will take any string, instead of a properly formatted port.
		port = args[0]
	}

	fmt.Printf("Started web server on http://localhost:%v\n", port)

	http.HandleFunc("/", serve)

	//TODO: Add proper error handling to this.
	http.ListenAndServe(":"+port, nil)
}
