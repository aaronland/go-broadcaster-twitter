package broadcaster

import (
	"image"
)

type Message struct {
	Title  string
	Body   string
	Images []image.Image
}
