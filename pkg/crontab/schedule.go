package main

import (
	//"github.com/robfig/cron"
	"github.com/jakecoffman/cron"   // add RemoveJob func
	"log"
)



func Schedule(c *cron.Cron) {
	// add params
	r := Reset{id: 5}
	c.AddJob("*/5 * * * * ?", r, "reset")  // 5 second once
	c.AddFunc("0 30 * * * *", Clear, "clear")  // Every hour on the half hour
}

func main() {
	log.Println("Starting...")

	c := cron.New()
	Schedule(c) // add tasks

	//start cron
	c.Start()

	//close cron
	defer c.Stop()

	select{} // block main process
}
