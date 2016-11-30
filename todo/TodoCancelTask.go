package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoCancelResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoCancelTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoCancelResult
}

func (T *TodoCancelTask) GetResult() interface{} {
	return &T.Result
}
