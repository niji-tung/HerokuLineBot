package clublinebot

import (
	"heroku-line-bot/service/googlescript"
	"heroku-line-bot/service/linebot"
	lineBotReqs "heroku-line-bot/service/linebot/domain/model/reqs"
)

type ClubLineBot struct {
	lineAdminID,
	lineRoomID string
	*linebot.LineBot
	*googlescript.GoogleScript
}

func (b *ClubLineBot) Handle(json string) error {
	return b.LinePage(json)
}

func (b *ClubLineBot) replyErr(err error, replyToken string) error {
	if _, err := b.ReplyMessage(
		&lineBotReqs.ReplyMessage{
			ReplyToken: replyToken,
			Messages: []interface{}{
				linebot.GetTextMessage("統發生錯誤，已通知管理員"),
			},
		}); err != nil {
		return err
	}

	if _, err := b.PushMessage(
		&lineBotReqs.PushMessage{
			To: b.lineAdminID,
			Messages: []interface{}{
				linebot.GetTextMessage(err.Error()),
			},
		}); err != nil {
		return err
	}

	return nil
}

func (b *ClubLineBot) tryReply(replyToken string, messages []interface{}) error {
	if _, err := b.ReplyMessage(
		&lineBotReqs.ReplyMessage{
			ReplyToken: replyToken,
			Messages:   messages,
		}); err != nil {
		return b.replyErr(err, replyToken)
	}
	return nil
}
