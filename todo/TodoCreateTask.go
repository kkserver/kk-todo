package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoCreateResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoCreateTask struct {
	app.Task
	Pid    int64  `json:"pid,string"`
	Uid    int64  `json:"uid,string"`
	Title  string `json:"title,omitempty"`
	Result TodoCreateResult
}

func (T *TodoCreateTask) GetResult() interface{} {
	return &T.Result
}
