package client // import "github.com/docker/docker/client"

import (
	"context"
	"io"
	"log"
	"net/url"

	"github.com/docker/docker/api/types"
)

func (cli *Client) ImageMakeCache(ctx context.Context, input io.Reader) (types.ImageMakeCacheResponse, error) {
	log.Printf("Hitting ImageMakeCache method")
	v := url.Values{}
	headers := map[string][]string{"Content-Type": {"application/json"}}
	resp, err := cli.postRaw(ctx, "/images/makecache", v, input, headers)
	if err != nil {
		return types.ImageMakeCacheResponse{}, err
	}
	return types.ImageMakeCacheResponse{
		Body: resp.body,
		JSON: resp.header.Get("Content-Type") == "application/json",
	}, nil
}
