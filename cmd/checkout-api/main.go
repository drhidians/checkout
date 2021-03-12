package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
	"github.com/spf13/pflag"

	"github.com/drhidians/checkout/pkg/checkout"
	"github.com/drhidians/checkout/pkg/checkout/checkoutdriver"
	"github.com/drhidians/checkout/pkg/pricingrules"
	"github.com/drhidians/checkout/pkg/pricingrules/pricingrulesdriver"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
	buildDate  string
)

func main() {
	flags := pflag.NewFlagSet("Checkout Pricing Rule Engine", pflag.ExitOnError)

	httpAddr := flags.String("http-addr", ":8000", "HTTP Server address")
	publicURL := flags.String("public-url", "http://localhost:8000", "Publicly available base URL")

	_ = flags.Parse(os.Args[1:])

	log.Println("starting application version", version, fmt.Sprintf("(%s)", commitHash), "built on", buildDate)

	checkoutUrl := *publicURL + "/checkout"

	router := mux.NewRouter()

	var pricingRulesService pricingrules.Service
	{
		store := pricingrules.NewInMemoryStore()
		pricingRulesService = pricingrules.New(store)
		endpoints := pricingrulesdriver.MakeEndpoints(pricingRulesService)

		pricingrulesdriver.RegisterHTTPHandlers(
			endpoints,
			router.PathPrefix("/pricing-rules").Subrouter(),
			kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
				return context.WithValue(ctx, pricingrulesdriver.ContextKeyBaseURL, checkoutUrl)
			}),
		)
	}

	{
		checkoutService := checkout.New(pricingRulesService)
		endpoints := checkoutdriver.MakeEndpoints(checkoutService)

		checkoutdriver.RegisterHTTPHandlers(
			endpoints,
			router.PathPrefix("/checkout").Subrouter(),
			kithttp.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
				return context.WithValue(ctx, pricingrulesdriver.ContextKeyBaseURL, checkoutUrl)
			}),
		)
	}

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"content-type"}),
	)

	httpServer := &http.Server{
		Addr:    *httpAddr,
		Handler: cors(router),
	}

	log.Println("listening on", *httpAddr)

	httpLn, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		log.Fatal(err)
	}


	var group run.Group

	group.Add(
		func() error { return httpServer.Serve(httpLn) },
		func(err error) { _ = httpServer.Shutdown(context.Background()) },
	)
	defer httpServer.Close()

	// Setup signal handler
	group.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

	err = group.Run()
	if err != nil {
		if _, ok := err.(run.SignalError); ok {
			log.Println(err)

			return
		}

		// Fatal error
		// We don't use fatal, so deferred functions can do their jobs.
		log.Println(err)
	}
}