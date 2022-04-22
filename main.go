package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	//buffer := bytes.NewBuffer([]byte{})
	//process := exec.Command("git", "s")
	//process.Stdout = buffer
	//err2 := process.Run()
	//if err2 != nil {
	//	log.Fatalln(err2)
	//}
	//if process.ProcessState.Success() {
	//	println(buffer.String())
	//}
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-01 10:09:10", time.Local)
	if err != nil {
		return
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-31 16:32:13", time.Local)
	if err != nil {
		return
	}
	now := time.Now()

	startDay := int(now.Sub(startTime) / (time.Hour * 24))
	endDay := int(now.Sub(endTime) / (time.Hour * 24))

	//career := endTime.Sub(startTime)
	//fmt.Printf("hours: %s\n", career.String())
	//daysBetween := career / (time.Hour * 24)
	//fmt.Printf("daysBetween: %d\n", startDay-endDay)

	for i := startDay; i > endDay; i-- {
		//println(i)
		err := workOneDay(i)
		if err != nil {
			log.Fatalln(err)
		}
	}
	println("completed")
}

func workOneDay(daysBefore int) (err error) {
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
		_, err := botfile.WriteString(fmt.Sprintf("\n%d days ago", daysBefore))
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
			fmt.Sprintf("%d days ago (%d)", daysBefore, index),
			fmt.Sprintf("--date=format:relative:%d.days.ago", daysBefore))
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

func commitOne(daysBefore int) {
	stdOut := bytes.NewBuffer([]byte{})
	dateString := fmt.Sprintf("--date=format:relative:%d.days.ago", daysBefore)
	p1 := exec.Cmd{
		Path:   "/usr/bin/git",
		Args:   []string{"/usr/bin/git", "commit", "-m", "hello", dateString},
		Stdout: stdOut,
	}
	err := p1.Run()
	if err != nil {
		log.Fatalln(err)
	}
	if p1.ProcessState.Success() {
		println(stdOut.String())

	}
}

// modify file
// exec add
// exec commit
