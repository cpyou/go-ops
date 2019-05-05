package tasks

import (
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
	"go-ops/models"
	"go-ops/modules/crontab"
	"go-ops/pkg/json"
	"golang.org/x/perf/storage/db"
	"log"
	"time"
)

// Enum type
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

type Parameter map[string]string

func (t Parameter) Value() (driver.Value, error) {
	p, e := json.Marshal(t)
	return p, e
}

func (t *Parameter) Scan(src interface{}) error {
	if e := json.Unmarshal(src.([]byte), &t); e != nil {
		return e
	}
	return nil
}

func (t Parameter) IsNull() bool {
	return len(t) == 0
}

func NewParameter(m map[string]string) Parameter {
	return Parameter(m)
}

type Task struct {
	models.BaseModel
	UserId      uint           `json:"user_id"` // user of create task
	StartTime   mysql.NullTime `json:"start_time"` // once task runtime
	EndTime     mysql.NullTime `json:"end_time"`
	DeadLine    mysql.NullTime `json:"dead_line"`
	Status      Status         `json:"status" sql:"type:ENUM('waiting', 'doing', 'done', 'fail')"`
	Name        string         `json:"name"`     // unique by program
	Frequency   string         `json:"frequency"`                    // cron config or once
	Active      bool           `json:"active" gorm:"default:true"`
	NextRuntime time.Time      `json:"next_runtime" gorm:"-"`
	Parameter   Parameter 	   `json:"parameter" sql:"type:json"`
	TaskLogs    []TaskLog
	JobId   uint
}

// 作业
type Job struct {
	models.BaseModel
	Tasks []Task
	JobName     string         `json:"job_name";gorm:"unique_index"` // function to call
}



type TaskOperation struct {

}

// run task
func (task Task) Run() {
	job := task.GetJob()
	fc := crontab.JobMap[job.JobName]
	task.AuditLog() // action audit
	fc(task.Parameter)
}

// audit log
func (task *Task) AuditLog() {
	tasklog := TaskLog{
		UserId: task.UserId,
		TaskId: task.ID,
	}
	models.DB.Create(&tasklog)
}

func (task *Task) Create() (insertId uint) {
	err := models.DB.Create(&task).Error
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return task.ID
}

func (task *Task) GetJob() (job Job) {
	models.DB.First(&job, task.JobId)
}


func GetTaskByName(task *Task, name string) {
	models.DB.Where("name = ?", name).Find(&task, name)
}

func GetTaskById(task *Task, id uint) {
	models.DB.Find(&task, id)
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

// get all active tasks
func GetActiveList(page_num, page_size int) []Task {
	var tasks []Task
	models.DB.Where("active = ?", true).Find(&tasks).Limit(page_size).Offset(page_num*page_size - page_size)
	return tasks
}

//
