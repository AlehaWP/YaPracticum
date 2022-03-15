package repository

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
)

var (
	sr        *ServerRepo
	once      sync.Once
	userID    string
	ctx       context.Context
	noConnect bool = false
)

func TestServerRepo_CreateUser(t *testing.T) {
	initVar()
	if noConnect {
		return
	}

	t.Run("CreateUser", func(t *testing.T) {
		var err error
		userID, err = sr.CreateUser(ctx)
		if len(userID) == 0 || err != nil {
			t.Errorf("ServerRepo.CreateUser() = %v, want %v", err, "some string")
		}
	})
}

func TestServerRepo_FindUser(t *testing.T) {
	initVar()
	if noConnect {
		return
	}
	type args struct {
		ctx       context.Context
		userEncID string
	}
	tests := []struct {
		name       string
		args       args
		wantFinded bool
	}{
		{
			"Тест1",
			args{
				ctx,
				"123456",
			},
			false,
		},
		{
			"Тест2",
			args{
				ctx,
				userID,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFinded := sr.FindUser(tt.args.ctx, tt.args.userEncID); gotFinded != tt.wantFinded {
				t.Errorf("ServerRepo.FindUser() = %v, want %v", gotFinded, tt.wantFinded)
			}
		})
	}
}

func TestServerRepo_GetURL(t *testing.T) {
	if noConnect {
		return
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Тест1",
			args{
				ctx,
				"123456",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sr.GetURL(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerRepo.GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ServerRepo.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCreateUser(b *testing.B) {
	if noConnect {
		return
	}
	for i := 0; i < b.N; i++ {
		sr.CreateUser(ctx)
	}
}

func BenchmarkFindUser(b *testing.B) {
	if noConnect {
		return
	}
	for i := 0; i < b.N; i++ {
		sr.FindUser(ctx, userID)
	}
}

func BenchmarkGetURL(b *testing.B) {
	if noConnect {
		return
	}
	for i := 0; i < b.N; i++ {
		sr.GetURL(ctx, "123")
	}
}

func initVar() {
	once.Do(func() {
		ctx = context.Background()
		opt := defoptions.NewDefOptions()
		var err error
		sr, err = NewServerRepo(ctx, opt.DBConnString())

		if err != nil {
			noConnect = true
			fmt.Println("Ошибка при подключении к БД: ", err)
			return
		}
	})
}
