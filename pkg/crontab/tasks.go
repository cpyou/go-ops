package main

import "fmt"

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