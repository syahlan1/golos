package utils

import "time"

func GetDateNow() string {
	return time.Now().Format("2006-01-02")
}

func GetDateNowFormat() string {
	return time.Now().Format("02012006")
}

func GetDateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}