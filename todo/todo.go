package todo

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type Todo struct {
	Id    int64  `json:"id,string"`
	Pid   int64  `json:"pid,string"`
	Uid   int64  `json:"uid,string"`
	Path  string `json:"path"`
	Title string `json:"title,omitempty"`
	Mtime int64  `json:"mtime"` //修改时间
	Ctime int64  `json:"ctime"` //创建时间
}

type TodoUser struct {
	Id     int64 `json:"id,string"`
	TodoId int64 `json:"todoId,string"`
	Uid    int64 `json:"uid,string"`
	Ctime  int64 `json:"ctime"` //创建时间
}

type TodoApp struct {
	app.App
	DB     app.DBConfig
	Remote *remote.Service

	Todo     *TodoService
	TodoUser *TodoUserService

	TodoTable     kk.DBTable
	TodoUserTable kk.DBTable
}

func (C *TodoApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}
