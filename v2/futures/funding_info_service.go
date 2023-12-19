package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type FundingInfo struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor int64  `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int64  `json:"fundingIntervalHours"`
	Disclaimer               bool   `json:"disclaimer"`
}

type FundingInfoService struct {
	c *Client
}

func (s *FundingRateService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/fundingInfo",
		secType:  secTypeNone,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*FundingInfo{}, err
	}
	res = make([]*FundingInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*FundingInfo{}, err
	}
	return res, nil
}
