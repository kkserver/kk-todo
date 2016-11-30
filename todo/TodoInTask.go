package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoInResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoInTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoInResult
}

func (T *TodoInTask) GetResult() interface{} {
	return &T.Result
}
