package services

import (
	"github.com/code-raushan/go-mongodb-server/repositories"
	"github.com/code-raushan/go-mongodb-server/types"
)

type MongoServicer interface {
	Fetch(filter types.FilterOptions) []types.FetchResponse
}
type MongoService struct {
	repo *repositories.MongoRepo
}

func NewMongoService(repo *repositories.MongoRepo) *MongoService {
	return &MongoService{
		repo: repo,
	}
}

func (s *MongoService) FetchRecords(filter types.FilterOptions) *types.UserResponse {
	res := s.repo.Fetch(filter)

	return &types.UserResponse{
		Code: 0,
		Msg: "successful",
		Records: res,
	}
}
