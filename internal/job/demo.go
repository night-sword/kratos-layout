package job

import (
	"context"

	"github.com/night-sword/kratos-layout/internal/cnst"
	"github.com/night-sword/kratos-layout/internal/job/param"
)

func NewDemoJob() *DemoJob {
	return &DemoJob{
		abstractJob: newAbstractJob(cnst.Job(0)),
	}
}

type DemoJob struct {
	*abstractJob
	// add other biz or other object
}

func (inst *DemoJob) Run(ctx context.Context, params param.JobParam) (err error) {
	// p, err := param.ConvertParam[*param.DemoJobParams](params)
	// if err != nil {
	//     return
	// }

	return
}
