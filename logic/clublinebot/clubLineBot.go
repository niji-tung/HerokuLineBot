package clublinebot

import (
	"fmt"
	"heroku-line-bot/service/googlescript"
	"heroku-line-bot/service/linebot"
	lineBotModel "heroku-line-bot/service/linebot/domain/model"
	lineBotReqs "heroku-line-bot/service/linebot/domain/model/reqs"
	"strings"

	lineBotDomain "heroku-line-bot/service/linebot/domain"

	"github.com/tidwall/gjson"
)

type ClubLineBot struct {
	lineAdminID,
	lineRoomID string
	*linebot.LineBot
	*googlescript.GoogleScript
}

func (b *ClubLineBot) Handle(json string) error {
	go b.LinePage(json)

	eventsJs := gjson.Get(json, "events")
	for _, eventJs := range eventsJs.Array() {
		event := linebot.NewEventJson(eventJs.Raw)
		if err := b.handleEvent(event); err != nil {
			return err
		}
	}

	return nil
}

func (b *ClubLineBot) handleEvent(eventJson *linebot.EventJson) error {
	eventType, rawEvent := eventJson.Parse()
	switch eventType {
	case lineBotDomain.MEMBER_JOINED_EVENT_TYPE:
		event := rawEvent.(*lineBotModel.MemberJoinEvent)
		replyToken := event.ReplyToken
		if err := b.tryLine(
			func() error {
				userInfoMsgs := []string{}
				for _, source := range event.Joined.Members {
					userID := source.UserID
					userInfo, err := b.GetUserProfile(userID)
					if err != nil {
						return err
					}
					msg := fmt.Sprintf("%s : %s", userInfo.DisplayName, userID)
					userInfoMsgs = append(userInfoMsgs, msg)
				}
				userInfoMsg := strings.Join(userInfoMsgs, "\n")
				groupID := event.Source.GroupID
				pushMessages := []interface{}{
					linebot.GetTextMessage("member join group : " + groupID),
					linebot.GetTextMessage(userInfoMsg),
				}
				pushReqs := &lineBotReqs.PushMessage{
					To:       b.lineAdminID,
					Messages: pushMessages,
				}
				if _, err := b.PushMessage(pushReqs); err != nil {
					return err
				}

				replyMessges := []interface{}{
					linebot.GetTextMessage("歡迎加入，跟我加入好友可以獲取更多社團的資訊喔!，點擊連結加入好友 https://line.me/R/ti/p/%4001"),
				}
				replyReqs := &lineBotReqs.ReplyMessage{
					ReplyToken: replyToken,
					Messages:   replyMessges,
				}
				if _, err := b.ReplyMessage(replyReqs); err != nil {
					return err
				}

				return nil
			},
			replyToken,
		); err != nil {
			return err
		}
	}

	return nil
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
	if err := b.tryLine(
		func() error {
			if _, err := b.ReplyMessage(
				&lineBotReqs.ReplyMessage{
					ReplyToken: replyToken,
					Messages:   messages,
				}); err != nil {
				return err
			}

			return nil
		},
		replyToken,
	); err != nil {
		return err
	}
	return nil
}

func (b *ClubLineBot) tryLine(tryF func() error, replyToken string) error {
	if err := tryF(); err != nil {
		return b.replyErr(err, replyToken)
	}
	return nil
}
