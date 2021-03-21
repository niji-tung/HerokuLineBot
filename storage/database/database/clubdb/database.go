package clubdb

import (
	"heroku-line-bot/storage/database/common"
	"heroku-line-bot/storage/database/database/clubdb/table/income"
	"heroku-line-bot/storage/database/database/clubdb/table/member"
)

type Database struct {
	common.BaseDatabase
	Member member.Member
	Income income.Income
}
