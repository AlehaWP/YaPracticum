package main

import (
	"context"
	"fmt"
	"net"

	"github.com/AlehaWP/YaPracticum.git/cmd/grcp_server/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/cmd/grcp_server/internal/models"
	"github.com/AlehaWP/YaPracticum.git/cmd/grcp_server/internal/repository"

	pb "github.com/AlehaWP/YaPracticum.git/cmd/grcp_server/proto"

	"google.golang.org/grpc"
)

var Repo models.Repository
var BaseURL string

// UsersServer поддерживает все необходимые методы.
type URLsServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedURLsServer
}

// AddUser реализует интерфейс добавления пользователя.
func (s *URLsServer) AddYRL(ctx context.Context, in *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	var response pb.AddURLResponse

	userID := in.User
	retURL, err := Repo.SaveURL(ctx, in.Url.Url, BaseURL, userID)
	if err != nil {
		return nil, err
	}
	response.Url.Url = retURL

	return &response, err

}

func main() {
	var err error
	opt := defoptions.NewDefOptions()
	ctx := context.Background()
	Repo, err = repository.NewServerRepo(ctx, opt.DBConnString())
	BaseURL = opt.RespBaseURL()
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		fmt.Println(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterURLsServer(s, &URLsServer{})

	fmt.Println("сервер gRPC начал работу")
	// получаем запрос gRpc
	if err := s.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
