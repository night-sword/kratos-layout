package param

import (
	"github.com/pkg/errors"
)

// ------ params ------- //
type JobParam interface {
	GetTaskId() string
}

type abstractJobParam struct {
	TaskId string
}

func (inst *abstractJobParam) GetTaskId() string {
	return inst.TaskId
}

func ConvertParam[T JobParam](params JobParam) (p T, err error) {
	p, ok := params.(T)
	if !ok {
		err = errors.New("convert param fail")
	}

	return
}

func newAbstractJobParam(taskId string) *abstractJobParam {
	return &abstractJobParam{
		TaskId: taskId,
	}
}
