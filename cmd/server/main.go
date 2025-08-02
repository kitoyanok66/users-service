package main

import (
	"log"

	"github.com/kitoyanok66/users-service/internal/database"
	transportgrpc "github.com/kitoyanok66/users-service/internal/transport/grpc"
	"github.com/kitoyanok66/users-service/internal/user"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	repo := user.NewUserRepository(db)
	svc := user.NewUserService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("Users gRPC server error: %v", err)
	}
}
