package broadcast

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-broadcaster"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	br, err := broadcaster.NewMultiBroadcasterFromURIs(ctx, broadcaster_uris...)

	if err != nil {
		return fmt.Errorf("Failed to create broadcaster, %w", err)
	}

	br.SetLogger(ctx, logger)

	msg := &broadcaster.Message{
		Title: title,
		Body:  body,
	}

	id, err := br.BroadcastMessage(ctx, msg)

	if err != nil {
		return fmt.Errorf("Failed to broadcast message, %w", err)
	}

	fmt.Println(id.String())
	return nil
}
