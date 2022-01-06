package infra

import (
	"context"

	"github.com/lbozza/vwap/entity"
)

type ClientHandler interface {
	Subscribe(ctx context.Context, pairs []string, channel chan entity.ResponseInternal) error
}
