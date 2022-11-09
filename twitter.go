package twitter

// https://developer.twitter.com/en/docs/labs/overview/error-codes
// https://developer.twitter.com/en/docs/basics/authentication/overview/3-legged-oauth
// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-statuses-update.html
// https://developer.twitter.com/en/docs/media/upload-media/overview
// https://developer.twitter.com/en/docs/media/upload-media/uploading-media/chunked-media-upload
// https://developer.twitter.com/en/docs/media/upload-media/api-reference/post-media-upload.html

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/aaronland/go-broadcaster"	
	"github.com/aaronland/go-broadcaster-twitter/oauth"
	"github.com/aaronland/go-image-encode"
	"github.com/aaronland/go-uid"
	"github.com/sfomuseum/runtimevar"
	"image"
	"log"
	"net/url"
	"strconv"
	"time"
)

func init() {
	ctx := context.Background()
	broadcaster.RegisterBroadcaster(ctx, "twitter", NewTwitterBroadcaster)
}

type TwitterBroadcaster struct {
	broadcaster.Broadcaster
	twitter_client *anaconda.TwitterApi
	testing        bool
	encoder        encode.Encoder
	logger         *log.Logger
}

func NewTwitterBroadcaster(ctx context.Context, uri string) (broadcaster.Broadcaster, error) {

	parsed, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	query := parsed.Query()
	creds_uri := query.Get("credentials")

	if creds_uri == "" {
		return nil, fmt.Errorf("Missing ?credentials= parameter")
	}

	rt_ctx, rt_cancel := context.WithTimeout(ctx, 5*time.Second)
	defer rt_cancel()

	str_creds, err := runtimevar.StringVar(rt_ctx, creds_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to config from credentials, %w", err)
	}

	creds, err := oauth.NewOAuth1CredentialsFromString(ctx, str_creds)

	if err != nil {
		return nil, err
	}

	tw_client := anaconda.NewTwitterApiWithCredentials(creds.AccessToken, creds.AccessSecret, creds.ConsumerKey, creds.ConsumerSecret)

	_, err = tw_client.VerifyCredentials()

	if err != nil {
		return nil, err
	}

	enc, err := encode.NewEncoder(ctx, "png://")

	if err != nil {
		return nil, err
	}

	logger := log.Default()

	br := &TwitterBroadcaster{
		twitter_client: tw_client,
		testing:        false,
		encoder:        enc,
		logger:         logger,
	}

	return br, nil
}

func (b *TwitterBroadcaster) BroadcastMessage(ctx context.Context, msg *broadcaster.Message) (uid.UID, error) {

	params := url.Values{}

	if len(msg.Images) > 0 {

		for _, im := range msg.Images {

			media_id, err := b.uploadImage(im)

			if err != nil {
				return nil, err
			}

			str_media_id := strconv.FormatInt(media_id, 10)

			params.Add("media_ids", str_media_id)
		}
	}

	status := msg.Body

	if b.testing {
		status = fmt.Sprintf("this is a test and there may be more / please disregard and apologies for the distraction / meanwhile: %s", status)
	}

	tw, err := b.twitter_client.PostTweet(status, params)

	if err != nil {
		return nil, err
	}

	b.logger.Printf("twitter post %d (media id: %s) ", tw.Id, params.Get("media_id"))

	return uid.NewInt64UID(ctx, tw.Id)
}

func (b *TwitterBroadcaster) SetLogger(ctx context.Context, logger *log.Logger) error {
	b.logger = logger
	return nil
}

func (b *TwitterBroadcaster) uploadImage(im image.Image) (int64, error) {

	ctx := context.Background()

	// but what if GIF...

	out := new(bytes.Buffer)

	err := b.encoder.Encode(ctx, im, out)

	if err != nil {
		return -1, err
	}

	return b.uploadMedia(out.Bytes())
}

func (b *TwitterBroadcaster) uploadMedia(body []byte) (int64, error) {

	b64_body := base64.StdEncoding.EncodeToString(body)

	rsp, err := b.twitter_client.UploadMedia(b64_body)

	if err != nil {
		return -1, err
	}

	return rsp.MediaID, nil
}
