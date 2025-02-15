package dro

import (
	"msg/cmd/app/server/bot/status"
)

type BotSender struct {
	Key    string             `json:"key"`
	Name   string             `json:"name"`
	Status []*BotSenderStatus `json:"status"`
}

type BotSenderStatus struct {
	ChannelType        string              `json:"channel_type"`
	ChannelName        string              `json:"channel_name"`
	ChannelEnglishName string              `json:"channel_english_name"`
	SenderStatus       status.SenderStatus `json:"sender_status"`
	SenderStatusStr    string              `json:"sender_status_str"`
	ReceiverStatus     status.SenderStatus `json:"receiver_status"`
	ReceiverStatusStr  string              `json:"receiver_status_str"`
}
