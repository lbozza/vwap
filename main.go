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
)

const address string = "wss://ws-feed.exchange.coinbase.com"

type Handler struct {
	infra.ClientHandler
}

func main() {
	print("Hello WORLD")

	ctx := context.Background()
	wg := &sync.WaitGroup{}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	err := initialize(ctx, address, "BTC-USD", wg)

	if err != nil {
		print(err)
	}

	//err = initialize(ctx, address, "ETH-USD", wg)

	for {
		select {
		case <-interrupt:
			shutdown(wg)
			return

		}
	}

}

func initialize(ctx context.Context, address string, pair string, wg *sync.WaitGroup) (err error) {
	client, err := infra.NewClient(address)
	handler := Handler{&client}

	wg.Add(1)

	if err != nil {
		print(err)
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