package futures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fundingInfoServiceTestSuite struct {
	baseTestSuite
}

func TestFundingInfoService(t *testing.T) {
	suite.Run(t, new(fundingInfoServiceTestSuite))
}

func (s *fundingInfoServiceTestSuite) assertFundingInfoEqual(e, a *FundingInfo) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AdjustedFundingRateCap, a.AdjustedFundingRateCap, "AdjustedFundingRateCap")
	r.Equal(e.AdjustedFundingRateFloor, a.AdjustedFundingRateFloor, "AdjustedFundingRateFloor")
	r.Equal(e.FundingIntervalHours, a.FundingIntervalHours, "FundingIntervalHours")
	r.Equal(e.Disclaimer, e.Disclaimer, "Disclaimer")
}

func (s *fundingInfoServiceTestSuite) TestFundingInfo() {
	data := []byte(`[
    {
        "symbol": "BLZUSDT",
        "adjustedFundingRateCap": "0.02500000",
        "adjustedFundingRateFloor": "-0.02500000",
        "fundingIntervalHours": 8,
        "disclaimer": false
    }
]
`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{})
		s.assertRequestEqual(e, r)
	})

	fundingInfos, err := s.client.NewFundingInfoService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(fundingInfos, 1)
	e := &FundingInfo{
		Symbol:                   "BLZUSDT",
		AdjustedFundingRateCap:   "0.02500000",
		AdjustedFundingRateFloor: "-0.02500000",
		FundingIntervalHours:     8,
		Disclaimer:               false,
	}
	s.assertFundingInfoEqual(e, fundingInfos[0])
}
