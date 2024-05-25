package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/code-raushan/go-mongodb-server/types"
)

func ExtractFilters(query url.Values) (filters *types.FilterOptions, err error) {
	filters = &types.FilterOptions{}

	fmt.Printf("%v", query)

	if minCountStr := query["minCount"]; minCountStr[0] != "" {
		minCount, err := strconv.Atoi(minCountStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid minCount param: %v", err)
		}
		filters.MinCount = minCount
	}

	if maxCountStr := query["maxCount"]; maxCountStr[0] != "" {
		maxCount, err := strconv.Atoi(maxCountStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid maxCount param: %v", err)
		}
		filters.MaxCount = maxCount
	}

	if startDateStr := query["startDate"]; startDateStr[0] != "" {
		startDate, err := time.Parse(time.DateOnly, startDateStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start date param: %v", err)
		}
		filters.StartDate = startDate
	}

	if endDateStr := query["endDate"]; endDateStr[0] != "" {
		endDate, err := time.Parse(time.DateOnly, endDateStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid end date param: %v", err)
		}
		filters.EndDate = endDate
	}

	fmt.Printf("%v", filters)

	return filters, nil
}