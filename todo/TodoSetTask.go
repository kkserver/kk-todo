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
	Id       int64  `json:"id,string"`
	Pid      int64  `json:"pid,string"`
	Uid      int64  `json:"uid,string"`
	Title    string `json:"title,omitempty"`
	Summary  string `json:"summary,omitempty"`
	MinValue int64  `json:"minValue,string"` //最低价格
	MaxValue int64  `json:"maxValue,string"` //最高价格
	MaxCount int    `json:"maxCount"`        //最大派发数
	Result   TodoSetResult
}

func (T *TodoSetTask) GetResult() interface{} {
	return &T.Result
}
