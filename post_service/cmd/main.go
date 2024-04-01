package main

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
	"post_service/config"
	pb "post_service/genproto/post-service"
	"post_service/pkg/db"
	"post_service/pkg/logger"
	"post_service/service"
	"post_service/service/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post_service")
	defer logger.Cleanup(log)

	log.Info("main sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	grpcClient, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("grpc client dial error")
	}

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	postService := service.NewPostService(connDB, log, grpcClient)

	//connMongDB, err := db.ConnectToMongoDB(cfg)
	//if err != nil {
	//	log.Fatal("mongo connection to mongodb error", logger.Error(err))
	//}
	//
	//postService := service.NewPostServiceMongo(connMongDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)

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
