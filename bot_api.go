package main

import (
	"github.com/SergeyShpak/TgAlerts/types"
)

type BotAPI struct {
	Client *Client
}

func NewBotAPI(token string) (*BotAPI, error) {
	client, err := NewClient(token)
	if err != nil {
		return nil, err
	}
	botAPI := &BotAPI{
		Client: client,
	}
	return botAPI, nil
}

func (b *BotAPI) GetMe() error {
	if err := b.Client.DoRequest("GET", "getMe"); err != nil {
		return err
	}
	return nil
}

type GetUpdatesParams struct {
	Offset         int
	Limit          int
	Timeout        int
	AllowedUpdates []string
}

func defaultifyGetUpdatesParams(params *GetUpdatesParams) *GetUpdatesParams {
	const (
		offset  = 0
		limit   = 100
		timeout = 0
	)
	if params == nil {
		return &GetUpdatesParams{
			Offset:  offset,
			Limit:   limit,
			Timeout: timeout,
		}
	}
	return params
}

func (b *BotAPI) GetUpdates(params *GetUpdatesParams) error {
	params = defaultifyGetUpdatesParams(params)
	err := b.Client.DoRequest("GET", "getUpdates")
	if err != nil {
		return err
	}
	return nil
}

type UpdatesChannel <-chan types.Update

func (b *BotAPI) GetUpdatesChan(params *GetUpdatesParams) UpdatesChannel {
	return nil
}
