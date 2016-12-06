package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoSetResult struct {
	app.Result
	Todo *Todo `json:"todo,omitempty"`
}

type TodoSetTask struct {
	app.Task
	Id     int64       `json:"id,string"`
	Pid    interface{} `json:"pid,string"`
	Uid    interface{} `json:"uid,string"`
	Title  interface{} `json:"title,string,omitempty"`
	Result TodoSetResult
}

func (T *TodoSetTask) GetResult() interface{} {
	return &T.Result
}
