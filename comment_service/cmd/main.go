package main

import (
	"comment_service/config"
	pb "comment_service/genproto/comment-service"
	"comment_service/pkg/db"
	"comment_service/pkg/logger"
	"comment_service/service"
	"comment_service/service/grpc_client"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
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

	commentService := service.NewCommentService(connDB, log, grpcClient)

	//connMongDB, err := db.ConnectToMongoDB(cfg)
	//if err != nil {
	//	log.Fatal("mongo connection to mongodb error", logger.Error(err))
	//}
	//
	//commentService := service.NewCommentServiceMongo(connMongDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, commentService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
