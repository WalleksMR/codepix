package main

import (
	"github/walleksmr/codepix/application/grpc"
	"github/walleksmr/codepix/infra/db"
	"os"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
