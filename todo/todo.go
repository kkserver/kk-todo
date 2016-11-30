package todo

import (
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

const TodoStatusNone = 0     // 新创建的
const TodoStatusIn = 1       // 派单中
const TodoStatusWorking = 2  // 工作中
const TodoStatusCheckin = 3  // 验收中
const TodoStatusFinish = 200 // 完成
const TodoStatusCancel = 300 // 取消任务

type Todo struct {
	Id          int64  `json:"id,string"`
	Pid         int64  `json:"pid,string"`
	Uid         int64  `json:"uid,string"`
	Title       string `json:"title,omitempty"`
	Summary     string `json:"summary,omitempty"`
	MinValue    int64  `json:"minValue,string"` //最低价格
	MaxValue    int64  `json:"maxValue,string"` //最高价格
	MaxCount    int    `json:"maxCount"`        //最大派发数
	Count       int    `json:"count"`           //派发数
	Status      int    `json:"status"`
	CheckinTime int64  `json:"checkinTime"` //验收时间
	Mtime       int64  `json:"mtime"`       //修改时间
	Ctime       int64  `json:"ctime"`       //创建时间
}

const TodoUserStatusNone = 0    // 等待确认
const TodoUserStatusIn = 1      // 确认接单
const TodoUserStatusWorking = 2 // 工作中
const TodoUserStatusCheckin = 3 // 验收中
const TodoUserStatusGoback = 4  // 退回

const TodoUserStatusFinish = 200   // 完成
const TodoUserStatusCancel = 300   // 用户取消
const TodoUserStatusRejected = 400 // 被拒绝当 Todo.status==TodoStatusIn 时可被拒绝

type TodoUser struct {
	Id      int64 `json:"id,string"`
	TodoId  int64 `json:"todoId,string"`
	Uid     int64 `json:"uid,string"`
	Value   int64 `json:"value,string"` //收益
	Status  int   `json:"status"`
	Expires int64 `json:"expires"` //确认超时时间
	Ctime   int64 `json:"ctime"`   //创建时间
}

type TodoLog struct {
	Id     int64  `json:"id,string"`
	TodoId int64  `json:"todoId,string"`
	Uid    int64  `json:"uid,string"`
	Body   string `json:"body"`
	Ctime  int64  `json:"ctime"` //创建时间
}

type TodoTag struct {
	Id       int64 `json:"id,string"`
	TodoId   int64 `json:"todoId,string"`
	TagId    int64 `json:"tagId,string"`
	MinValue int   `json:"minValue"` //最小能力值
	MaxValue int   `json:"maxValue"` //最大能力值
	Ctime    int64 `json:"ctime"`    //创建时间
}

type TodoUserTag struct {
	Id     int64 `json:"id,string"`
	TodoId int64 `json:"todoId,string"`
	TagId  int64 `json:"tagId,string"`
	Uid    int64 `json:"uid,string"`
	Value  int   `json:"value"` //用户能力值
	Ctime  int64 `json:"ctime"` //创建时间
}

type TodoApp struct {
	app.App
	DB     app.DBConfig
	Remote *remote.Service

	Todo *TodoService

	TodoTable    kk.DBTable
	UserTable    kk.DBTable
	LogTable     kk.DBTable
	TagTable     kk.DBTable
	UserTagTable kk.DBTable
}
