package natslib

import (
    "github.com/nats-io/nats.go"
)

type (
    Channels struct {
        InvalidateRefreshToken chan *InvalidateRefreshToken
    }

    client struct {
        *nats.EncodedConn
    }

    Client interface {
        Channels() *Channels
        BindSendChan(string, interface{}) error
        BindRecvChan(string, interface{}) (*nats.Subscription, error)
        Close()
    }
)

func (c *client) Channels() *Channels {
    return &Channels{
        InvalidateRefreshToken: make(chan *InvalidateRefreshToken),
    }
}

func New(url string) (Client, error) {
    nc, err := nats.Connect(url)
    if err != nil {
        return nil, err
    }

    ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
    if err != nil {
        return nil, err
    }

    return &client{ec}, nil
}
