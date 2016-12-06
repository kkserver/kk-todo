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
	Id        int64       `json:"id,string"`
	Pid       interface{} `json:"pid,string"`
	Uid       interface{} `json:"uid,string"`
	Path      interface{} `json:"path,string"`
	OrderBy   string      `json:"orderBy"` // desc, asc
	PageIndex int         `json:"p"`
	PageSize  int         `json:"size"`
	Counter   bool        `json:"counter"`
	Result    TodoQueryResult
}

func (T *TodoQueryTask) GetResult() interface{} {
	return &T.Result
}
