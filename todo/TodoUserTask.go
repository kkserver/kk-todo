package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoUserResult struct {
	app.Result
	User *TodoUser `json:"user,omitempty"`
}

type TodoUserTask struct {
	app.Task
	TodoId int64 `json:"todoId,string"`
	Uid    int64 `json:"uid,string"`
	Result TodoUserResult
}

func (T *TodoUserTask) GetResult() interface{} {
	return &T.Result
}
