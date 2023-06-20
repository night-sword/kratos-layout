package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/night-sword/kratos-layout/internal/cnst"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/job"
	"github.com/night-sword/kratos-layout/internal/job/param"
)

type QueueServers struct {
	config     *conf.Job
	dataConfig *conf.Data
	servers    []*QueueServer
}

func NewQueueServers(
	config *conf.Job, dataConfig *conf.Data,
	demoJob *job.DemoJob,
) *QueueServers {
	servers := []*QueueServer{
		NewQueueServer[*param.DemoJobParams](config, dataConfig, cnst.Job_Demo, demoJob),
	}

	return &QueueServers{
		config:     config,
		dataConfig: dataConfig,
		servers:    servers,
	}
}

func (inst *QueueServers) GetServers() []*QueueServer {
	return inst.servers
}

// ------- QueueServer ------- //

type QueueServer struct {
	jobConfig  *conf.Job
	dataConfig *conf.Data

	server *asynq.Server
	mux    *asynq.ServeMux
}

func (inst *QueueServer) Start(_ context.Context) (err error) {
	return inst.server.Start(inst.mux)
}

func (inst *QueueServer) Stop(_ context.Context) (err error) {
	inst.server.Shutdown()
	return
}

func NewQueueServer[P param.JobParam](config *conf.Job, dataConfig *conf.Data, name cnst.Job, job job.Runable) (server *QueueServer) {
	mux := newQueueMux[P](name, job)

	server = &QueueServer{
		jobConfig:  config,
		dataConfig: dataConfig,
		mux:        mux,
	}
	server.server = server.newServer(name)
	return
}

func (inst *QueueServer) newServer(job cnst.Job) (srv *asynq.Server) {
	redisCfg := inst.dataConfig.GetRedis()
	redisOpts := asynq.RedisClientOpt{
		Addr:     redisCfg.GetAddr(),
		Password: redisCfg.GetPwd(),
	}

	concurrency := inst.concurrency(job)
	asynqCfg := asynq.Config{
		Concurrency: concurrency,
		Queues: map[string]int{
			job.String(): 100,
		},
		RetryDelayFunc: inst.getRetryDelayFn(job),
	}

	log.Infof("start queue server: job=%s concurrency=%d ", job, concurrency)

	return asynq.NewServer(redisOpts, asynqCfg)
}

func (inst *QueueServer) getRetryDelayFn(job cnst.Job) asynq.RetryDelayFunc {
	retryDelay := inst.retryDelay(job)
	if retryDelay == nil {
		return asynq.DefaultRetryDelayFunc
	}

	return func(n int, e error, t *asynq.Task) time.Duration {
		return retryDelay.AsDuration()
	}
}

func (inst *QueueServer) concurrency(job cnst.Job) (c int) {
	c = int(inst.jobConfig.GetDefaultConfig().GetConcurrency())

	opts := inst.options(job)
	if opts != nil {
		c = int(opts.GetConcurrency())
	}

	if c <= 0 {
		log.Warnf("job concurrency is 0, job will not run, job=%s", job)
	}

	return
}

func (inst *QueueServer) retryDelay(job cnst.Job) (delay *durationpb.Duration) {
	delay = inst.jobConfig.GetDefaultConfig().GetRetryDelay()

	opts := inst.options(job)
	if opts != nil {
		d := opts.GetRetryDelay()
		if d.AsDuration() != 0 {
			delay = d
		}
	}

	return
}

func (inst *QueueServer) options(job cnst.Job) (options *conf.Job_JobOption) {
	jobs := inst.jobConfig.GetJobs()
	if len(jobs) == 0 {
		log.Warn("without jobs options")
		return
	}

	options, ok := jobs[job.String()]
	if !ok {
		log.Warnf("job without options, job=%s", job)
	}

	return
}

func newQueueMux[P param.JobParam](name cnst.Job, runable job.Runable) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	handler := job.NewHandler[P](runable.Run)
	mux.HandleFunc(name.String(), handler)

	return mux
}
