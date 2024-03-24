package cmd

import (
	"api/internal/common/app"
	server_common "api/internal/common/server"
	"api/internal/common/server/spec"
	"context"
	"fmt"
	oapi_middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

func bootstrap() {
	app := app.NewApplication(context.Background())
	swagger, err := spec.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	payments := server_common.NewServer(app)
	paymentsStrictHandler := spec.NewStrictHandler(payments, nil)

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(oapi_middleware.OapiRequestValidator(swagger))
		r.Mount("/api", spec.HandlerFromMux(paymentsStrictHandler, r))
	})

	r.Group(
		server_common.InitRoutes(app),
	)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start application",
	Long:  `Start application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
