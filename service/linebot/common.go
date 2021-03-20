package linebot

import "heroku-line-bot/service/linebot/domain/model"

func New(channelAccessToken string) *LineBot {
	return &LineBot{
		channelAccessToken: channelAccessToken,
	}
}

func NewEventJson(json string) *EventJson {
	return &EventJson{
		json: json,
	}
}

func GetTextMessage(text string) *model.TextMessage {
	return &model.TextMessage{
		Type: model.Type{
			Type: "text",
		},
		Text: text,
	}
}
