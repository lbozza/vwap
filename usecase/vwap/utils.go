package vwap

import "math/big"

func ParseFloat(s string) (*big.Float, error) {
	float := new(big.Float)
	float, _, err := float.Parse(s, 10)

	if err != nil {
		return nil, err
	}
	return float, nil
}
