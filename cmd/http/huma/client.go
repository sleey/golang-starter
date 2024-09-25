package hc

import (
	huma "github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	chi "github.com/go-chi/chi/v5"
)

var Api huma.API

func InitHumaClient(router chi.Router) {
	humaConfig := huma.DefaultConfig("Backend API", "1.0.0")
	humaConfig.CreateHooks = []func(huma.Config) huma.Config{}

	// if util.Environment == "production" {
	// 	humaConfig.Servers = []*huma.Server{
	// 		{
	// 			URL: "https://backend-node.impossible.finance",
	// 		},
	// 	}
	// }

	Api = humachi.New(router, humaConfig)
}
