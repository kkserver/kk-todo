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
	Pid      int64  `json:"pid,string"`
	Uid      int64  `json:"uid,string"`
	Title    string `json:"title,omitempty"`
	Summary  string `json:"summary,omitempty"`
	MinValue int64  `json:"minValue,string"` //最低价格
	MaxValue int64  `json:"maxValue,string"` //最高价格
	MaxCount int    `json:"maxCount"`        //最大派发数
	Result   TodoCreateResult
}

func (T *TodoCreateTask) GetResult() interface{} {
	return &T.Result
}
