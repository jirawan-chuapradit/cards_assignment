package config

import "time"

var Location *time.Location
var AccessSecretKey = "test-1234"
var RefreshSecretKey = "refresh-test-1234"

func SetUp() {
	var loc, _ = time.LoadLocation("Asia/Bangkok")
	Location = loc
}
