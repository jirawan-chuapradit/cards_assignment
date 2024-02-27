package config

import "time"

var Location *time.Location

func SetUp() {
	var loc, _ = time.LoadLocation("Asia/Bangkok")
	Location = loc
}
