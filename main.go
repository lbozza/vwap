package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/lbozza/vwap/entity"
	"github.com/lbozza/vwap/infra"
	usecase "github.com/lbozza/vwap/usecase"
	"github.com/lbozza/vwap/usecase/vwap"
	"github.com/pkg/errors"
)

type Handler struct {
	infra.ClientHandler
}

var pairList = []string{entity.BTCUSD, entity.ETHUSD, entity.ETHBTC}

func main() {

	ctx := context.Background()
	wg := &sync.WaitGroup{}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for _, pair := range pairList {
		err := initialize(ctx, pair, wg)

		if err != nil {
			errors.Wrap(err, "Error to Initialize VWAP Calculator for : "+pair)
		}
	}

	for {
		select {
		case <-interrupt:
			shutdown(wg)
			return

		}
	}

}

func initialize(ctx context.Context, pair string, wg *sync.WaitGroup) (err error) {
	client, err := infra.NewClient()
	handler := Handler{&client}

	wg.Add(1)

	if err != nil {
		return err
	}

	vwapCalc := vwap.NewVwapCalculator()
	tradeChannel := make(chan entity.ResponseInternal)
	service := usecase.NewService(tradeChannel, pair, *vwapCalc)

	go handler.Subscribe(ctx, []string{pair}, tradeChannel)
	go service.Execute()

	return

}

func shutdown(wg *sync.WaitGroup) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()
}
