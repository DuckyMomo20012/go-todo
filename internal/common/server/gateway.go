package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGatewayServer(registerServer func(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	// Register gRPC server endpoint
	// NOTE: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := registerServer(ctx, mux, opts)
	if err != nil {
		return err
	}

	// Ref: https://github.com/grpc-ecosystem/grpc-gateway/issues/769#issuecomment-478307237
	r := chi.NewRouter()
	r.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		// Gateway is generated to match for /v1/ and not /api/v1
		// we could update the gateway proto to match for /api/v1 but
		// it shouldn't care where it's mounted to, hence we just rewrite the path here
		r.URL.Path = strings.ReplaceAll(r.URL.Path, "/api", "")
		mux.ServeHTTP(w, r)
	})

	// Create an HTTP server with desired timeouts
	const timeoutSeconds = 10

	address := fmt.Sprintf("%v:%v", viper.Get("HOST"), viper.Get("PORT"))

	server := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  time.Second * timeoutSeconds, // Set the read timeout to 10 seconds
		WriteTimeout: time.Second * timeoutSeconds, // Set the write timeout to 10 seconds
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.WithField("address", address).Info("Starting: HTTP Listener")

	return server.ListenAndServe()
}
