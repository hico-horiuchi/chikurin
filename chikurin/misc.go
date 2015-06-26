package chikurin

import (
	"fmt"
	"os"
	"time"

	"github.com/hico-horiuchi/ohgi/sensu"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func at(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006/01/02 15:04:05")
}

func since(event sensu.EventStruct) string {
	elapsed := time.Duration(event.Occurrences) * time.Duration(event.Check.Interval) * time.Second
	return time.Unix(event.Check.Executed, 0).Add(-1 * elapsed).Format("2006/01/02 15:04:05")
}
