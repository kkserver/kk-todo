package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoWorkingResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoWorkingTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoWorkingResult
}

func (T *TodoWorkingTask) GetResult() interface{} {
	return &T.Result
}
