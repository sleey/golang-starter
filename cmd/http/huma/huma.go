package hc

import (
	"context"
	"net/http"
	"strings"

	huma "github.com/danielgtaylor/huma/v2"
)

// Router is a wrapper around huma.Router
type Router struct {
	Api         huma.API
	Prefix      string
	Middlewares huma.Middlewares
}

type RouterDoc struct {
	OperationID string
	Summary     string
	Description string
	Tags        []string
	Hidden      bool
}

func BaseHandler[I, O any](r Router, method string, path string, op RouterDoc, fn func(context.Context, *I) (*O, error)) {
	// remove trailing slash
	fullPath := strings.TrimRight(r.Prefix+path, "/")

	// if util.Environment == "" {
	// 	log.Println(fullPath)
	// }

	huma.Register(r.Api, huma.Operation{
		OperationID: op.OperationID,
		Method:      method,
		Path:        fullPath,
		Summary:     op.Summary,
		Description: op.Description,
		Tags:        op.Tags,
		Middlewares: r.Middlewares,
		Hidden:      op.Hidden,
	}, fn)
}

func Get[I, O any](r Router, path string, op RouterDoc, fn func(context.Context, *I) (*O, error)) {
	BaseHandler(r, http.MethodGet, path, op, fn)
}

func Post[I, O any](r Router, path string, op RouterDoc, fn func(context.Context, *I) (*O, error)) {
	BaseHandler(r, http.MethodPost, path, op, fn)
}

func Delete[I, O any](r Router, path string, op RouterDoc, fn func(context.Context, *I) (*O, error)) {
	BaseHandler(r, http.MethodDelete, path, op, fn)
}

func Patch[I, O any](r Router, path string, op RouterDoc, fn func(context.Context, *I) (*O, error)) {
	BaseHandler(r, http.MethodPatch, path, op, fn)
}

func InitRouter(api huma.API, prefix string) Router {
	return Router{
		Api:    api,
		Prefix: prefix,
	}
}

func (r *Router) Use(middlewares huma.Middlewares) {
	r.Middlewares = append(r.Middlewares, middlewares...)
}

func (r *Router) Route(prefix string, fn func(Router)) {
	temp := *r
	temp.Prefix += prefix

	fn(temp)
}

func (r *Router) Group(fn func(Router)) {
	fn(*r)
}
