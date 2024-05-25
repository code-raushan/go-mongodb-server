package services

import (
	"github.com/code-raushan/go-mongodb-server/types"
)

type MongoServicer interface {
	Fetch(filter types.FilterOptions) []types.FetchResponse
}
type MongoService struct {
	client MongoServicer
}

func NewMongoService(client MongoServicer) *MongoService {
	return &MongoService{
		client: client,
	}
}

func (s *MongoService) FetchRecords(filter types.FilterOptions) *types.UserResponse {
	res := s.client.Fetch(filter)

	return &types.UserResponse{
		Code: 0,
		Msg: "successful",
		Records: res,
	}
}
