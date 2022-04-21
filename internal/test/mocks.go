package test

import (
	"context"

	"github.com/AlehaWP/YaPracticum.git/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (m *RepoMock) SaveURL(ctx context.Context, url, baseURL, userID string) (string, error) {
	args := m.Called(ctx, url, baseURL, userID)
	return args.String(0), args.Error(1)
}

func (m *RepoMock) SaveURLs(ctx context.Context, urls map[string]string, baseURL, userID string) (map[string]string, error) {
	args := m.Called(ctx, urls, baseURL, userID)
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *RepoMock) GetURL(ctx context.Context, id string) (string, error) {
	args := m.Called(ctx, id)
	return args.String(0), args.Error(1)
}

func (m *RepoMock) FindUser(ctx context.Context, id string) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *RepoMock) CreateUser(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *RepoMock) GetUserURLs(context.Context, string) ([]models.URLs, error) {
	return nil, nil
}

func (m *RepoMock) CheckDBConnection(context.Context) error {
	return nil
}

func (m *RepoMock) SetURLsToDel(context.Context, []string, string) error {
	return nil
}

func (m *RepoMock) GetStatistics(context.Context) (models.Statistics, error) {
	return models.Statistics{}, nil
}

type OptsMock struct {
	mock.Mock
}

func (o *OptsMock) ServAddr() string {
	args := o.Called()
	return args.String(0)
}

func (o *OptsMock) RespBaseURL() string {
	args := o.Called()
	return args.String(0)
}

func (o *OptsMock) RepoFileName() string {
	args := o.Called()
	return args.String(0)
}

func (o *OptsMock) DBConnString() string {
	args := o.Called()
	return args.String(0)
}
func (o *OptsMock) HTTPS() bool {
	args := o.Called()
	return args.Bool(0)
}

func (o *OptsMock) IsTrustedIp(ip string) bool {
	args := o.Called(ip)
	return args.Bool(0)
}
