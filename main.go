package main

import (
	Pkgwf "transmitter/workflow"

	"github.com/robfig/cron/v3"
)

func init() {
	workflow = Pkgwf.New()
	workflow.GetConfig()
}
func task() {
	workflow.GetValue()
	workflow.Notify()
	workflow.WriteDB()
}
func report() {

	workflow.GetValue()
	workflow.NotifyReport()
}
func sched() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", task)
	c.AddFunc("0 21 * * *", report)
	c.Start()
}

var workflow Pkgwf.Workflower

func main() {
	sched()
	select {}
	/*for {
		task()
		time.Sleep(time.Second * 10)
	}*/

}
