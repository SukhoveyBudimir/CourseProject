package Game

import (
	"context"
	"fmt"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/player"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/repository"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/server"
	"github.com/SukhoveyBudimir/CourseProject/Game/internal/service"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

var (
	poolP pgxpool.Pool
	//poolM mongo.Client
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :50051...")
	key := []byte("super-key")
	cfg := player.Config{JwtKey: key}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	conn := DBConnection(&cfg)
	fmt.Println("DB successfully connect...")
	/*defer func() {
		poolP.Close()
		if err = poolM.Disconnect(context.Background()); err != nil {
			log.Errorf("cannot disconnect with mongodb")
		}
	}()*/
	ns := grpc.NewServer()
	newService := service.NewService(conn, cfg.JwtKey)
	srv := server.NewServer(newService)
	pb.RegisterPharmacyplayerServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

// DBConnection create connection with db
func DBConnection(cfg *player.Config) repository.Repository {
	log.Info(cfg.PostgresDBURL)
	poolP, err := pgxpool.Connect(context.Background(), cfg.PostgresDBURL) //"postgres://postgres:player@postgres:5432/player?sslmode=disable"
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
		return nil
	}
	return &repository.PRepository{Pool: poolP}
}
