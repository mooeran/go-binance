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

func (s *fundingInfoServiceTestSuite) TestFundingInfo() {
}
