package clublinebot

import (
	lineBotDomain "heroku-line-bot/service/linebot/domain"
	lineBotModel "heroku-line-bot/service/linebot/domain/model"
)

func (b *ClubLineBot) handleMessageEvent(event *lineBotModel.MessageEvent) error {
	replyToken := event.ReplyToken
	if err := b.tryLine(
		func() error {
			message := event.Message.(lineBotModel.MessageEventMessage)
			switch message.Type {
			case lineBotDomain.TEXT_MESSAGE_EVENT_MESSAGE_TYPE:

			}
			return nil
		},
		replyToken,
	); err != nil {
		return err
	}

	return nil
}
