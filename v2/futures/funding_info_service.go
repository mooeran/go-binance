package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type FundingInfo struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int    `json:"fundingIntervalHours"`
	Disclaimer               bool   `json:"disclaimer"`
}

type FundingInfoService struct {
	c *Client
}

func (s *FundingInfoService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingInfo, err error) {
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
