package events

import (
    "github.com/fireteamsupport/identity/pkg/natslib"
)

type (
    Channels struct {
        *natslib.Channels
    }

    client struct {
        natslib.Client
    }

    Client interface {
        Send() *Channels
        Recv() *Channels
        Close()
    }
)

func (c *client) Send() *Channels {
    channels := c.Channels()
    c.BindSendChan("events.invalidate_refresh_token", channels.InvalidateRefreshToken)
    return &Channels{channels}
}

func (c *client) Recv() *Channels {
    return &Channels{c.Channels()}
}

func New(url string) (Client, error) {
    nats, err := natslib.New(url)
    if err != nil {
        return nil, err
    }

    return &client{nats}, nil
}
