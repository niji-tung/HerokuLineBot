package main

import (
	"embed"
	"heroku-line-bot/entry"
)

//go:embed config/*
var f embed.FS

// @title Line Bot
// @version 1.0
func main() {
	if err := entry.Run(f); err != nil {
		panic(err)
	}
}
