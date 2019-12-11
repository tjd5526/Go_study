package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504"}
	for _, logEntry := range logs {
		r := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r.MatchString(logEntry) {
			match := r.FindStringSubmatch(logEntry)
			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("Not a valid date time format!")
			}
		}

	}
}
