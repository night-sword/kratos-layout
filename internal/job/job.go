package job

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/night-sword/kratos-layout/internal/cnst"
	"github.com/night-sword/kratos-layout/internal/job/param"
)

// ------- Job 接口定义 ------- //
type Runable interface {
	GetJobName() (name string)
	Run(ctx context.Context, params param.JobParam) (err error)
}

// ------- 抽象Job类 ------- //
type abstractJob struct {
	job cnst.Job
}

func newAbstractJob(job cnst.Job) *abstractJob {
	return &abstractJob{
		job: job,
	}
}

func (inst *abstractJob) GetJobName() (name string) {
	return inst.job.String()
}

// 任务是否可重试
func (inst *abstractJob) isRetryable(ctx context.Context, runErr error) (retryable bool) {
	if errors.Is(runErr, asynq.SkipRetry) {
		return true
	}

	retried, err := inst.getRetryCount(ctx)
	if err != nil {
		return
	}
	maxRetry, err := inst.getMaxRetry(ctx)
	if err != nil {
		return
	}

	return retried >= maxRetry
}

func (inst *abstractJob) getRetryCount(ctx context.Context) (count int, err error) {
	count, ok := asynq.GetRetryCount(ctx)
	if !ok {
		err = errors.New("cannot get retried from ctx, skip set runtime")
	}
	return
}

func (inst *abstractJob) getMaxRetry(ctx context.Context) (count int, err error) {
	count, ok := asynq.GetMaxRetry(ctx)
	if !ok {
		err = errors.New("cannot get retried from ctx, skip set runtime")
	}
	return
}

// ------ 获得 asynqHandler -------//
type AsynqHandler func(ctx context.Context, t *asynq.Task) (err error)

//  Run(ctx context.Context, params JobParams) (err error)

func NewHandler[T param.JobParam](run func(ctx context.Context, params param.JobParam) error) AsynqHandler {
	return func(ctx context.Context, task *asynq.Task) (err error) {
		var p T
		err = json.Unmarshal(task.Payload(), &p)
		if err != nil {
			return
		}

		err = run(ctx, p)
		return
	}
}
