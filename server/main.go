package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"github.com/asa-taka/practice-grpc-blog/server/services/account"
	"github.com/asa-taka/practice-grpc-blog/server/services/document"
)

var (
	port    = flag.Int("port", 10000, "gRPC server port")
	webPort = flag.Int("web-port", 10001, "gRPC-Web server port")
)

func main() {
	flag.Parse()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logUnaryInterceptor),
	)

	// Register service implementations
	account.RegisterNewServer(grpcServer)
	document.RegisterNewServer(grpcServer)

	reflection.Register(grpcServer) // for grpc_cli

	// Run both gRPC/gRPC-Web
	go runGrpcWebServer(grpcServer, *webPort)
	runGrpcServer(grpcServer, *port)
}

func runGrpcServer(grpcServer *grpc.Server, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpclog.Infof("gRPC server starts on localhost:%d", port)
	grpcServer.Serve(lis)
}

func runGrpcWebServer(grpcServer *grpc.Server, port int) {
	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Infof("gRPC-Web server starts on http://localhost:%d", port)
	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}

// gRPC Interceptor
// ----------------

func logUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	res, err = handler(ctx, req)
	method := info.FullMethod
	if err == nil {
		grpclog.Infof("%s: OK", method)
		if grpclog.V(1) {
			grpclog.Infof("%s Request: %v", method, req)
			grpclog.Infof("%s Response: %v", method, res)
		}
	} else {
		grpclog.Errorf("%s: %v", method, err)
		if grpclog.V(1) {
			grpclog.Infof("%s Request: %v", method, req)
		}
	}
	return res, err
}
