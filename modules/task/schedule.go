package task

import (
	"go-ops/config"
	"go-ops/models"
	"go-ops/modules/cron"
	"go-ops/modules/log"
)

var (
	TaskMap map[string]cron.Job  // 任务集
)

type TaskCron struct {
	*cron.Cron
}

func (tc TaskCron) CreateTask(task models.Task) {
	if task.Frequency == "once" {
		runtime := task.StartTime.Time
		tc.Schedule(cron.RunAt(runtime), TaskMap[task.Name], task.Name)
	} else {
		tc.AddFunc(task.Frequency, func() {}, task.Name)
	}
}

// 任务初始化
func ScheduleInit(tc TaskCron) {
	var tasks []models.Task
	models.DB.Find(&tasks)
	for _, task := range tasks {
		tc.CreateTask(task)
	}
}

// 定时清理日志
func init() {
	c := cron.New()
	//r := Reset{id: 5}
	//c.AddJob("*/5 * * * * ?", r, "reset")
	common := config.GetConfig().Common
	// 切割日志
	c.AddFunc("0 59 23 * * *", func() {
		cutLog(common.ACCESS_LOG_PATH)
		cutLog(common.INFO_LOG_PATH)
		cutLog(common.ERROR_LOG_PATH)
		log.InitAllLogger()
	}, "cutlog")
	c.Start()
}
