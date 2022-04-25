package grcpserver

import (
	"context"
	"fmt"
	"net"

	"github.com/AlehaWP/YaPracticum.git/internal/models"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"

	pb "github.com/AlehaWP/YaPracticum.git/internal/grcpserver/proto"

	"google.golang.org/grpc"
)

var repo models.Repository
var baseURL string

// UsersServer поддерживает все необходимые методы.
type URLsServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedURLsServer
}

// AddUser реализует интерфейс добавления пользователя.
func (s *URLsServer) AddURL(ctx context.Context, in *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	var response pb.AddURLResponse

	userID := in.User
	retURL, err := repo.SaveURL(ctx, in.Url.Url, baseURL, userID)
	if err != nil {
		return nil, err
	}
	response.Url.Url = retURL

	return &response, err

}

func Start(ctx context.Context, opt models.Options) {
	sr, err := repository.NewServerRepo(ctx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()

	repo = sr
	baseURL = opt.RespBaseURL()
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
