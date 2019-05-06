package crontab

import (
	"go-ops/config"
	"go-ops/models"
	"go-ops/models/tasks"
	"go-ops/modules/cron"
	"go-ops/modules/log"
)

var (
	JobMap map[string]func(tasks.Parameter) // 任务集
	Crontab TaskCron
)

type TaskCron struct {
	*cron.Cron
}

func NewCron() TaskCron {
	c := cron.New()
	return TaskCron{c}
}

func (tc TaskCron) CreateTask(task tasks.Task) {
	if task.Frequency == "once" {
		runtime := task.StartTime.Time
		tc.Schedule(cron.RunAt(runtime), task, task.Name)
	} else {
		tc.AddFunc(task.Frequency, func() {}, task.Name)
	}
}

// cron init
func ScheduleInit(tc TaskCron) {
	var tasks []tasks.Task
	models.DB.Find(&tasks)
	for _, task := range tasks {
		tc.CreateTask(task)
	}
}


// 定时清理日志
func init() {
	Crontab = NewCron()
	//r := Reset{id: 5}
	//c.AddJob("*/5 * * * * ?", r, "reset")
	common := config.GetConfig().Common
	// 切割日志
	Crontab.AddFunc("0 59 23 * * *", func() {
		cutLog(common.ACCESS_LOG_PATH)
		cutLog(common.INFO_LOG_PATH)
		cutLog(common.ERROR_LOG_PATH)
		log.InitAllLogger()
	}, "cutlog")
	Crontab.Start()
}
