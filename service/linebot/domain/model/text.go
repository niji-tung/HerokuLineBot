package model

import "heroku-line-bot/service/linebot/domain"

type TextMessage struct {
	Type
	Text   string                   `json:"text"`
	Weight domain.TextMessageWeight `json:"weight"`
	Size   domain.TextMessageSize   `json:"size"`
}
