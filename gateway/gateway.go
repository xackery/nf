package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/xackery/nf/insecure"
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
	aasv1 "github.com/xackery/nf/proto/aas/v1"
	accountsv1 "github.com/xackery/nf/proto/accounts/v1"
	adventuresv1 "github.com/xackery/nf/proto/adventures/v1"
	"github.com/xackery/nf/third_party"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string) error {

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	err = aasv1.RegisterAaServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("register aa: %w", err)
	}
	err = ranksv1.RegisterRankServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("register aa_rank: %w", err)
	}
	err = accountsv1.RegisterAccountServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("register account: %w", err)
	}
	err = adventuresv1.RegisterAdventureServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("register adventure: %w", err)
	}

	oa := getOpenAPIHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "11000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}
	// Empty parameters mean use the TLS Config specified with the server.
	if strings.ToLower(os.Getenv("SERVE_HTTP")) == "true" {
		log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
		return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
	}

	gwServer.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{insecure.Cert},
	}
	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on https://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServeTLS("", ""))
}
