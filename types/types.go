package types

import "time"

type FilterOptions struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MaxCount  int       `json:"maxCount"`
	MinCount  int       `json:"minCount"`
}

type FetchResponse struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type UserResponse struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Records []FetchResponse `json:"records"`
}
