package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lbozza/vwap/entity"
	"golang.org/x/net/websocket"
)

type ClientHandler interface {
	Subscribe(ctx context.Context, pairs []string, channel chan entity.ResponseInternal) error
}

type Request struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []Channel `json:"channels"`
}

type Channel struct {
	Name       string
	ProductIDs []string
}

type Response struct {
	Type      string    `json:"type"`
	Channels  []Channel `json:"channels"`
	Message   string    `json:"message,omitempty"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	ProductID string    `json:"product_id"`
}

type CliHandler struct {
	conn *websocket.Conn
}

func NewClient(address string) (CliHandler, error) {
	conn, err := websocket.Dial(address, "", "http://localhost")

	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}

	return CliHandler{
		conn: conn,
	}, nil

}

func (c *CliHandler) Subscribe(ctx context.Context, pairs []string, channel chan entity.ResponseInternal) error {

	subscription := Request{
		Type:       "subscribe",
		ProductIDs: pairs,
		Channels: []Channel{
			{Name: "matches"},
		},
	}

	payload, _ := json.Marshal(subscription)

	err := websocket.Message.Send(c.conn, payload)

	if err != nil {
		print(err)
	}

	var response entity.ResponseInternal

	websocket.JSON.Receive(c.conn, &response)

	go readClientMessage(ctx, c.conn, channel)

	return nil
}

func readClientMessage(ctx context.Context, conn *websocket.Conn, incomingMessages chan entity.ResponseInternal) {
	for {
		select {
		case <-ctx.Done():
			err := conn.Close()
			if err != nil {
				log.Printf("failed closing ws connection: %s", err)
			}
		default:
			var message Response

			err := websocket.JSON.Receive(conn, &message)
			if err != nil {
				log.Printf("failed receiving message: %s", err)

				break
			}
			incomingMessages <- entity.ResponseInternal{
				Type:      message.Type,
				Size:      message.Size,
				Price:     message.Price,
				ProductID: message.ProductID,
			}
		}
	}
}
