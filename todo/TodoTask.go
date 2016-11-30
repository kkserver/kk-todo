package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoResult
}

func (T *TodoTask) GetResult() interface{} {
	return &T.Result
}
