package usecase

import (
	"fmt"

	"github.com/lbozza/vwap/entity"
	"github.com/lbozza/vwap/usecase/vwap"
	"github.com/pkg/errors"
)

type Service struct {
	responseChannel chan entity.ResponseInternal
	ProductID       string
	VwapCalculator  vwap.DataPointList
}

func NewService(responseChannel chan entity.ResponseInternal, productID string, vwapCalculator vwap.DataPointList) *Service {
	return &Service{
		responseChannel: responseChannel,
		ProductID:       productID,
		VwapCalculator:  vwapCalculator,
	}
}

func (s *Service) Execute() {

	for data := range s.responseChannel {

		price, err := vwap.ParseFloat(data.Price)
		if err != nil {
			errors.Wrap(err, "error parsing price to float")
		}

		volume, err := vwap.ParseFloat(data.Size)

		if err != nil {
			errors.Wrap(err, "error parsing volume to float")
		}

		vwapValue := s.VwapCalculator.Calculate(vwap.DataPoint{
			Price:     price,
			Volume:    volume,
			ProductID: data.ProductID,
		})

		fmt.Printf("VWAP Calculated for %s is %v", data.ProductID, vwapValue)
		fmt.Println("")

	}

}
