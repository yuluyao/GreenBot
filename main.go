package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	err := createOne()
	if err != nil {
		log.Fatalln(err)
	}
	println("success")
}

// 修改一次，并add
func createOne() (err error) {
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

	// modify
	effected, err := botfile.WriteString("hello\n")
	if err != nil {
		return err
	}
	println(effected)

	// git add
	prcAdd := exec.Command("git", "add", ".")
	err = prcAdd.Run()
	if err != nil {
		return err
	}
	// git commit
	dateString := fmt.Sprintf("--date=format:relative:%d.days.ago", 4)
	prcCommit := exec.Command("git", "commit", "-m", "hello", dateString)
	err = prcCommit.Run()
	if err != nil {
		return err
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
