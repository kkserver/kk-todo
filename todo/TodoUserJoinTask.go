package todo

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type TodoUserJoinResult struct {
	app.Result
	User *TodoUser `json:"user,omitempty"`
}

type TodoUserJoinTask struct {
	app.Task
	TodoId int64 `json:"todoId,string"`
	Uid    int64 `json:"uid,string"`
	Result TodoUserJoinResult
}

func (T *TodoUserJoinTask) GetResult() interface{} {
	return &T.Result
}
