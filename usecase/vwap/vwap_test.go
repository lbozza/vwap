package vwap

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type VwapSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller
}

func TestVwapSuite(t *testing.T) {
	suite.Run(t, new(VwapSuite))
}

func (s *VwapSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
}

func (s *VwapSuite) ShutDown() {
	s.ctrl.Finish()
}

func (s *VwapSuite) TestCalculateVwap() {
	vwapCalc := NewVwapCalculator()

	datapoints := populateTrades()

	vwapValues := []string{}

	for _, datapoint := range datapoints {
		vwapValues = append(vwapValues, vwapCalc.Calculate(datapoint))
	}

	vwapCalculatedExpected := getExpectedVwapCalculated()

	s.Equal(vwapValues, vwapCalculatedExpected)

}
func getExpectedVwapCalculated() []string {

	vwaps := []string{"35212", "35298.00382", "35299.06707", "35299.68169", "35283.84885"}

	return vwaps

}
func populateTrades() []DataPoint {

	dataPoints := []DataPoint{}

	price, _ := ParseFloat("35212")
	vol, _ := ParseFloat("0.65223")

	dp := DataPoint{
		ProductID: "BTC-USD",
		Price:     price,
		Volume:    vol,
	}

	dataPoints = append(dataPoints, dp)

	price, _ = ParseFloat("36312")
	vol, _ = ParseFloat("0.05532")

	dp = DataPoint{
		ProductID: "BTC-USD",
		Price:     price,
		Volume:    vol,
	}

	dataPoints = append(dataPoints, dp)

	price, _ = ParseFloat("37123.23")
	vol, _ = ParseFloat("0.00041241")

	dp = DataPoint{
		ProductID: "BTC-USD",
		Price:     price,
		Volume:    vol,
	}

	dataPoints = append(dataPoints, dp)

	price, _ = ParseFloat("35902.99")
	vol, _ = ParseFloat("0.00072124")

	dp = DataPoint{
		ProductID: "BTC-USD",
		Price:     price,
		Volume:    vol,
	}
	dataPoints = append(dataPoints, dp)

	price, _ = ParseFloat("35199.11")
	vol, _ = ParseFloat("0.1324124")

	dp = DataPoint{
		ProductID: "BTC-USD",
		Price:     price,
		Volume:    vol,
	}
	dataPoints = append(dataPoints, dp)

	return dataPoints

}
