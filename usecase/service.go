package usecase

import (
	"fmt"
	"log"

	"github.com/lbozza/vwap/entity"
	"github.com/lbozza/vwap/usecase/vwap"
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
			log.Printf("error parsing price to float")
		}

		volume, err := vwap.ParseFloat(data.Size)

		if err != nil {
			log.Printf("error parsing volume to float")
		}

		vwapValue := s.VwapCalculator.Calculate(vwap.DataPoint{
			Price:     price,
			Volume:    volume,
			ProductID: data.ProductID,
		})

		fmt.Println("PRODUCT: ", data.ProductID, "PRECO: ", data.Price, "VOL: ", volume)
		fmt.Printf("VWAP Calculated for %s is %v", data.ProductID, vwapValue)
		fmt.Println("")

		// fmt.Println(list.DataPoints)

		//fmt.Println(data)
	}

}
