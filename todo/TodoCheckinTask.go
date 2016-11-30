package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoCheckinResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoCheckinTask struct {
	app.Task
	Id     int64 `json:"id,string"`
	Result TodoCheckinResult
}

func (T *TodoCheckinTask) GetResult() interface{} {
	return &T.Result
}
