package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoUserRemoveResult struct {
	app.Result
	User *TodoUser `json:"user,omitempty"`
}

type TodoUserRemoveTask struct {
	app.Task
	TodoId int64 `json:"todoId,string"`
	Uid    int64 `json:"uid,string"`
	Result TodoUserRemoveResult
}

func (T *TodoUserRemoveTask) GetResult() interface{} {
	return &T.Result
}
