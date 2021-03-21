package resp

import "time"

type ID struct {
	ID int
}

type ID64 struct {
	ID int64
}

type UID struct {
	ID uint
}

type Name struct {
	Name string
}

type ID64Name struct {
	ID   int64
	Name string
}

type ID64NameTime struct {
	ID64Name
	Time *time.Time
}

type Time struct {
	Time time.Time
}

type ID64Istest struct {
	ID     int64
	IsTest string
}
