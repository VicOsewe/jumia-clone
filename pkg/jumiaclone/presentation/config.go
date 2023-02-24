package presentation

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	serverTimeoutSeconds = 120
)

// Router sets up the ginContext router
func Router(ctx context.Context) (*mux.Router, error) {
	r := mux.NewRouter()
	// h := InitHandlers()

	r.Path("/health").HandlerFunc(HealthStatusCheck)
	// RESTRoutes := r.PathPrefix("/api/v1").Subrouter()

	return r, nil
}

// HealthStatusCheck endpoint to check if the server is working.
func HealthStatusCheck(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(true)
	if err != nil {
		log.Fatal(err)
	}
}

// InitHandlers initializes all the handlers dependencies
func InitHandlers() {

}

// PrepareServer prepares the http server
func PrepareServer(ctx context.Context, port int) *http.Server {
	r, err := Router(ctx)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}

	addr := fmt.Sprintf(":%d", port)
	h := handlers.CompressHandlerLevel(r, gzip.BestCompression)

	h = handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST"}),
	)(h)
	h = handlers.CombinedLoggingHandler(os.Stdout, h)
	h = handlers.ContentTypeHandler(
		h,
		"application/json",
		"application/x-www-form-urlencoded",
	)
	return &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: serverTimeoutSeconds * time.Second,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
	}
}
