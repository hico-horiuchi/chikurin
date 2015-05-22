package chikurin

import (
	"fmt"
	"os"
	"time"
)

func utoa(timestamp int64) string {
	format := "2006/01/02 15:04:05"
	return time.Unix(timestamp, 0).Format(format)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
