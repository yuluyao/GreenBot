package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-01 10:09:10", time.Local)
	if err != nil {
		return
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-31 16:32:13", time.Local)
	if err != nil {
		return
	}
	now := time.Now()

	daysAgoStart := int(now.Sub(startTime) / (time.Hour * 24))
	daysAgoEnd := int(now.Sub(endTime) / (time.Hour * 24))

	for i := daysAgoStart; i > daysAgoEnd; i-- {
		err := workOneDay(i)
		if err != nil {
			log.Fatalln(err)
		}
	}
	println("completed")
}

func workOneDay(daysAgo int) (err error) {
	var botfile *os.File
	info, err := os.Stat("botfile")
	if err != nil {
		botfile, err = os.Create("botfile")
		if err != nil {
			return err
		}
	} else {
		botfile, err = os.OpenFile(info.Name(), os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	}

	commitTimes := rand.Intn(15) + 1

	oneCommit := func(index int) error {
		// modify
		_, err := botfile.WriteString(fmt.Sprintf("\n%d days ago", daysAgo))
		if err != nil {
			return err
		}
		//println(effected)

		// git add
		prcAdd := exec.Command("git", "add", ".")
		err = prcAdd.Run()
		if err != nil {
			return err
		}
		// git commit
		prcCommit := exec.Command("git", "commit", "-m",
			fmt.Sprintf("%d days ago (%d)", daysAgo, index),
			fmt.Sprintf("--date=format:relative:%d.days.ago", daysAgo))
		err = prcCommit.Run()
		if err != nil {
			return err
		}
		return nil
	}

	for i := 0; i < commitTimes; i++ {
		err = oneCommit(i + 1)
		if err != nil {
			return err
		}
	}

	return nil
}
