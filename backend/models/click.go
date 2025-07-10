package models

import "time"

type ClickEvent struct {
	ShortKey  string    `bson:"short_key"`
	Timestamp time.Time `bson:"timestamp"`
	IP        string    `bson:"ip"`
	UserAgent string    `bson:"user_agent"`
}
