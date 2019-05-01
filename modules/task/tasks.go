package task

import (
	"fmt"
	"os"
	"time"
	"go-ops/modules/log"
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
	id  int
}

func (r Reset) Run() {
	fmt.Println(r.id)
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