package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoUserQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
}

type TodoUserQueryResult struct {
	app.Result
	Users   []TodoUser            `json:"users,omitempty"`
	Counter *TodoUserQueryCounter `json:"counter,omitempty"`
}

type TodoUserQueryTask struct {
	app.Task
	Uid       interface{} `json:"uid,string"`
	TodoId    interface{} `json:"todoId,string"`
	OrderBy   string      `json:"orderBy"` // asc,desc
	PageIndex int         `json:"p"`
	PageSize  int         `json:"size"`
	Counter   bool        `json:"counter"`
	Result    TodoUserQueryResult
}

func (T *TodoUserQueryTask) GetResult() interface{} {
	return &T.Result
}
