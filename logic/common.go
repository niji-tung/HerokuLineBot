package logic

import (
	"heroku-line-bot/bootstrap"
	"heroku-line-bot/logic/clublinebot"
)

func Init(cfg *bootstrap.Config) {
	clublinebot.Init(cfg)
}
