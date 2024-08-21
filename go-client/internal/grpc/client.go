package grpc

import "errors"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Serve() error {
	return errors.New("Unimplemented")
}
