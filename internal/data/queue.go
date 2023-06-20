package data

import (
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/night-sword/kratos-layout/internal/cnst"
	"github.com/night-sword/kratos-layout/internal/conf"
)

func NewQueue(data *Data, jobConfig *conf.Job) *Queue {
	return &Queue{
		data:      data,
		jobConfig: jobConfig,
	}
}

type Queue struct {
	data      *Data
	jobConfig *conf.Job
}

func (inst *Queue) Enqueue(job cnst.Job, params any, options ...asynq.Option) (taskId string, err error) {
	if !job.IsValid() {
		err = errors.Errorf("job not support, job=%s", job.String())
		return
	}

	task, err := inst.newTask(job, params)
	if err != nil {
		return
	}

	opts := inst.options(options)
	info, err := inst.data.queueClient.Enqueue(task, opts...)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	taskId = info.ID
	return
}

func (inst *Queue) GetJobConcurrency(job cnst.Job) (concurrency int, err error) {
	servers, err := inst.data.queueInspector.Servers()
	if err != nil {
		return
	}
	for _, server := range servers {
		for k := range server.Queues {
			if k == job.String() {
				concurrency += server.Concurrency
			}
		}
	}

	return
}

func (inst *Queue) Archive(job cnst.Job, taskId string) (err error) {
	_, err = inst.data.queueInspector.GetTaskInfo(job.String(), taskId)
	if err != nil {
		if errors.Is(err, asynq.ErrTaskNotFound) {
			log.Errorf("task not found: %s", taskId)
			err = nil
		}
		return
	}

	// running task cannot archive, cancel first
	_ = inst.data.queueInspector.CancelProcessing(taskId)

	err = inst.data.queueInspector.ArchiveTask(job.String(), taskId)
	if err != nil {
		time.Sleep(time.Microsecond * 100)
		// if archvie fail, retry once
		err = inst.data.queueInspector.ArchiveTask(job.String(), taskId)
	}

	return
}

// wrap a new task with params
func (inst *Queue) newTask(key cnst.Job, params any) (task *asynq.Task, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		err = errors.Wrap(err, "convert params to payload fail")
		return
	}

	task = asynq.NewTask(key.String(), payload)
	return
}

// 填充默认记录保留时间
func (inst *Queue) options(opts []asynq.Option) (options []asynq.Option) {
	defaultOptions := []asynq.Option{
		asynq.Retention(inst.jobConfig.GetDefaultConfig().GetRetention().AsDuration()), // 任务默认保留时间
	}

	// 实例传入的opt会覆盖默认opt
	options = append(defaultOptions, opts...)

	return
}
