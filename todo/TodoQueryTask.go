package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
}

type TodoQueryResult struct {
	app.Result
	Counter *TodoQueryCounter `json:"counter,omitempty"`
	Todos   []Todo            `json:"todos,omitempty"`
}

type TodoQueryTask struct {
	app.Task
	Id        int64  `json:"id,string"`
	Pid       int64  `json:"pid,string"`
	Uid       int64  `json:"uid,string"`
	OrderBy   string `json:"orderBy"`
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    TodoQueryResult
}

func (T *TodoQueryTask) GetResult() interface{} {
	return &T.Result
}
