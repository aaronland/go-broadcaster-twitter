package broadcaster

import (
	"context"
	"github.com/aaronland/go-uid"
	"log"
	"strconv"
	"time"
)

func init() {
	ctx := context.Background()
	RegisterBroadcaster(ctx, "log", NewLogBroadcaster)
}

type LogBroadcaster struct {
	Broadcaster
	logger *log.Logger
}

func NewLogBroadcaster(ctx context.Context, uri string) (Broadcaster, error) {
	logger := log.Default()
	b := LogBroadcaster{
		logger: logger,
	}
	return &b, nil
}

func (b *LogBroadcaster) BroadcastMessage(ctx context.Context, msg *Message) (uid.UID, error) {
	b.logger.Println(msg.Body)

	now := time.Now()
	ts := now.Unix()

	// pending uid.NewInt64UID
	str_ts := strconv.FormatInt(ts, 10)

	return uid.NewStringUID(ctx, str_ts)
}

func (b *LogBroadcaster) SetLogger(ctx context.Context, logger *log.Logger) error {
	b.logger = logger
	return nil
}
