package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-microfrontend/items-repository/internal/repository"
	"github.com/go-microfrontend/items-repository/internal/server"
	pb "github.com/go-microfrontend/items-repository/pkg/items"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	list, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		slog.Error("failed to listen", slog.String("error", err.Error()))
		os.Exit(1)
	}

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repo := repository.New(conn)
	srv := server.New(repo)

	server := grpc.NewServer()
	pb.RegisterItemServiceServer(server, srv)

	reflection.Register(server)

	slog.Info("gRPC server started", slog.String("address", list.Addr().String()))

	if err := server.Serve(list); err != nil {
		slog.Error("failed to serve", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
