package grcpserver

import (
	"context"
	"fmt"
	"net"

	"github.com/AlehaWP/YaPracticum.git/internal/models"

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

func Start(ctx context.Context, sr models.Repository, opt models.Options) {
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
	// получаем запрос gRpc
	go s.Serve(listen)

	<-ctx.Done()
	s.Stop()
}
