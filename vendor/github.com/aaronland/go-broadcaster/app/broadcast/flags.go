package broadcast

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var broadcaster_uris multi.MultiCSVString

var title string

var body string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("broadcast")

	fs.Var(&broadcaster_uris, "broadcaster", "...")

	fs.StringVar(&title, "title", "", "...")
	fs.StringVar(&body, "body", "", "...")

	return fs
}
