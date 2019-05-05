package crontab

import (
	"fmt"
	"go-ops/modules/log"
	"os"
	"os/exec"
	"time"
)

func Clear() {

}

/*
# Job is an interface for submitted cron jobs.
type Job interface {
	Run()
}
*/
type Reset struct {
	id int
}

func (r Reset) Run() {
	fmt.Println(r.id)
}

type Shell interface {
	Run() string
}


// 切割日志
func cutLog(path string) {
	date := time.Now().Format("20060102")
	err := os.Rename(path, path+"."+date+".log")
	if err != nil {
		log.Println(path + " rename Error!")
	} else {
		log.Println(path + " rename OK!")
	}
}

func RunShell(command string) string {
	cmd := exec.Command(command)
	if output, err := cmd.Output(); err != nil {
		log.Error(err)
		log.Println(output)
		panic(err)
	} else {
		return string(output)
	}
}
