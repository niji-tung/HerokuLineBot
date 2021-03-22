package model

import "heroku-line-bot/service/linebot/domain"

type EventBase struct {
	Type
	ReplyToken string `json:"replyToken"`
}

type MemberJoinEventJoined struct {
	Members []*Source `json:"members"`
}

type MemberJoinEvent struct {
	EventBase
	Joined MemberJoinEventJoined `json:"joined"`
	Source Source                `json:"source"`
}

type MessageEvent struct {
	EventBase
	Message interface{} `json:"message"`
	Source  Source      `json:"source"`
}

type MessageEventMessage struct {
	Type domain.MessageEventMessageType `json:"type"`
}

type MessageEventText struct {
	MessageEventMessage
	Text string `json:"text"`
}
