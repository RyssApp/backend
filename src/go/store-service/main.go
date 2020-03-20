package main

import (
	"log"
	"net"

	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/repository"
	"google.golang.org/grpc"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "ryss_dev",
		Password: "lrnzt4z716lfwxa9",
		Addr:     "db-postgresql-fra1-29218-do-user-7255094-0.a.db.ondigitalocean.com:25060",
		Database: "development",
	})
	defer db.Close()

	r := repository.NewPostgresRepository(db)

	lis, err := net.Listen("tcp", ":6546")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStoreServiceServer(grpcServer, NewServer(r))
	log.Fatal(grpcServer.Serve(lis))
}
