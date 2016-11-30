package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoFinishResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoFinishTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoFinishResult
}

func (T *TodoFinishTask) GetResult() interface{} {
	return &T.Result
}
