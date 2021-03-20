package linebot

import (
	"encoding/json"
	"heroku-line-bot/service/linebot/domain"
	"heroku-line-bot/service/linebot/domain/model"

	"github.com/tidwall/gjson"
)

type EventJson struct {
	json string
}

func (e *EventJson) Raw() string {
	return e.json
}

func (e *EventJson) Parse() (eventType domain.EventType, result interface{}) {
	eventTypeJs := gjson.Get(e.json, "type")
	eventType = domain.EventType(eventTypeJs.String())
	switch eventType {
	case domain.MEMBER_JOINED_EVENT_TYPE:
		result = &model.MemberJoinEvent{}
		if err := json.Unmarshal([]byte(e.json), result); err != nil {
			return
		}
	}
	return
}
