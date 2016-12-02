package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoRemoveResult struct {
	app.Result
}

type TodoRemoveTask struct {
	app.Task
	Id     int64  `json:"id,string"`
	Pid    int64  `json:"pid,string"`
	Uid    int64  `json:"uid,string"`
	Ids    string `json:"ids,string"`
	Pids   string `json:"pids,string"`
	Uids   string `json:"uids,string"`
	Result TodoRemoveResult
}

func (T *TodoRemoveTask) GetResult() interface{} {
	return &T.Result
}
