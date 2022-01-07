package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lbozza/vwap/entity"
	"golang.org/x/net/websocket"
)

const address string = "wss://ws-feed.exchange.coinbase.com"

type request struct {
	Type       string    `json:"type"`
	ProductIDs []string  `json:"product_ids"`
	Channels   []channel `json:"channels"`
}

type channel struct {
	Name       string
	ProductIDs []string
}

type response struct {
	Type      string    `json:"type"`
	Channels  []channel `json:"channels"`
	Message   string    `json:"message,omitempty"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	ProductID string    `json:"product_id"`
}

type Client struct {
	conn *websocket.Conn
}

func NewClient() (Client, error) {
	conn, err := websocket.Dial(address, "", "http://localhost")

	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		return Client{}, err
	}

	return Client{
		conn: conn,
	}, nil

}

func (c *Client) Subscribe(ctx context.Context, pairs []string, tradeChannel chan entity.ResponseInternal, fatalErrors chan error) {

	subscription := request{
		Type:       "subscribe",
		ProductIDs: pairs,
		Channels: []channel{
			{Name: "matches"},
		},
	}

	payload, _ := json.Marshal(subscription)

	err := websocket.Message.Send(c.conn, payload)

	if err != nil {
		fatalErrors <- err
	}

	var response entity.ResponseInternal

	websocket.JSON.Receive(c.conn, &response)

	go readClientMessage(ctx, c.conn, tradeChannel)

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
			var message response

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
