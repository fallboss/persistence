package persistence

import "time"

type PostgresProp struct {
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassword       string
	DbName           string
	DbSslMode        string
	DbSelectLimit    int
	DbMaxCnx         int
	DbWaitAfterQuery time.Duration
}
