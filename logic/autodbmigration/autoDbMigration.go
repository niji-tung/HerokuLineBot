package autodbmigration

import (
	"heroku-line-bot/storage/database"
	"heroku-line-bot/storage/database/common"
)

func MigrationNotExist() error {
	tables := []*common.BaseTable{
		database.Club.Member.BaseTable,
		database.Club.Income.BaseTable,
	}
	for _, table := range tables {
		if !table.IsExist() {
			if err := table.CreateTable(); err != nil {
				return err
			}
		}
	}

	return nil
}
