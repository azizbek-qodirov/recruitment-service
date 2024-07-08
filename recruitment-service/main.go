package main

import (
	"log"
	"net"
	"path/filepath"
	cfg "recruitment/config"
	"recruitment/config/logger"
	er "recruitment/genproto"
	src "recruitment/services"
	post "recruitment/storage/postgres"
	"runtime"

	"google.golang.org/grpc"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	logger := logger.NewLogger(basepath, "logs/info.log")
	db, err := post.DbCon()
	if err != nil {
		log.Fatal(err)
	}

	// config := cfg.Load()
	liss, err := net.Listen("tcp", cfg.Load().HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	er.RegisterCartItemServiceServer(server, src.NewCartItemService(db, *logger))
	er.RegisterCartServiceServer(server, src.NewCartService(db))
	er.RegisterOrderRecommendationServiceServer(server, src.NewOrderRecommendationService(db))
	er.RegisterOrderServiceServer(server, src.NewOrderService(db))
	er.RegisterProductServiceServer(server, src.NewProductService(db))

	log.Println("Server is running on port :2020")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}
