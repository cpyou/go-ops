package models

import (
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
	"go-ops/modules/cron"
	"log"
	"time"
)

type Status string

const (
	Waiting Status = "waiting"
	Doing   Status = "doing"
	Done    Status = "done"
	Fail    Status = "fail"
)

func (s *Status) Scan(value interface{}) error {
	*s = Status(value.([]byte))
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}

type Task struct {
	Model
	UserId      uint           `json:"user_id"`
	StartTime   mysql.NullTime `json:"start_time"`  // once task runtime
	EndTime     mysql.NullTime `json:"end_time"`
	DeadLine    mysql.NullTime `json:"dead_line"`
	Status      string         `json:"status" sql:"type:ENUM('waiting', 'doing', 'done', 'fail')"`
	JobName     string         `json:"job_name";gorm:"unique_index"` // function to call
	Name        string         `json:"name" gorm:"unique_index"`     // unique
	Frequency   string         `json:"frequency"`                    // cron config or once
	Active      bool           `json:"active" gorm:"default:true"`
	NextRuntime time.Time      `json:"next_runtime" gorm:"-"`
}

func (task *Task) Create() (insertId uint) {
	err := DB.Create(&task).Error
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return task.ID
}

func GetTaskByName(task *Task, name string) {
	DB.Where("name = ?", name).Find(&task, name)
}

func GetTaskById(task *Task, id uint) {
	DB.Find(&task, id)
}

func (task *Task) Stop() error {
	return nil
}

func (task *Task) GetStatus() string {
	return task.Status
}

func (task *Task) GetDetail() *Task {
	return task
}

// run task
func (task *Task) Run() error {
	fc := cron.TaskMap[task.Function]
	err := fc()
	return err
}

// get all active tasks
func GetActiveList(page_num, page_size int) []Task {
	var tasks []Task
	DB.Where("active = ?", true).Find(&tasks).Limit(page_size).Offset(page_num*page_size - page_size)
	return tasks
}

//
