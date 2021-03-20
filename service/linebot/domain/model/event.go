package model

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
