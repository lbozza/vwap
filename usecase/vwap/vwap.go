package vwap

import (
	"math/big"
)

const maxDataPointSize = 200

type DataPoint struct {
	ProductID string
	Price     *big.Float
	Volume    *big.Float
}

type DataPointList struct {
	DataPoints        []DataPoint
	SumVolumeAndPrice *big.Float
	SumVolume         *big.Float
	Position          int
	MaxSize           int
}

func NewVwapCalculator() *DataPointList {
	return &DataPointList{
		DataPoints:        []DataPoint{},
		SumVolumeAndPrice: new(big.Float),
		SumVolume:         new(big.Float),
		Position:          0,
		MaxSize:           maxDataPointSize,
	}

}

func (l *DataPointList) Calculate(d DataPoint) string {

	oldestPrice, oldestVolume := new(big.Float), new(big.Float)
	if len(l.DataPoints) >= maxDataPointSize {
		oldestPrice, oldestVolume = l.getOldestPriceAndVolume(d)
	}

	newPriceAndVol := new(big.Float)
	newPriceAndVol.Mul(d.Price, d.Volume)

	if oldestPrice != nil && oldestVolume != nil {
		oldPriceAndVol := new(big.Float)
		oldPriceAndVol.Mul(oldestPrice, oldestVolume)

		l.SumVolumeAndPrice.Sub(l.SumVolumeAndPrice, oldPriceAndVol)
		l.SumVolume.Sub(l.SumVolume, oldestVolume)

		l.SumVolumeAndPrice.Add(l.SumVolumeAndPrice, newPriceAndVol)
		l.SumVolume.Add(l.SumVolume, d.Volume)

	} else {
		l.SumVolumeAndPrice.Add(l.SumVolumeAndPrice, newPriceAndVol)
		l.SumVolume.Add(l.SumVolume, d.Volume)
	}

	vwapCalculated := new(big.Float)
	vwapCalculated.Quo(l.SumVolumeAndPrice, l.SumVolume)

	// fmt.Println("PRODUCT: ", d.ProductID, "PRECO: ", d.Price, "VOL: ", d.Volume)
	// fmt.Printf("VWAP Calculated for %s is %v", d.ProductID, vwapCalculated)
	// fmt.Println("")

	l.DataPoints = append(l.DataPoints, d)

	return vwapCalculated.String()

}

func (l *DataPointList) getOldestPriceAndVolume(d DataPoint) (oldestPrice, oldestVolume *big.Float) {

	if l.MaxSize == l.Position {
		l.Position = 0
	}

	//	fmt.Println("Index is: ", l.Position)
	oldest := l.DataPoints[l.Position]
	//	fmt.Println("OLDEST IS: ", oldest)
	//	fmt.Println("NEW TRADE IS: ", d)

	if &oldest != nil {
		oldestPrice, oldestVolume = oldest.Price, oldest.Volume
	}

	l.DataPoints[l.Position] = d
	l.Position++

	return oldestPrice, oldestVolume

}
