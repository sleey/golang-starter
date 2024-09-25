package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	hc "github.com/sleey/golang-starter/cmd/http/huma"
	"github.com/sleey/golang-starter/config"
	"github.com/sleey/golang-starter/internal/datastore/db"
	users_handler "github.com/sleey/golang-starter/internal/handler/users"
	"github.com/sleey/golang-starter/migrations"
	"github.com/sleey/golang-starter/util"
)

func init() {
	// load .env value
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	config.InitConfig()
	log.Info().Msg("Done loading .env file and init config")

	// init log
	util.InitializeLog()
}

func main() {
	config := config.GetConfig()

	///////////////////////
	// init dependencies //
	///////////////////////

	dbClient := util.InitializeDB(config.DatabaseURL)

	mainDB := db.NewMainDB(dbClient)

	userHandler := users_handler.NewUserHandler(mainDB)

	/////////////////////
	// check for flags //
	/////////////////////

	if f.MigrationDown {
		log.Info().Msg("Running migrations down")
		migrations.Down(dbClient.DB)
		os.Exit(0)
	}

	if f.MigrationUp {
		log.Info().Msg("Running migrations up")
		migrations.Up(dbClient.DB)
		os.Exit(0)
	}

	// default is to always run migrations
	if !f.SkipMigration {
		log.Info().Msg("Running migrations up")
		migrations.Up(dbClient.DB)
	}

	if f.DryRun {
		os.Exit(0)
	}

	///////////////////////
	// start http server //
	///////////////////////

	huma.NewError = func(status int, message string, errs ...error) huma.StatusError {
		details := make([]string, len(errs))
		for i, err := range errs {
			details[i] = err.Error()
		}
		return &hc.CustomHumaError{
			Status:  status,
			Message: message,
			Details: details,
		}
	}

	cli := humacli.New(func(hooks humacli.Hooks, _ *struct{}) {
		router := chi.NewRouter()
		router.Use(middleware.RealIP)
		router.Use(cors.Default().Handler)
		router.Use(middleware.RequestID)
		router.Use(middleware.Recoverer)

		// healthcheck
		router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Initialize Huma client
		hc.InitHumaClient(router)

		// root huma router
		hr := hc.InitRouter(hc.Api, "/api")

		hr.Route("", func(hr hc.Router) {

		})

		// only called inside the cluster
		hr.Route("/users", func(hr hc.Router) {
			hc.Get(hr, "/", hc.RouterDoc{
				OperationID: "get-user-list",
				Summary:     "Get list of users",
				Description: "Get list of users",
				Tags:        []string{"users"},
			}, userHandler.GetUserList)

			hc.Get(hr, "/{id}", hc.RouterDoc{
				OperationID: "get-user",
				Summary:     "Get a user",
				Description: "Get a user",
				Tags:        []string{"users"},
			}, userHandler.GetUser)
		})

		server := http.Server{
			Addr:    ":" + config.Port,
			Handler: router,
		}

		hooks.OnStart(func() {
			server.ListenAndServe()
		})

		hooks.OnStop(func() {
			// Give the server 5 seconds to gracefully shut down, then give up.
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	log.Info().Msg("Start server on port " + config.Port)
	cli.Run()

}
