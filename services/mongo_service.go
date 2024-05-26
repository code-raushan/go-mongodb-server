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

func (s *MongoService) FetchRecords(filter *types.FilterOptions) *types.UserResponse {
	res, err := s.repo.Fetch(filter)

	if err != nil {
		return &types.UserResponse{
			Code: 1,
			Msg: err.Error(),
			Records: make([]types.FetchResponse, 0),
		}
	}

	return &types.UserResponse{
		Code: 0,
		Msg: "Success",
		Records: res,
	}
}
