package main

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/SergeyShpak/TgAlerts/types"
)

type BotAPI struct {
	Client            *Client
	updatesBufferSize int
}

func NewBotAPI(token string) (*BotAPI, error) {
	client, err := NewClient(token)
	if err != nil {
		return nil, err
	}
	botAPI := &BotAPI{
		Client:            client,
		updatesBufferSize: 100,
	}
	return botAPI, nil
}

func (b *BotAPI) GetMe() error {
	if err := b.Client.DoRequest("GET", "getMe", nil, nil); err != nil {
		return err
	}
	return nil
}

type GetUpdatesParams struct {
	Offset         int32
	Limit          int32
	Timeout        int32
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

func (b *BotAPI) GetUpdates(params *GetUpdatesParams) ([]types.Update, error) {
	params = defaultifyGetUpdatesParams(params)
	v := url.Values{}
	if params.Offset != 0 {
		v.Add("offset", strconv.Itoa(int(params.Offset)))
	}
	if params.Limit > 0 {
		v.Add("limit", strconv.Itoa(int(params.Limit)))
	}
	if params.Timeout > 0 {
		v.Add("timeout", strconv.Itoa(int(params.Timeout)))
	}
	var updates []types.Update
	err := b.Client.DoRequest("GET", "getUpdates", v, &updates)
	if err != nil {
		return nil, err
	}
	return updates, nil
}

type UpdatesChannel <-chan types.Update

func (b *BotAPI) GetUpdatesChan(params *GetUpdatesParams) (UpdatesChannel, error) {
	ch := make(chan types.Update, b.updatesBufferSize)

	params = defaultifyGetUpdatesParams(params)
	go func() {
		for {
			updates, err := b.GetUpdates(params)
			if err != nil {
				log.Println("failed to get updates: ", err)
				time.Sleep(time.Second * 2)
				continue
			}
			for _, update := range updates {
				if update.UpdateID >= params.Offset {
					params.Offset = update.UpdateID + 1
				}
				ch <- update
			}
		}
	}()
	return ch, nil
}
