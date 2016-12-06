package todo

import (
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	Value "github.com/kkserver/kk-lib/kk/value"
	"reflect"
	"time"
)

type TodoUserService struct {
	app.Service
	Join   *TodoUserJoinTask
	Remove *TodoUserRemoveTask
	Get    *TodoUserTask
	Query  *TodoQueryTask
}

func (S *TodoUserService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *TodoUserService) HandleTodoUserJoinTask(a *TodoApp, task *TodoUserJoinTask) error {

	if task.TodoId == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found todoId"
	}

	if task.Uid == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_UID
		task.Result.Errmsg = "Not found uid"
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var v = TodoUser{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &a.TodoUserTable, a.DB.Prefix, " WHERE todoid=? AND uid=?", task.TodoId, task.Uid)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.User = &v

	} else {

		v.TodoId = task.TodoId
		v.Ctime = time.Now().Unix()
		_, err = kk.DBInsert(db, &a.TodoUserTable, a.DB.Prefix, &v)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.User = &v
	}

	return nil
}

func (S *TodoUserService) HandleTodoUserRemoveTask(a *TodoApp, task *TodoUserRemoveTask) error {

	if task.TodoId == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found todoId"
	}

	if task.Uid == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_UID
		task.Result.Errmsg = "Not found uid"
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	rows, err := kk.DBQuery(db, &a.TodoUserTable, a.DB.Prefix, " WHERE todoid=? AND uid=?", task.TodoId, task.Uid)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	var v = TodoUser{}
	var scanner = kk.NewDBScaner(&v)

	if rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = kk.DBDelete(db, &a.TodoUserTable, a.DB.Prefix, " WHERE todoid=? AND uid=?", task.TodoId, task.Uid)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.User = &v

	} else {
		task.Result.Errno = ERROR_TODO_USER_NOT_FOUND
		task.Result.Errmsg = fmt.Sprintf("Not found todo user todoId:%d uid:%d", task.TodoId, task.Uid)
	}

	return nil
}

func (S *TodoUserService) HandleTodoUserTask(a *TodoApp, task *TodoUserTask) error {

	if task.TodoId == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_ID
		task.Result.Errmsg = "Not found todoId"
	}

	if task.Uid == 0 {
		task.Result.Errno = ERROR_TODO_NOT_FOUND_UID
		task.Result.Errmsg = "Not found uid"
	}

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	rows, err := kk.DBQuery(db, &a.TodoUserTable, a.DB.Prefix, " WHERE todoid=? AND uid=?", task.TodoId, task.Uid)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	var v = TodoUser{}
	var scanner = kk.NewDBScaner(&v)

	if rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		task.Result.User = &v

	} else {
		task.Result.Errno = ERROR_TODO_USER_NOT_FOUND
		task.Result.Errmsg = fmt.Sprintf("Not found todo user todoId:%d uid:%d", task.TodoId, task.Uid)
	}

	return nil
}

func (S *TodoUserService) HandleTodoQueryTask(a *TodoApp, task *TodoUserQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	var users = []TodoUser{}

	var args = []interface{}{}

	var sql = " WHERE 1"

	if task.TodoId != nil {
		sql = sql + " AND todoid=?"
		args = append(args, Value.IntValue(reflect.ValueOf(task.TodoId), 0))
	}

	if task.Uid != nil {
		sql = sql + " AND uid=?"
		args = append(args, Value.IntValue(reflect.ValueOf(task.Uid), 0))
	}

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = TodoUserQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.PageSize, err = kk.DBQueryCount(db, &a.TodoUserTable, a.DB.Prefix, sql, args...)
		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}
	}

	if task.OrderBy == "asc" {
		sql = sql + " ORDER BY id ASC"
	} else {
		sql = sql + " ORDER BY id DESC"
	}

	sql = fmt.Sprintf("%s LIMIT %d,%d", sql, (pageIndex-1)*pageSize, pageSize)

	var v = TodoUser{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := kk.DBQuery(db, &a.TodoUserTable, a.DB.Prefix, sql, args...)

	if err != nil {
		task.Result.Errno = ERROR_TODO
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_TODO
			task.Result.Errmsg = err.Error()
			return nil
		}

		users = append(users, v)
	}

	task.Result.Users = users

	return nil
}
