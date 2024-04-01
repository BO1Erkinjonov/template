package main

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
	"user_service/config"
	pb "user_service/genproto/user-service"
	"user_service/pkg/db"
	"user_service/pkg/logger"
	"user_service/sevice"
	"user_service/sevice/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "test_server")
	defer logger.Cleanup(log)

	log.Info("main sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("Error while dealing", logger.Error(err))
	}

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := sevice.NewUserService(connDB, log, grpcClient)

	//connMongDB, err := db.ConnectToMongoDB(cfg)
	//if err != nil {
	//	log.Fatal("mongo connection to mongodb error", logger.Error(err))
	//}
	//userService := sevice.NewUserServiceMongo(connMongDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)

	if err != nil {
		log.Fatal("error", logger.Error(err))
	}

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
