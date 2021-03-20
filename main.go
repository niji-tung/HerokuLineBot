package main

import (
	"embed"
	"heroku-line-bot/entry"
)

//go:embed config/*
var f embed.FS

func main() {
	entry.Run(f)
}
